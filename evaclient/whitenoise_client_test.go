package evaclient

import (
	"context"
	"testing"
	"time"
)

func TestNormalClient(t *testing.T) {
	whitenoiseid := "0BnQeww21evjeccPvfsDSPY4KZ5TWDTaznqVghPg2D9E6"
	url := "wn:///ip4/10.0.0.197/tcp/3331/p2p/12D3KooWAz8hqZXf9y6mpbdh4XquXMB2yak3k584CKrKjnNEb89V#"
	url = url + whitenoiseid
	cli, err := Dial(url)
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Second * 3)

	ctx := context.Background()
	num, err := cli.BlockNumber(ctx)
	if err != nil {
		t.Fatal(err)
	}
	println(num)
}
