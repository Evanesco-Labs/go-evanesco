package zkpminer

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/zkpminer/keypair"
	"github.com/ethereum/go-ethereum/zkpminer/problem"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var (
	ErrorMinerWorkerOutOfRange    = errors.New("miner's workers reach MaxWorkerCnt, can not add more workers")
	ErrorLocalMinerWithoutBackend = errors.New("new local miner with nil backend")
	ErrorBlockHeaderSubscribe     = errors.New("block header subscribe error")
)

type TaskStep int

const (
	TASKSTART TaskStep = iota
	TASKWAITCHALLENGEBLOCK
	TASKGETCHALLENGEBLOCK
	TASKPROBLEMSOLVED
	TASKSUBMITTED
)

const (
	COINBASEINTERVAL = types.CoinBaseInterval
	SUBMITADVANCE    = types.SubmitAdvance
	RPCTIMEOUT       = time.Minute
)

var WSUrlTryRound = 3
var RetryWSRPCWaitDuration = time.Second * 5

type Backend interface {
	BlockChain() *core.BlockChain
	EventMux() *event.TypeMux
}

type Task struct {
	CoinbaseAddr     common.Address
	minerAddr        common.Address
	Step             TaskStep
	lastCoinBaseHash [32]byte
	challengeHeader  *types.Header
	challengeIndex   Height
	lottery          *types.Lottery
	signature        [65]byte
}

func (t *Task) SetHeader(h *types.Header) {
	t.challengeHeader = h
	t.lottery.ChallengeHeaderHash = h.Hash()
	t.Step = TASKGETCHALLENGEBLOCK
}

func (t *Task) SetCoinbaseAddr(coinbaseAddr common.Address) {
	t.CoinbaseAddr = coinbaseAddr
}

//SetTaskMinerAddr only use in TASKSTART step
func SetTaskMinerAddr(template *Task, minerAddr common.Address) Task {
	if template.Step != TASKSTART {
		panic("only use it to update task in step TASKSTART")
	}
	//Deep Copy task
	return Task{
		minerAddr:        minerAddr,
		CoinbaseAddr:     template.CoinbaseAddr,
		Step:             TASKSTART,
		lastCoinBaseHash: template.lastCoinBaseHash,
		challengeIndex:   Height(uint64(0)),
		lottery: &types.Lottery{
			MinerAddr:    minerAddr,
			CoinbaseAddr: template.CoinbaseAddr,
		},
	}
}

type Config struct {
	MinerList        []keypair.Key
	MaxWorkerCnt     int32
	MaxTaskCnt       int32
	CoinbaseInterval uint64
	SubmitAdvance    uint64
	CoinbaseAddr     common.Address
	WsUrl            []string
	RpcTimeout       time.Duration
	PkPath           string
}

func DefaultConfig() Config {
	return Config{
		MinerList:        make([]keypair.Key, 0),
		MaxWorkerCnt:     1,
		MaxTaskCnt:       1,
		CoinbaseInterval: COINBASEINTERVAL,
		SubmitAdvance:    SUBMITADVANCE,
		CoinbaseAddr:     common.Address{},
		WsUrl:            []string{},
		RpcTimeout:       RPCTIMEOUT,
		PkPath:           "./QmNpJg4jDFE4LMNvZUzysZ2Ghvo4UJFcsjguYcx4dTfwKx",
	}
}

func (config *Config) Customize(minerList []keypair.Key, coinbase common.Address, url []string, pkPath string) {
	config.MinerList = minerList

	config.CoinbaseAddr = coinbase

	if len(url) != 0 {
		config.WsUrl = append(config.WsUrl, url...)
	}

	if pkPath != "" {
		config.PkPath = pkPath
	}
}

type Miner struct {
	mu               sync.RWMutex
	config           Config
	zkpProver        *problem.Prover
	MaxWorkerCnt     int32
	MaxTaskCnt       int32
	CoinbaseAddr     common.Address
	workers          map[common.Address]*Worker
	scanner          *Scanner
	coinbaseInterval Height
	submitAdvance    Height
	urlList          []string
	exitCh           chan struct{}
}

func NewLocalMiner(config Config, backend Backend) (*Miner, error) {
	runtime.GOMAXPROCS(1)
	zkpProver, err := problem.NewProblemProver(config.PkPath)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	log.Info("Init ZKP Problem worker success!")

	miner := Miner{
		mu:               sync.RWMutex{},
		config:           config,
		zkpProver:        zkpProver,
		MaxWorkerCnt:     config.MaxWorkerCnt,
		MaxTaskCnt:       config.MaxTaskCnt,
		CoinbaseAddr:     config.CoinbaseAddr,
		workers:          make(map[common.Address]*Worker),
		coinbaseInterval: Height(config.CoinbaseInterval),
		submitAdvance:    Height(config.SubmitAdvance),
		exitCh:           make(chan struct{}),
		urlList:          config.WsUrl,
	}

	if backend == nil {
		return nil, ErrorLocalMinerWithoutBackend
	}

	explorer := LocalExplorer{
		Backend:  backend,
		headerCh: make(chan *types.Header),
	}
	blockEventCh := make(chan core.ChainHeadEvent)
	sub := backend.BlockChain().SubscribeChainHeadEvent(blockEventCh)
	go func() {
		for {
			select {
			case blockEvent := <-blockEventCh:
				explorer.headerCh <- blockEvent.Block.Header()
			case <-sub.Err():
				log.Error(ErrorBlockHeaderSubscribe.Error())
				miner.Close()
			}
		}
	}()
	miner.NewScanner(&explorer)
	miner.StartScanner()

	go miner.Loop()
	//add new workers
	for _, key := range config.MinerList {
		miner.NewWorker(key)
	}
	log.Info("miner start")
	return &miner, nil
}

func NewMiner(config Config) (*Miner, error) {
	runtime.GOMAXPROCS(1)
	zkpProver, err := problem.NewProblemProver(config.PkPath)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	log.Info("Init ZKP Problem worker success!")
	if len(config.WsUrl) == 0 {
		panic("Evanesco websocket url unset")
	}
	miner := Miner{
		mu:               sync.RWMutex{},
		config:           config,
		zkpProver:        zkpProver,
		MaxWorkerCnt:     config.MaxWorkerCnt,
		MaxTaskCnt:       config.MaxTaskCnt,
		CoinbaseAddr:     config.CoinbaseAddr,
		workers:          make(map[common.Address]*Worker),
		coinbaseInterval: Height(config.CoinbaseInterval),
		submitAdvance:    Height(config.SubmitAdvance),
		exitCh:           make(chan struct{}),
		urlList:          config.WsUrl,
	}

	explorer := RpcExplorer{
		Client:     new(rpc.Client),
		Sub:        new(rpc.ClientSubscription),
		HeaderCh:   make(chan *types.Header),
		rpcTimeOut: config.RpcTimeout,
		WsUrl:      "",
	}

	miner.NewScanner(&explorer)
	miner.updateWS()

	go func() {
		for {
			err := <-explorer.Sub.Err()
			log.Warn(ErrorBlockHeaderSubscribe.Error(), "err", err)
			log.Info("try to connect another node")
			miner.updateWS()
		}
	}()

	go miner.Loop()
	//add new workers
	for _, key := range config.MinerList {
		miner.NewWorker(key)
	}
	log.Info("miner start")
	return &miner, nil
}

func (m *Miner) updateWS() {

	if m.scanner.IsUpdating() {
		return
	}

	m.scanner.close()
	exp, ok := m.scanner.explorer.(*RpcExplorer)
	if !ok {
		panic("Full node miner cannot update ws client")
		return
	}

	//set scanner status updating
	atomic.StoreInt32(&m.scanner.running, int32(2))

	res := false
	var err error
	for i := 0; i < WSUrlTryRound; i++ {
		for _, url := range m.urlList {
			time.Sleep(RetryWSRPCWaitDuration)
			if url == exp.WsUrl {
				continue
			}
			exp.Client, err = rpc.Dial(url)
			if err != nil {
				log.Warn("Websocket dial Evanesco node err", "err", err)
				continue
			}
			exp.Sub, err = exp.Client.EthSubscribe(context.Background(), exp.HeaderCh, "newHeads")
			if err != nil {
				log.Warn("Subscribe block err", "err", err)
				continue
			}

			res = true
			exp.WsUrl = url
			log.Info("Connected node WebSocket URL", "url", url)
			break
		}
		if res == true {
			break
		} else {
			//reset url to try this url again
			exp.WsUrl = ""
		}
	}

	if res == false {
		log.Error("Dial all websocket urls failed")
		os.Exit(1)
	}

	m.StartScanner()
	return
}

func (m *Miner) Close() {
	defer func() {
		if recover() != nil {
		}
	}()
	//close workers
	for _, worker := range m.workers {
		worker.close()
	}
	//close scanner
	m.scanner.close()
	close(m.exitCh)
	os.Exit(1)
}

func (m *Miner) NewWorker(minerKey keypair.Key) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if len(m.workers) == int(m.MaxWorkerCnt) {
		log.Error(ErrorMinerWorkerOutOfRange.Error())
		return
	}
	worker := Worker{
		mu:               sync.RWMutex{},
		running:          0,
		MaxTaskCnt:       m.MaxTaskCnt,
		CoinbaseAddr:     m.CoinbaseAddr,
		minerAddr:        minerKey.Address,
		pk:               minerKey.PrivateKey.Public(),
		sk:               &minerKey.PrivateKey,
		workingTaskCnt:   0,
		coinbaseInterval: m.coinbaseInterval,
		inboundTaskCh:    make(chan *Task),
		submitAdvance:    m.submitAdvance,
		scanner:          m.scanner,
		zkpProver:        m.zkpProver,
		exitCh:           make(chan struct{}),
	}

	m.workers[minerKey.Address] = &worker
	go worker.Loop()
	worker.start()
	log.Debug("worker start")
}

func (m *Miner) CloseWorker(addr common.Address) {
	if worker, ok := m.workers[addr]; ok {
		worker.close()
		delete(m.workers, addr)
	}
}

func (m *Miner) StopWorker(addr common.Address) {
	if worker, ok := m.workers[addr]; ok {
		worker.stop()
	}
}

func (m *Miner) StartWorker(addr common.Address) {
	if worker, ok := m.workers[addr]; ok {
		worker.start()
	}
}

func (m *Miner) Loop() {
	for {
		select {
		case <-m.exitCh:
			return
		case taskTem := <-m.scanner.outboundTaskCh:
			if taskTem.Step == TASKSTART {
				for _, worker := range m.workers {
					task := SetTaskMinerAddr(taskTem, worker.minerAddr)
					worker.inboundTaskCh <- &task
				}
				continue
			}
			if taskTem.Step == TASKGETCHALLENGEBLOCK {
				if worker, ok := m.workers[taskTem.minerAddr]; ok {
					worker.inboundTaskCh <- taskTem
					continue
				}
				log.Warn("worker for this task not exist")
			}
			if taskTem.Step == TASKSUBMITTED {
				//todo: store submitted lotteries for later queries
			}
		}
	}
}

func (m *Miner) NewScanner(explorer Explorer) {
	m.scanner = &Scanner{
		miner:              m,
		mu:                 sync.RWMutex{},
		CoinbaseAddr:       m.CoinbaseAddr,
		BestScore:          zero,
		LastBlockHeight:    0,
		CoinbaseInterval:   m.coinbaseInterval,
		LastCoinbaseHeight: 0,
		taskWait:           make(map[Height][]*Task),
		inboundTaskCh:      make(chan *Task),
		outboundTaskCh:     make(chan *Task),
		explorer:           explorer,
		exitCh:             make(chan struct{}),
		running:            int32(0),
	}
}

func (m *Miner) StartScanner() {
	m.scanner.close()
	m.scanner.exitCh = make(chan struct{})
	go m.scanner.Loop()
	atomic.StoreInt32(&m.scanner.running, int32(1))
}
