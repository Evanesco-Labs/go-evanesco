package evaclient

import (
	"context"
	"testing"
	"time"
)

func TestEVAWhiteNoiseClient(t *testing.T) {
	bootstrapAddress := "/ip4/10.0.0.197/tcp/3331/p2p/12D3KooWAz8hqZXf9y6mpbdh4XquXMB2yak3k584CKrKjnNEb89V"
	whitenoiseid := "08T22uD6H7V1MZmVNeAKU9SX46UtytmumpW99W8nZTyaB"
	urlScheme := "wn://"
	url := urlScheme + bootstrapAddress + "#" + whitenoiseid
	cli, err := Dial(url)
	if err != nil {
		t.Fatal(err)
	}

	//wait connection generation
	time.Sleep(time.Second)

	//rpc get block numbers
	ctx := context.Background()
	num, err := cli.BlockNumber(ctx)
	if err != nil {
		t.Fatal(err)
	}
	println(num)

	//rpc get chainID
	chainID, err := cli.ChainID(ctx)
	if err != nil {
		t.Fatal(err)
	}
	println(chainID.String())
}
