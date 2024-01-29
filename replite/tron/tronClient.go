package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"strconv"
// 	"sync"
// )

// var tronClient = http.DefaultClient
// var tronOnce sync.Once

// var searchBlockByNumberURL = "https://api.trongrid.io/wallet/getblockbynum"

// func GetTransactionByBlock(number int) error {
// 	valueByt, err := json.Marshal(strconv.FormatInt(int64(number), 10))
// 	if err != nil {
// 		log.Printf("serialize the blocknumber %d error:%s", number, err.Error())
// 		return err
// 	}
// 	resp, err := tronClient.Post(searchBlockByNumberURL, "application/json;charset=utf-8", bytes.NewReader(valueByt))
// 	if err != nil {
// 		log.Printf("get transactions from %d block failed:%s", number, err.Error())
// 		return err
// 	}
// 	respBytes, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Printf("read the response from search trasactions from %d block method failed:%s", number, err.Error())
// 		return err
// 	}
// 	fmt.Println(string(respBytes))
// 	return nil
// }
