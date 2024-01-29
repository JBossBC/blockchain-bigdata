package core

import "unsafe"

type matchAccountNode struct {
	Address  string
	Balance  int64
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

func EnconomyMatch(transactions []*TransactionNode, mistake int64, relative int64, maxDepth int64) (accounts map[string]*matchAccountNode) {
	var gradMatch func(preBalance int64, tos []*TransactionNode) [][]*TransactionNode
	accounts = make(map[string]*matchAccountNode)
	gradMatch = func(preBalance int64, tos []*TransactionNode) [][]*TransactionNode {
		if len(tos) <= 0 {
			return nil
		}
		transactionBalanceSort(tos).Sort()
		var res [][]*TransactionNode
		for i := 0; i < len(tos)-1; i++ {
			low, high := i+1, len(tos)-1
			for low < high {
				sum := tos[i].Balance + tos[low].Balance + tos[high].Balance
				if sum == preBalance {
					res = append(res, []*TransactionNode{tos[i], tos[low], tos[high]})
					low++
					high--
				} else if sum < preBalance {
					low++
				} else {
					high--
				}
			}
		}
		return res
	}
	var enconomyMatch func(transactions []*TransactionNode, depth int64)
	enconomyMatch = func(transactions []*TransactionNode, depth int64) {
		if depth > maxDepth {
			return
		}
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
		var nodes = make([]*TransactionNode, 0)
		for k, v := range accToTransactions {
			var sortV = transactionDateSearch(v)
			putSortFlag(uintptr(unsafe.Pointer(&sortV)))
			remains := batchSearch(k, sortV.mergeRanges(mistake))
			var sortRemain = transactionDateSearch(remains)
			sortRemain.Sort()
			for i := 0; i < len(v); i++ {
				var trans = v[i]
				nexts := gradMatch(trans.Balance, sortRemain.rangeQuery(trans.Date, mistake))
				for j := 0; j < len(nexts); j++ {
					var path = nexts[j]
					for k := 0; k < len(path); k++ {
						var t = path[k]
						nodes = append(nodes, t)
						if accounts[t.To] == nil {
							accounts[t.To] = &matchAccountNode{
								Address: t.To,
							}
						}
						accounts[t.To].Balance += t.Balance
						accounts[t.To].Relative++
					}
					saveRoute(path)
				}
			}
		}
		enconomyMatch(nodes, depth+1)
	}
	enconomyMatch(transactions, 0)
	return accounts
}

func saveRoute(route []*TransactionNode) {
	var emptyTrans = make([]*emptyTransaction, len(route))
	for i := 0; i < len(route); i++ {
		emptyTrans[i] = &route[i].emptyTransaction
	}
	var routes = make([]*routeNode, len(route))
	for i := 0; i < len(route); i++ {
		routes[i] = newRouteNodeFromTransactionNode(route[i], emptyTrans)
	}

}

func batchSearch(address string, daterange []*dateRange) []*TransactionNode {

	return nil
}
