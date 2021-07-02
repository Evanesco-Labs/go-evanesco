package rpc

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func Test_Parse(t *testing.T) {
	bootAddressExpect := "/ip4/127.0.0.1/tcp/3331/p2p/12D3KooWEyoppNCUx8Yx66oV9fJnriXwCcXwDDUA2kj6vnc6iDEp"
	whiteNoiseIDExpect := "0GS2hZQpGN8kZ87Y54dpU3D3PYkjyL9N9gNK6Q9dgvF9k"
	wnPrefix := "wn://"
	urlText := wnPrefix + bootAddressExpect + "#" + whiteNoiseIDExpect
	u, err := url.Parse(urlText)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, u.Scheme, "wn")

	bootAddress, whitenoiseID, err := ParseWhiteNoiseEndpoint(urlText)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, bootAddress, bootAddressExpect)
	assert.Equal(t, whitenoiseID, whiteNoiseIDExpect)
}
