package core


type MatchAccountNode struct{
	Address string
	Balance string
	Relative int
}

type emptyTransaction string


type TransactionNode struct{
	emptyTransaction 
	From string
	To string
	Balance string
	date int64
}

type RouteNode struct{
	PreNode *RouteNode
	NextNode *RouteNode
	Info emptyTransaction
	RootTransaction emptyTransaction
	RelativeTransaction []emptyTransaction
}


func Match([]*TransactionNode)