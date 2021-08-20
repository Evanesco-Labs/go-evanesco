package types

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"time"
)

var (
	max256    = new(big.Int).Exp(big.NewInt(2), big.NewInt(256), big.NewInt(0))
	keccak256 = crypto.Keccak256
)

const (
	ResetTickerDuration = time.Millisecond * 10
	CoinBaseInterval    = uint64(100)
	SubmitAdvance       = uint64(20)
	RewardAmount        = uint64(500)
)

type ZKPReward struct {
	CoinbaseAddr common.Address `json:"coinbase_addr"`
	Score        [32]byte       `json:"score"`
}

type LotterySubmit struct {
	Lottery   `json:"lottery"`
	Signature [65]byte `json:"signature"`
}

type Lottery struct {
	CoinbaseAddr        common.Address `json:"coinbase_addr"`
	MinerAddr           common.Address `json:"miner_addr"`            //20 bytes
	ChallengeHeaderHash [32]byte       `json:"challenge_header_hash"` //challenge block header Hash
	Index               [32]byte       `json:"index"`
	MimcHash            []byte         `json:"mimc_hash"` //32 bytes
	ZkpProof            []byte         `json:"zkp_proof"`
	VrfProof            []byte         `json:"vrf_proof"`
}

func (l *Lottery) SetMinerAddr(addr common.Address) {
	l.MinerAddr = addr
}

func (l *Lottery) SetVrfProof(proof []byte) {
	l.VrfProof = proof
}

func (l *Lottery) SetZKPProof(proof []byte) {
	l.ZkpProof = proof
}

func (l *Lottery) Serialize() ([]byte, error) {
	return json.Marshal(l)
}

func (l *Lottery) Deserialize(data []byte) error {
	return json.Unmarshal(data, l)
}

func (l *Lottery) ScoreBytes() [32]byte {
	var res [32]byte
	b := append(l.MinerAddr.Bytes(), l.MimcHash...)
	b = xor(keccak256(b), l.ChallengeHeaderHash[:])
	copy(res[:], b)
	return res
}

func (l *Lottery) Score() *big.Int {
	b := l.ScoreBytes()
	return new(big.Int).SetBytes(b[:])
}

func IfPassDiff(score []byte, diff *big.Int) bool {
	target := new(big.Int).Div(max256, diff)
	if new(big.Int).SetBytes(score).Cmp(target) > 0 {
		return false
	} else {
		return true
	}
}

func xor(one, other []byte) (xor []byte) {
	if len(one) != len(other) {
		return nil
	}
	xor = make([]byte, len(one))
	for i := 0; i < len(one); i++ {
		xor[i] = one[i] ^ other[i]
	}
	return xor
}

func (lp *LotterySubmit) Hash() common.Hash {
	b, err := json.Marshal(lp)
	if err != nil {
		return common.Hash{}
	}
	return crypto.Keccak256Hash(b)
}

func (h *Header) IsZKPRewardBlock() bool {
	zero := new(big.Int).SetUint64(uint64(0))
	return new(big.Int).Mod(h.Number, new(big.Int).SetUint64(CoinBaseInterval)).Cmp(zero) == 0
}