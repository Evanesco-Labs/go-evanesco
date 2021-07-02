package rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Evanesco-Labs/WhiteNoise/sdk"
	"github.com/Evanesco-Labs/WhiteNoise/secure"
	"github.com/ethereum/go-ethereum/log"
	"github.com/libp2p/go-libp2p-core/peer"
	"io"
	"math/rand"
	"strings"
	"time"
)

const DefaultWhiteNoiseKeyType = 0
const MaxPeersList = 10

func DialWhiteNoise(ctx context.Context, endpoint string) (*Client, error) {
	bootstrapAddr, whitenoiseID, err := ParseWhiteNoiseEndpoint(endpoint)
	if err != nil {
		return nil, err
	}
	sdk.BootStrapPeers = bootstrapAddr
	wnClient, err := sdk.NewOneTimeClient(ctx, DefaultWhiteNoiseKeyType)
	if err != nil {
		return nil, err
	}

	peers, err := wnClient.GetMainNetPeers(MaxPeersList)
	if err != nil {
		return nil, err
	}

	//choose random to register
	if len(peers) == 0 {
		return nil, errors.New("no peers exist")
	}

	index := rand.New(rand.NewSource(time.Now().UnixNano())).Int() % len(peers)
	entry := peers[index]

	if err := wnClient.Register(entry); err != nil {
		return nil, err
	}

	client, err := newClient(ctx, func(ctx context.Context) (ServerCodec, error) {
		conn, sessionid, err := wnClient.Dial(whitenoiseID)
		if err != nil {
			return nil, err
		}
		return JsonWhiteNoise{
			Client:     wnClient,
			SessioinID: sessionid,
			Proxy:      entry,
			Conn:       conn,
			Closed:     make(chan interface{}),
		}, nil
	})
	if err != nil {
		return nil, err
	}
	client.isWhiteNoise = true
	return client, nil
}

func ParseWhiteNoiseEndpoint(url string) (string, string, error) {
	s := strings.Split(url, "//")
	if len(s) != 2 {
		return "", "", errors.New("whitenoise rpc url err")
	}
	t := strings.Split(s[1], "#")
	if len(t) != 2 {
		return "", "", errors.New("whitenoise rpc url err")
	}
	return t[0], t[1], nil
}

type WhiteNoiseConn struct {
	sdk.SecureConnection
}

func (conn *WhiteNoiseConn) WriteJson(v interface{}) error {
	return json.NewEncoder(conn).Encode(v)
}

type JsonWhiteNoise struct {
	Client     sdk.Client
	SessioinID string
	Proxy      peer.ID
	Conn       sdk.SecureConnection
	Closed     chan interface{}
}

func (wn JsonWhiteNoise) readBatch() (messages []*jsonrpcMessage, batch bool, err error) {
	// Decode the next JSON object in the input stream.
	// This verifies basic syntax, etc.
	log.Info("whitenoise read batch")
	rawmsg := make([]byte, 0)
	for {
		rawmsg, err = secure.ReadPayload(wn.Conn)
		if err != nil {
			if err == io.EOF {
				continue
			}
			return nil, false, errors.New("secure read payload err; " + err.Error())
		} else {
			break
		}
	}

	fmt.Println(rawmsg)
	messages, batch = parseMessage(rawmsg)
	for i, msg := range messages {
		if msg == nil {
			// Message is JSON 'null'. Replace with zero value so it
			// will be treated like any other invalid message.
			messages[i] = new(jsonrpcMessage)
		}
	}

	return messages, batch, nil
}

func (wn JsonWhiteNoise) close() {
	wn.Client.UnRegister()
}

func (j JsonWhiteNoise) writeJSON(ctx context.Context, v interface{}) error {
	b := make([]byte, 0)
	buf := bytes.NewBuffer(b)
	err := json.NewEncoder(buf).Encode(v)
	if err != nil {
		return err
	}
	msg := buf.Bytes()
	fmt.Println(msg)
	encoded := secure.EncodePayload(msg)
	return j.Client.SendMessage(encoded, j.SessioinID)
}

func (wn JsonWhiteNoise) closed() <-chan interface{} {
	return wn.Closed
}

func (wn JsonWhiteNoise) remoteAddr() string {
	return wn.Conn.RemoteWhiteNoiseID()
}
