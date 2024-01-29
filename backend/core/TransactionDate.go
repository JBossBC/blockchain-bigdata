package core

import (
	"sort"
	"sync"
	"unsafe"
)

type transactionDateSearch []*TransactionNode

func (nodes transactionDateSearch) Len() int {
	return len(nodes)
}

func (nodes transactionDateSearch) Less(i, j int) bool {
	return nodes[i].Date < nodes[j].Date
}

func (nodes transactionDateSearch) Swap(i, j int) {
	nodes[i], nodes[j] = nodes[j], nodes[i]
}

var sortOnce = make(map[uintptr]any, 0)
var mutex sync.Mutex

func _onceSort(nodes *transactionDateSearch) {
	mutex.Lock()
	defer mutex.Unlock()
	ptr := uintptr(unsafe.Pointer(nodes))
	if _, ok := sortOnce[ptr]; ok {
		return
	}
	sort.Sort(*nodes)
	sortOnce[ptr] = struct{}{}
}
func putSortFlag(ptr uintptr) {
	mutex.Lock()
	defer mutex.Unlock()
	sortOnce[ptr] = struct{}{}
}
func (nodes transactionDateSearch) Sort() {
	sort.Sort(nodes)
	putSortFlag(uintptr(unsafe.Pointer(&nodes)))
}

func (nodes transactionDateSearch) rangeQuery(date, mistake int64) []*TransactionNode {
	_onceSort((*transactionDateSearch)(&nodes))
	result := []*TransactionNode{}
	start := date - mistake
	end := date + mistake
	lowerBound := lowerBound(nodes, start)
	upperBound := upperBound(nodes, end)
	for i := lowerBound; i < upperBound; i++ {
		if nodes[i].Date >= start && nodes[i].Date <= end {
			result = append(result, nodes[i])
		}
	}

	return result
}
func lowerBound(nodes transactionDateSearch, target int64) int {
	left, right := 0, len(nodes)

	for left < right {
		mid := left + (right-left)/2
		if nodes[mid].Date < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func upperBound(nodes transactionDateSearch, target int64) int {
	left, right := 0, len(nodes)

	for left < right {
		mid := left + (right-left)/2
		if nodes[mid].Date <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

type dateRange struct {
	From int64
	To   int64
}

func (nodes transactionDateSearch) mergeRanges(mistake int64) []*dateRange {
	if len(nodes) == 0 {
		return nil
	}
	_onceSort((*transactionDateSearch)(&nodes))
	mergedRanges := []*dateRange{&dateRange{nodes[0].Date - mistake, nodes[0].Date + mistake}}
	for i := 1; i < len(nodes); i++ {
		currentNode := nodes[i]
		lastMergedRange := mergedRanges[len(mergedRanges)-1]

		if currentNode.Date-mistake <= lastMergedRange.To {
			// Merge overlapping ranges
			lastMergedRange.To = max(lastMergedRange.To, currentNode.Date+mistake)
		} else {
			// Add a new range
			mergedRanges = append(mergedRanges, &dateRange{currentNode.Date - mistake, currentNode.Date + mistake})
		}
	}

	return mergedRanges
}
