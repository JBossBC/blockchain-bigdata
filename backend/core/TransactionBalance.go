package core

import "sort"

type transactionBalanceSort []*TransactionNode

func (nodes transactionBalanceSort) Len() int {
	return len(nodes)
}

func (nodes transactionBalanceSort) Less(i, j int) bool {
	return nodes[i].Balance < nodes[j].Balance
}

func (nodes transactionBalanceSort) Swap(i, j int) {
	nodes[i], nodes[j] = nodes[j], nodes[i]
}

func (nodes transactionBalanceSort) Sort() {
	sort.Sort(nodes)
}
