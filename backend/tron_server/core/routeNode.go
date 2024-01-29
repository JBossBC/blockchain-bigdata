package core

type routeNode struct {
	PreNode             string
	NextNode            string
	Info                emptyTransaction
	RelativeTransaction []*emptyTransaction
}

func newRouteNodeFromTransactionNode(cur *TransactionNode, uncles []*emptyTransaction) *routeNode {
	return &routeNode{
		PreNode:             cur.From,
		NextNode:            cur.To,
		Info:                cur.emptyTransaction,
		RelativeTransaction: uncles,
	}
}
