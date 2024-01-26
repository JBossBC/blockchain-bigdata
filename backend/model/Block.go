package Model


type Block struct{
	Blockchain *BlockChain
	header *BlockHeader
	Uncles []*BlockHeader
	Transactions Transactions
	
	hash string
	size  int
}