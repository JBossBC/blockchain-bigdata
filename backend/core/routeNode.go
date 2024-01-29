package core

type routeNode struct {
	PreNode             *routeNode
	NextNode            *routeNode
	Info                emptyTransaction
	RootTransaction     emptyTransaction
	RelativeTransaction []emptyTransaction
}

func NewRouteNodeFromTransactionNode(preNodes *routeNode, uncles []*TransactionNode) *routeNode {

}
