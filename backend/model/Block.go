package Model


type Block struct{
	Blockchain *Blockchain
	header *BlockHeader
	Uncles []*BlockHeader
	Transactions Transactions
	
	hash string
	size  int
}