package zkpminer

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/evaclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"io"
	"os"
	"runtime"
)

var PledgeContract = common.HexToAddress("0x9Cf113d5c6f3aA616D690e7B2acd220f3b63E6ed")

func Iseffective(miner common.Address, server *rpc.Server) (bool, common.Address) {
	client := evaclient.NewClient(rpc.DialInProc(server))
	defer client.Close()
	caller, err := NewPledgeCaller(PledgeContract, client)
	if err != nil {
		log.Error("New Pledge Contract error", "contract address", PledgeContract.String(), "err", err.Error())
		return false, common.Address{}
	}
	ok, coinbaseAddr, err := caller.IseffectiveNew(&bind.CallOpts{Pending: false}, miner)
	if err != nil {
		log.Error("Iseffective error", "err", err)
		return false, common.Address{}
	}
	return ok, coinbaseAddr
}

func Fatalf(format string, args ...interface{}) {
	w := io.MultiWriter(os.Stdout, os.Stderr)
	if runtime.GOOS == "windows" {
		// The SameFile check below doesn't work on Windows.
		// stdout is unlikely to get redirected though, so just print there.
		w = os.Stdout
	} else {
		outf, _ := os.Stdout.Stat()
		errf, _ := os.Stderr.Stat()
		if outf != nil && errf != nil && os.SameFile(outf, errf) {
			w = os.Stderr
		}
	}
	fmt.Fprintf(w, "Fatal: "+format+"\n", args...)
	os.Exit(1)
}
