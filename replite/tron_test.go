package main

import (
	"fmt"
	"log"
	"testing"
)

func TestGetCurBlockNumber(t *testing.T) {

	GetCurBlockNumber()
}

func TestGetTransactionsByBlockNumber(t *testing.T) {
	transactions, err := GetTransactionByBlock(46691388)
	if err != nil {
		panic(fmt.Sprintf("run GetTransactionByBlock failed:blockNumber:%d,failedinfo:%s", 46691388, err.Error()))
	}
	log.Printf("blocknumber %d一共有%d笔交易", 46691388, len(transactions.List))
	if len(transactions.List) <= 0 {
		return
	}
	for i := 0; i < len(transactions.List); i++ {
		fmt.Printf("%v\n", transactions.List[i])
	}
}
