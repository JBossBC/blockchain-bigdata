


package Model


type BlockBody interface{
	GetTransactions()Transactions
}


type BaseBody struct{
	Transactions Transactions
}

func(base *BaseBody)GetTransactions()Transactions{
	return base.Transactions
}