package node

import (
	"context"
	"errors"
	"github.com/Evanesco-Labs/WhiteNoise/common/account"
	"github.com/Evanesco-Labs/WhiteNoise/sdk"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"math/rand"
	"time"
)

type WhiteNoiseServer struct {
	log          log.Logger
	host         string
	port         int
	whitenoiseID string
	client       sdk.Client
	rpcServer    *rpc.Server
	account      *account.Account
	bootstrap    string
	proxy        string
}

func NewWhiteNoiseServer(log log.Logger, account *account.Account) *WhiteNoiseServer {
	return &WhiteNoiseServer{
		log:          log,
		host:         "localhost",
		port:         0,
		whitenoiseID: "",
		client:       nil,
		rpcServer:    rpc.NewServer(),
		account:      account,
		proxy:        "",
	}
}

func (wnServer *WhiteNoiseServer) RegisterName(name string, receiver interface{}) error {
	return wnServer.rpcServer.RegisterName(name, receiver)
}

func (wnServer *WhiteNoiseServer) start(ctx context.Context, bootstrap string) error {
	log.Info("starting whitenoise")
	wnServer.bootstrap = bootstrap
	sdk.BootStrapPeers = wnServer.bootstrap
	client, err := sdk.NewClient(ctx, wnServer.account)
	if err != nil {
		return err
	}

	wnServer.client = client
	wnServer.whitenoiseID = client.GetWhiteNoiseID()

	log.Info("whitenosie id: " + wnServer.whitenoiseID)
	peers, err := client.GetMainNetPeers(10)
	if err != nil {
		return err
	}

	//choose random to register
	if len(peers) == 0 {
		return errors.New("no peers exist")
	}

	index := rand.New(rand.NewSource(time.Now().UnixNano())).Int() % len(peers)
	entry := peers[index]

	if err := client.Register(entry); err != nil {
		return err
	}

	wnServer.proxy = entry.Pretty()

	err = wnServer.client.EventBus().Subscribe(sdk.GetCircuitTopic, func(sessionID string) {
		log.Info("get whitenosie circuit")
		conn, ok := wnServer.client.GetCircuit(sessionID)
		if !ok {
			return
		}
		codec := rpc.JsonWhiteNoise{
			Client:     wnServer.client,
			SessioinID: sessionID,
			Proxy:      entry,
			Conn:       conn,
			Closed:     make(chan interface{}),
		}
		log.Info("rpcServer serve conn")
		go wnServer.rpcServer.ServeCodec(codec, 0)
	})

	return err
}
