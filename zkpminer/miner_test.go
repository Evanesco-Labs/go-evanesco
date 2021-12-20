package zkpminer

import (
	"github.com/Evanesco-Labs/go-evanesco/common"
	"github.com/Evanesco-Labs/go-evanesco/consensus/ethash"
	"github.com/Evanesco-Labs/go-evanesco/core"
	"github.com/Evanesco-Labs/go-evanesco/core/rawdb"
	"github.com/Evanesco-Labs/go-evanesco/log"
	"github.com/Evanesco-Labs/go-evanesco/params"
	"github.com/Evanesco-Labs/go-evanesco/zkpminer/keypair"
	"testing"
	"time"
)

var (
	testdb     = rawdb.NewMemoryDatabase()
	gspec      = &core.Genesis{Config: params.TestChainConfig}
	genesis    = gspec.MustCommit(testdb)
	blocks, _  = core.GenerateChain(params.TestChainConfig, genesis, ethash.NewFaker(), testdb, 100, nil)
	testURL    = "ws://127.0.0.1:8549"
	testPKPath = "./provekeytest.txt"
)

func DefaultTestConfig() Config {
	return Config{
		MinerList:        make([]keypair.Key, 0),
		MaxWorkerCnt:     10,
		MaxTaskCnt:       10,
		CoinbaseInterval: uint64(5),
		SubmitAdvance:    SUBMITADVANCE,
		CoinbaseAddr:     common.Address{},
		WsUrl:            []string{},
		RpcTimeout:       RPCTIMEOUT,
		PkPath:           "./QmQL4k1hKYiW3SDtMREjnrah1PBsak1VE3VgEqTyoDckz9",
	}
}

func TestMiner(t *testing.T) {
	minerList := make([]keypair.Key, 0)
	{
		_, sk := keypair.GenerateKeyPair()
		key, err := keypair.NewKey(sk.PrivateKey)
		if err != nil {
			t.Fatal(err)
		}
		minerList = append(minerList, key)
	}

	_, coinbaseSk := keypair.GenerateKeyPair()
	coinbaseKey, err := keypair.NewKey(coinbaseSk.PrivateKey)
	if err != nil {
		t.Fatal(err)
	}

	config := DefaultTestConfig()
	config.Customize(minerList, coinbaseKey.Address, []string{testURL}, testPKPath)

	miner, err := NewMiner(config)
	if err != nil {
		t.Fatal(err)
	}
	log.Debug("coinbase address:", "coinbase address", miner.scanner.CoinbaseAddr)

	//add another worker
	time.Sleep(time.Second * 20)

	_, sk := keypair.GenerateKeyPair()
	key, err := keypair.NewKey(sk.PrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	miner.NewWorker(key)

	select {}
}
