package core

import "unsafe"

type matchAccountNode struct {
	Address  string
	Balance  string
	Relative int
}

type emptyTransaction string

type TransactionNode struct {
	emptyTransaction
	From    string
	To      string
	Balance int64
	Date    int64
}

func EnconomyMatch(transactions []*TransactionNode, mistake int64, relative int64, maxDepth int64) (accounts [2]map[string]*matchAccountNode, crimePath map[string][]*routeNode) {
	var gradMatch func(preBalance int64, tos []*TransactionNode) [][]*TransactionNode
	gradMatch = func(preBalance int64, tos []*TransactionNode) [][]*TransactionNode {

	}

	var enconomyMatch func(transactions []*TransactionNode, depth int64, walk []*routeNode)
	enconomyMatch = func(transactions []*TransactionNode, depth int64, walk []*routeNode) {
		sortTransactions := transactionDateSearch(transactions)
		sortTransactions.Sort()
		var accToTransactions map[string][]*TransactionNode = make(map[string][]*TransactionNode, 0)
		for i := 0; i < len(transactions); i++ {
			var tmp = transactions[i]
			if accToTransactions[tmp.To] == nil {
				accToTransactions[tmp.To] = make([]*TransactionNode, 0)
			}
			accToTransactions[tmp.To] = append(accToTransactions[tmp.To], tmp)
		}
		for k, v := range accToTransactions {
			var sortV = transactionDateSearch(v)
			putSortFlag(uintptr(unsafe.Pointer(&sortV)))
			remains := batchSearch(k, sortV.mergeRanges(mistake))
			var sortRemain = transactionDateSearch(remains)
			sortRemain.Sort()
			for i := 0; i < len(v); i++ {
				var trans = v[i]
				walk = append(walk)
				RouteMemory(walk, gradMatch(trans.Balance, sortRemain.rangeQuery(trans.Date, mistake)))
			}
		}
	}
	enconomyMatch(transactions, 0, []*routeNode{})
}
func RouteMemory(walk []*routeNode, nexts [][]*TransactionNode) {

}

func batchSearch(address string, daterange []*dateRange) []*TransactionNode {

	return nil
}
