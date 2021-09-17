package zkpminer

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"sync"
	"sync/atomic"
	"time"
)

var zero = new(big.Int).SetInt64(int64(0))

var NewHeaderTimeoutDuration = time.Minute

var (
	InvalidTaskStepErr = errors.New("invalid task step")
)

var (
	NotCliqueConsensusError = errors.New("no clique engine, invalid Evanesco node")
	NotEffectiveAddrError   = errors.New("miner address not staked")
	ZKPProofVerifyError     = errors.New("ZKP proof verify failed")
	NotPledgeCoinbaseError  = errors.New("coinbase address conflict, check the coinbase address setting in Fortress")
)

type Height uint64

type Explorer interface {
	GetHeaderChan() chan *types.Header
	GetHeaderByNum(num uint64) *types.Header
}

type RpcExplorer struct {
	Client     *rpc.Client
	Sub        ethereum.Subscription
	HeaderCh   chan *types.Header
	rpcTimeOut time.Duration
	WsUrl      string
}

func (r *RpcExplorer) GetHeaderChan() chan *types.Header {
	return r.HeaderCh
}

func (r *RpcExplorer) GetHeaderByNum(num uint64) *types.Header {
	ctx, cancel := context.WithTimeout(context.Background(), r.rpcTimeOut)
	defer cancel()
	var head *types.Header
	err := r.Client.CallContext(ctx, &head, "eth_getBlockByNumber", toBlockNumArg(new(big.Int).SetUint64(num)), false)
	if err != nil {
		log.Error("rpc call eth_getBlockByNumber error", "err", err)
		return nil
	}
	return head
}

type LocalExplorer struct {
	Backend
	headerCh chan *types.Header
}

func (l *LocalExplorer) GetHeaderChan() chan *types.Header {
	return l.headerCh
}

func (l *LocalExplorer) GetHeaderByNum(num uint64) *types.Header {
	return l.BlockChain().GetHeaderByNumber(num)
}

type Scanner struct {
	mu                 sync.RWMutex
	miner              *Miner
	CoinbaseAddr       common.Address
	BestScore          *big.Int
	LastBlockHeight    Height
	CoinbaseInterval   Height
	LastCoinbaseHeight Height
	explorer           Explorer
	taskWait           map[Height][]*Task //single concurrent
	inboundTaskCh      chan *Task         //channel to receive tasks from worker
	outboundTaskCh     chan *Task         //channel to send challenged task to miner
	running            int32
	exitCh             chan struct{}
}

func (s *Scanner) NewTask(h *types.Header) Task {
	return Task{
		Step:             TASKSTART,
		CoinbaseAddr:     s.CoinbaseAddr,
		lastCoinBaseHash: h.Hash(),
		challengeIndex:   Height(uint64(0)),
		lottery: &types.Lottery{
			CoinbaseAddr: s.CoinbaseAddr,
		},
	}
}

func (s *Scanner) close() {
	defer func() {
		if recover() != nil {
		}
	}()
	close(s.exitCh)
	atomic.StoreInt32(&s.running, int32(0))
}

func (s *Scanner) Loop() {
	headerCh := s.explorer.GetHeaderChan()
	timer := time.NewTimer(NewHeaderTimeoutDuration)
	for {
		select {
		case <-s.exitCh:
			timer.Stop()
			return
		case <-timer.C:
			log.Warn("receive new header timeout")
			timer.Stop()
			s.miner.updateWS()
		case header := <-headerCh:
			timer.Reset(NewHeaderTimeoutDuration)
			log.Debug("best score:", "score", s.BestScore)
			height := Height(header.Number.Uint64())
			//index := height - s.LastCoinbaseHeight
			index := Height(new(big.Int).Mod(header.Number, new(big.Int).SetUint64(uint64(s.CoinbaseInterval))).Uint64())
			log.Debug("chain status", "height", height, "index", index)

			s.LastBlockHeight = height
			if s.IfCoinBase(header) {
				if _, ok := s.miner.Workers[header.BestLottery.MinerAddr]; ok {
					log.Info("Congratulations you got the best score!", "height", header.Number.Uint64())
				}

				log.Info("start new mining epoch")
				task := s.NewTask(header)
				s.outboundTaskCh <- &task
				s.LastCoinbaseHeight = height
				s.CleanStatus()
			}

			if taskList, ok := s.taskWait[index]; ok {
				//add challengeHeader and send tasks to miner task channel
				for _, task := range taskList {
					if task.Step != TASKWAITCHALLENGEBLOCK {
						log.Error(InvalidTaskStepErr.Error())
						continue
					}
					task.SetHeader(header)
					s.outboundTaskCh <- task
				}
				delete(s.taskWait, index)
				continue
			}

		case task := <-s.inboundTaskCh:
			if task.Step == TASKWAITCHALLENGEBLOCK {
				if taskList, ok := s.taskWait[task.challengeIndex]; ok {
					taskList = append(taskList, task)
					s.taskWait[task.challengeIndex] = taskList
					continue
				}
				taskList := []*Task{task}
				s.taskWait[task.challengeIndex] = taskList
				continue
			}
			if task.Step == TASKPROBLEMSOLVED {
				taskScore := task.lottery.Score()
				log.Debug("get solved score:", "score", taskScore)
				if taskScore.Cmp(s.BestScore) != 1 {
					log.Debug("less than best", "score", s.BestScore)
					continue
				}
				s.BestScore = taskScore
				//todo:abort lottery if exceed deadline
				go func() {
					s.Submit(task)
					task.Step = TASKSUBMITTED
					s.outboundTaskCh <- task
					log.Info("waiting for next mining epoch", "time duration (second)", (uint64(s.LastCoinbaseHeight)+types.CoinBaseInterval-uint64(s.LastBlockHeight))*6)
				}()
			}
		}
	}
}

func (s *Scanner) IfCoinBase(h *types.Header) bool {
	return new(big.Int).Mod(h.Number, new(big.Int).SetUint64(uint64(s.CoinbaseInterval))).Cmp(zero) == 0
}

//todo: improve robustness, add some retries
func (s *Scanner) GetHeader(height Height) (*types.Header, error) {
	header := s.explorer.GetHeaderByNum(uint64(height))
	if header == nil {
		return header, fmt.Errorf("get header failed, height: %v", height)
	}
	return header, nil
}

func (s *Scanner) Submit(task *Task) {
	// Submit check if the lottery has the best score
	score := (*hexutil.Big)(task.lottery.Score())
	log.Info("submiting work",
		"\nminer address", task.lottery.MinerAddr,
		"\ncoinbase address", task.lottery.CoinbaseAddr,
		"\nscore", score.String(),
	)

	if localExp, ok := s.explorer.(*LocalExplorer); ok {
		err := localExp.EventMux().Post(core.NewSolvedLotteryEvent{Lot: types.LotterySubmit{
			Lottery:   *task.lottery,
			Signature: task.signature,
		},
		})
		if err != nil {
			log.Error("submit with local explorer", "err", err)
			s.miner.Close()
		}
		return
	}

	if rpcExp, ok := s.explorer.(*RpcExplorer); ok {
		ctx, cancel := context.WithTimeout(context.Background(), RPCTIMEOUT)
		defer cancel()
		submit := types.LotterySubmit{
			Lottery:   *task.lottery,
			Signature: task.signature,
		}

		err := rpcExp.Client.CallContext(ctx, nil, "eth_lotterySubmit", submit)
		if err != nil {
			log.Error("submit work error", "err", err)
			if err.Error() == NotEffectiveAddrError.Error() {
				Fatalf("Miner address not effective")
			}
			if err.Error() == ZKPProofVerifyError.Error() {
				Fatalf("ZKP proof verify failed, please check miner setting and config")
			}
			if err.Error() == NotPledgeCoinbaseError.Error() {
				Fatalf(NotPledgeCoinbaseError.Error())
			}
			if err.Error() != NotEffectiveAddrError.Error() && err.Error() != ZKPProofVerifyError.Error() && err.Error() != NotPledgeCoinbaseError.Error() {
				log.Info("try to connect another node")
				s.miner.updateWS()
				//todo: retry submit after updateWs success
			}
		}
	}
}

func (s *Scanner) CleanStatus() {
	s.taskWait = make(map[Height][]*Task)
	s.BestScore = zero
}

func (s *Scanner) IsRunning() bool {
	if s.running == int32(1) {
		return true
	}
	return false
}

func (s *Scanner) IsUpdating() bool {
	if s.running == int32(2) {
		return true
	} else {
		return false
	}
}

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	pending := big.NewInt(-1)
	if number.Cmp(pending) == 0 {
		return "pending"
	}
	return hexutil.EncodeBig(number)
}
