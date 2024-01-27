package Model

type Block interface {
	GetBlockChain() *BlockChain
	GetBlockHeader() *BlockHeader
	GetTransactions() Transactions
	GetSize() int64
	GetHash() string
}

type BaseBlock struct {
	Blockchain BlockChain
	Header     BlockHeader
	//特定区块含有
	// Uncles       []*BlockHeader
	Transactions Transactions
	Hash string
	Size int64
}

func (base *BaseBlock) GetBlockChain() BlockChain {
	return base.Blockchain
}

func (base *BaseBlock) GetBlockHeader() BlockHeader {
	return base.Header
}
func (base *BaseBlock) GetTransactions() Transactions {
	return base.Transactions
}
func (base *BaseBlock) GetSize() int64 {
	return base.Size
}
func (base *BaseBlock) GetHash() string {
	return base.Hash
}
