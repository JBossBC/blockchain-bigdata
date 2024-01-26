package Model


type BlockHeader struct{
	ParentHash string
	UncleHash string
	Coinbase Address
	TxHash string
}