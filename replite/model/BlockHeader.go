package Model

import (
	"math/big"
	"time"
)

type BlockHeader interface {
	GetParentHash() string
	GetBlockOnwer() Address
	// 交易集合Trie的根hash
	GetTxsHash() string
	// 有关于共识机制的区分
	// GetDifficulty()string
	GetNumber() *big.Int
	GetTotalFee() Fee
	GetTime() time.Time
	GetExtraData() []byte
	// nonce?
	// GetNonce()BlockNonce
}

type BaseHeader struct {
	ParentHash string
	Coinbase   Address
	TxsHash    string
	Number     *big.Int
	Time       time.Time
	ExtraData  []byte
}

func (base *BaseHeader) GetParentHash() string {
	return base.ParentHash
}
func (base *BaseHeader) GetTxsHash() string {
	return base.TxsHash
}

func (base *BaseHeader) GetNumber() *big.Int {
	return base.Number
}
func (base *BaseHeader) GetTotalFee() Fee {
	return nil
}
func (base *BaseHeader) GetBlockOnwer() Address {
	return base.Coinbase
}
func (base *BaseHeader) GetExtraData() []byte {
	return base.ExtraData
}
func (base *BaseHeader) GetTime() time.Time {
	return base.Time
}
