package zkpminer

import (
	"context"
	"errors"
	"github.com/Evanesco-Labs/go-evanesco/accounts/abi/bind"
	"github.com/Evanesco-Labs/go-evanesco/common"
	"github.com/Evanesco-Labs/go-evanesco/core"
	"github.com/Evanesco-Labs/go-evanesco/core/types"
	"github.com/Evanesco-Labs/go-evanesco/evaclient"
	"github.com/Evanesco-Labs/go-evanesco/event"
	"github.com/Evanesco-Labs/go-evanesco/log"
	"github.com/Evanesco-Labs/go-evanesco/rpc"
	"github.com/Evanesco-Labs/go-evanesco/zkpminer/keypair"
	"github.com/Evanesco-Labs/go-evanesco/zkpminer/problem"
	"math/rand"
	"os"
	"runtime"
	"runtime/debug"
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

var WSUrlTryRound = 5
var RetryWSRPCWaitDuration = time.Second * 5
var RetryJitterMaxDuration = 1000 //millisecond

type Backend interface {
	BlockChain() *core.BlockChain
	EventMux() *event.TypeMux
}

type Task struct {
	CoinbaseAddr     common.Address
	minerAddr        common.Address
	Step             TaskStep
	lastCoinBaseHash [32]byte
	challengeHeader  types.HeaderShort
	challengeIndex   Height
	lottery          *types.Lottery
	signature        [65]byte
}

func (t *Task) SetHeader(h types.HeaderShort) {
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
		PkPath:           "./QmQL4k1hKYiW3SDtMREjnrah1PBsak1VE3VgEqTyoDckz9",
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
	isEffective      sync.Once
	config           Config
	zkpProver        *problem.Prover
	MaxWorkerCnt     int32
	MaxTaskCnt       int32
	CoinbaseAddr     common.Address
	Workers          map[common.Address]*Worker
	scanner          *Scanner
	coinbaseInterval Height
	submitAdvance    Height
	urlList          []string
	exitCh           chan struct{}
}

func NewLocalMiner(config Config, backend Backend) (*Miner, error) {
	runtime.GOMAXPROCS(1)
	if backend == nil {
		return nil, ErrorLocalMinerWithoutBackend
	}
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
		Workers:          make(map[common.Address]*Worker),
		coinbaseInterval: Height(config.CoinbaseInterval),
		submitAdvance:    Height(config.SubmitAdvance),
		exitCh:           make(chan struct{}),
		urlList:          config.WsUrl,
		isEffective:      sync.Once{},
	}

	checkEffective := func() {
		//check effective
		minerAddress := config.MinerList[0].Address
		ok, coinbasePledge := Iseffective(minerAddress, backend.BlockChain().InprocHandler)
		if !ok {
			log.Error("Miner address not staked", "address", minerAddress.String())
		}
		emptyAddr := common.Address{}
		//coinbase address is not set on pledge, use miner address by default
		if coinbasePledge == emptyAddr {
			if miner.CoinbaseAddr == emptyAddr {
				miner.CoinbaseAddr = minerAddress
				return
			}
			return
		}
		//coinbase address is set, use pledge coinbase address by default
		if coinbasePledge != emptyAddr {
			if miner.CoinbaseAddr == emptyAddr {
				miner.CoinbaseAddr = coinbasePledge
				return
			}
			if miner.CoinbaseAddr != coinbasePledge {
				log.Error(NotPledgeCoinbaseError.Error())
				log.Info("miner coinbase address:" + miner.CoinbaseAddr.String() + ", fortress coinbase address:" + coinbasePledge.String())
				return
			}
			return
		}
	}

	miner.isEffective.Do(checkEffective)

	explorer := LocalExplorer{
		Backend:  backend,
		headerCh: make(chan types.HeaderShort),
	}
	blockEventCh := make(chan core.ChainHeadEvent)
	sub := backend.BlockChain().SubscribeChainHeadEvent(blockEventCh)
	go func() {
		for {
			select {
			case blockEvent := <-blockEventCh:
				short := blockEvent.Block.Header().Short()
				explorer.headerCh <- short
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
	log.Info("waiting for next mining epoch")
	return &miner, nil
}

func NewMiner(config Config) (*Miner, error) {
	runtime.GOMAXPROCS(1)
	zkpProver, err := problem.NewProblemProver(config.PkPath)
	debug.FreeOSMemory()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	log.Info("Init ZKP Problem worker success!")
	if len(config.WsUrl) == 0 {
		Fatalf("Evanesco websocket url unset")
	}
	miner := Miner{
		mu:               sync.RWMutex{},
		config:           config,
		zkpProver:        zkpProver,
		MaxWorkerCnt:     config.MaxWorkerCnt,
		MaxTaskCnt:       config.MaxTaskCnt,
		CoinbaseAddr:     config.CoinbaseAddr,
		Workers:          make(map[common.Address]*Worker),
		coinbaseInterval: Height(config.CoinbaseInterval),
		submitAdvance:    Height(config.SubmitAdvance),
		exitCh:           make(chan struct{}),
		urlList:          config.WsUrl,
		isEffective:      sync.Once{},
	}

	explorer := RpcExplorer{
		Client:     new(rpc.Client),
		Sub:        new(rpc.ClientSubscription),
		HeaderCh:   make(chan types.HeaderShort),
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
	log.Info("waiting for next mining epoch")
	return &miner, nil
}

func (m *Miner) updateWS() {
	if m.scanner.IsUpdating() {
		return
	}
	m.scanner.close()
	exp, ok := m.scanner.explorer.(*RpcExplorer)
	if !ok {
		Fatalf("Full node miner disconnected from Avis Network more than %v", NewHeaderTimeoutDuration.String())
		return
	}
	//clean old rpc explorer and new
	oldURL := exp.WsUrl
	go func() {
		if exp.Sub == nil || exp.Client == nil {
			return
		}
		exp.Sub.Unsubscribe()
		exp.Client.Close()
	}()
	remoteExp := RpcExplorer{
		Client:     new(rpc.Client),
		Sub:        new(rpc.ClientSubscription),
		HeaderCh:   make(chan types.HeaderShort),
		rpcTimeOut: RPCTIMEOUT,
		WsUrl:      oldURL,
	}
	m.scanner.explorer = &remoteExp

	//set scanner status updating
	atomic.StoreInt32(&m.scanner.updating, int32(1))
	defer func() {
		atomic.StoreInt32(&m.scanner.updating, int32(0))
	}()

	res := false
	var err error
	for i := 0; i < WSUrlTryRound; i++ {
		for _, url := range m.urlList {
			jitter := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(RetryJitterMaxDuration)
			time.Sleep(RetryWSRPCWaitDuration + time.Millisecond*time.Duration(jitter))
			if url == remoteExp.WsUrl {
				continue
			}
			remoteExp.Client, err = rpc.Dial(url)
			if err != nil {
				log.Warn("Websocket dial Evanesco node err", "err", err)
				continue
			}

			remoteExp.Sub, err = remoteExp.Client.EthSubscribe(context.Background(), remoteExp.HeaderCh, "newHeadShort")
			if err != nil {
				log.Warn("Subscribe block err", "err", err)
				continue
			}

			res = true
			remoteExp.WsUrl = url
			log.Info("Connected node WebSocket URL", "url", url)
			break
		}
		if res == true {
			//check miner address effective
			checkEffective := func() {
				defer func() {
					m.scanner.CoinbaseAddr = m.CoinbaseAddr
				}()
				evaClient := evaclient.NewClient(remoteExp.Client)
				caller, err := NewPledgeCaller(PledgeContract, evaClient)
				if err != nil {
					Fatalf("New Pledge caller error %v", err)
				}
				minerKey := m.config.MinerList[0]
				ok, coinbasePledge, err := caller.IseffectiveNew(&bind.CallOpts{Pending: false}, minerKey.Address)
				if err != nil {
					Fatalf("call pledge contract abi IseffectiveNew error %v", err)
				}
				if !ok {
					log.Error(NotEffectiveAddrError.Error(), "address", minerKey.Address.String())
				}
				emptyAddr := common.Address{}
				//coinbase address is not set on pledge, use miner address by default
				if coinbasePledge == emptyAddr {
					if m.CoinbaseAddr == emptyAddr {
						m.CoinbaseAddr = minerKey.Address
						return
					}
					return
				}
				//coinbase address is set, use pledge coinbase address by default
				if coinbasePledge != emptyAddr {
					if m.CoinbaseAddr == emptyAddr {
						m.CoinbaseAddr = coinbasePledge
						return
					}
					if m.CoinbaseAddr != coinbasePledge {
						log.Error(NotPledgeCoinbaseError.Error())
						log.Info("miner coinbase address:" + m.CoinbaseAddr.String() + ", fortress coinbase address:" + coinbasePledge.String())
						return
					}
					return
				}
			}

			m.isEffective.Do(checkEffective)
			break
		} else {
			//reset url to try this url again
			remoteExp.WsUrl = ""
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
	for _, worker := range m.Workers {
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
	if len(m.Workers) == int(m.MaxWorkerCnt) {
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

	m.Workers[minerKey.Address] = &worker
	go worker.Loop()
	worker.start()
	log.Debug("worker start")
}

func (m *Miner) CloseWorker(addr common.Address) {
	if worker, ok := m.Workers[addr]; ok {
		worker.close()
		delete(m.Workers, addr)
	}
}

func (m *Miner) StopWorker(addr common.Address) {
	if worker, ok := m.Workers[addr]; ok {
		worker.stop()
	}
}

func (m *Miner) StartWorker(addr common.Address) {
	if worker, ok := m.Workers[addr]; ok {
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
				for _, worker := range m.Workers {
					task := SetTaskMinerAddr(taskTem, worker.minerAddr)
					worker.inboundTaskCh <- &task
				}
				continue
			}
			if taskTem.Step == TASKGETCHALLENGEBLOCK {
				if worker, ok := m.Workers[taskTem.minerAddr]; ok {
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
		updating:           int32(0),
	}
}

func (m *Miner) StartScanner() {
	i := 0
	for {
		i++
		if m.scanner.IsClosed() {
			break
		}
		time.Sleep(time.Millisecond * 100)
		m.scanner.close()
		if i == 10 {
			Fatalf("start carrier scanner failed")
		}
	}
	m.scanner.exitCh = make(chan struct{})
	go m.scanner.Loop()
}
