package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
)

var defualtStateFileLocation = fmt.Sprintf("%s%c%s", "D:", os.PathSeparator, "tron-State.txt")

var searchCurBlockNumber = "https://apilist.tronscanapi.com/api/block"

var getBlockNumberRequest *http.Request

func GetCurBlockNumber() int64 {
	if getBlockNumberRequest == nil {
		req, err := http.NewRequest("GET", searchCurBlockNumber, nil)
		if err != nil {
			return math.MinInt64
		}
		getBlockNumberRequest = req
	}
	getBlockNumberRequest.Header = singleHeader

	resp, err := client.Do(getBlockNumberRequest)
	if err != nil {
		return math.MinInt64
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return math.MinInt64
	}
	err = json.Unmarshal(body, &CurBlockNumberResponse)
	if err != nil {
		return math.MinInt64
	}
	log.Printf("current the blocknumber is:%d", CurBlockNumberResponse.RangeTotal)
	return CurBlockNumberResponse.RangeTotal
}

var CurBlockNumberResponse struct {
	RangeTotal int64 `json:"rangeTotal"`
}

func main() {
	flag.Parse()
	var customAPIKey string
	flag.StringVar(&customAPIKey, "apikey", "", "the tron api key ")
	if apiKey != "" {
		apiKey = customAPIKey
		singleHeader.Add("TRON-PRO-API-KEY", apiKey)
	}
	run()
}

var CurState struct {
	BlockNumber      int64 `json:"blockNumber"`
	ShiftBlockNumber int64 `json:"-"`
	TotalBlockNumber int64 `json:"-"`
}
var rw sync.RWMutex

func run() {
	initCurState()
	syncBlockNumber()
	syncTransaction()
}

func syncTransaction() {
	for {
		rw.RLock()
		var tmpBlockNumber = CurState.ShiftBlockNumber
		var targetBlockNumber = CurState.TotalBlockNumber
		rw.RUnlock()
		start := time.Now()
		workChan := make(chan struct{}, 10)
		for tmpBlockNumber < targetBlockNumber {
			workChan <- struct{}{}
			go func(curTransaction int64) {
				work(curTransaction)
				<-workChan
			}(tmpBlockNumber)
			tmpBlockNumber++
			var curBlockNumber = atomic.LoadInt64(&CurState.BlockNumber)
			if curBlockNumber+int64(cap(workChan)) < tmpBlockNumber {
				log.Printf("Push finish the blocknumber to %d", curBlockNumber+int64(cap(workChan)))
				atomic.StoreInt64(&CurState.BlockNumber, curBlockNumber+int64(cap(workChan)))
			}
		}
		end := time.Now()
		if end.Sub(start) <= 3*time.Minute {
			time.Sleep(10 * time.Second)
		}
	}
}

func work(blockNumber int64) {
	transactions, err := GetTransactionByBlock(blockNumber)
	if err != nil {
		panic(fmt.Sprintf("run GetTransactionByBlock failed:blockNumber:%d,failedinfo:%s", blockNumber, err.Error()))
	}
	log.Printf("blocknumber %d一共有%d笔交易", blockNumber, len(transactions.List))
	if len(transactions.List) <= 0 {
		return
	}
	var messages = make([]kafka.Message, 0, len(transactions.List))
	for i := 0; i < len(transactions.List); i++ {
		transactionByt, err := json.Marshal(transactions.List[i])
		if err != nil {
			panic(fmt.Sprintf("serialize the transaction %d:%d failed:%s", blockNumber, i, err.Error()))
		}
		message := kafka.Message{
			Key:   []byte(fmt.Sprintf("%d%d", blockNumber, i)),
			Value: transactionByt,
			Time:  time.Now(),
		}
		messages = append(messages, message)
	}
	err = getKafkaWriter().WriteMessages(context.TODO(), messages...)
	if err != nil {
		panic(fmt.Sprintf("push transaction %d for kafka failed:%s", blockNumber, err.Error()))
	}
}

func syncBlockNumber() {
	go func() {
		for {
			newBlockNumebr := GetCurBlockNumber()
			log.Printf("cur blocknumber is %d", newBlockNumebr)
			rw.Lock()
			if CurState.TotalBlockNumber < newBlockNumebr {
				CurState.TotalBlockNumber = newBlockNumebr
			}
			rw.Unlock()
			time.Sleep(1 * time.Minute)
		}
	}()
}
func initCurState() {
	file, err := os.Open(defualtStateFileLocation)
	defer file.Close()
	begin := make(chan struct{}, 0)
	go func() {
		<-begin
		close(begin)
		log.Println("presistent thread is running")
		timer := time.NewTicker(30 * time.Second)
		sysSignal := make(chan os.Signal, 0)
		signal.Notify(sysSignal, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGINT, syscall.SIGABRT, syscall.SIGHUP)
		flushDisk := time.NewTicker(1 * time.Minute)
		tmpFile, err := os.OpenFile(defualtStateFileLocation, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(fmt.Sprintf("open the state file %s error:%s", defualtStateFileLocation, err.Error()))
		}
		var tmpState []byte
		for {
			select {
			case <-sysSignal:
				tmpFile.Sync()
				log.Println("system exit the program,keeping the state for next execution")
				err = tmpFile.Truncate(0)
				if err != nil {
					panic(fmt.Sprintf("reset the state file %s content error:%s", defualtStateFileLocation, err.Error()))
				}
				tmpFile.Write(tmpState)
				err = file.Sync()
				if err != nil {
					log.Printf("syscall the file sync method error:%s", err.Error())
				}
				os.Exit(1)
			case <-flushDisk.C:
				log.Printf("cur state file %s keep to disk", defualtStateFileLocation)
				// tmpFile, err := os.OpenFile(defualtStateFileLocation, os.O_CREATE|os.O_TRUNC, 0644)
				// if err != nil {
				// 	panic(fmt.Sprintf("open the state file %s error:%s", defualtStateFileLocation, err.Error()))
				// }
				tmpFile.Truncate(0)
				tmpFile.Write(tmpState)
			case <-timer.C:
				// log.Println("starting  presist the state for disk")
				rw.Lock()
				tmpState, err = json.Marshal(CurState)
				rw.Unlock()
				if err != nil {
					panic(fmt.Sprintf("serilize the state file %s error:%s,maybe the file has destroyed", defualtStateFileLocation, err.Error()))
				}
			default:
				time.Sleep(1 * time.Second)
			}
		}
	}()
	if os.IsNotExist(err) {
		CurState.BlockNumber = 0
		CurState.ShiftBlockNumber = 0
		return
	}
	fileByt, err := io.ReadAll(file)
	if err != nil {
		os.Remove(defualtStateFileLocation)
		panic(fmt.Sprintf("open the state file %s error:%s", defualtStateFileLocation, err.Error()))
	}
	err = json.Unmarshal(fileByt, &CurState)
	if err != nil {
		panic(fmt.Sprintf("serilize the state file %s error:%s,maybe the file has destroyed", defualtStateFileLocation, err.Error()))
	}
	CurState.ShiftBlockNumber = CurState.BlockNumber
	log.Printf("Starting  work  from blockNumber %d", CurState.BlockNumber)
	begin <- struct{}{}
}

var tronClient = http.DefaultClient
var tronOnce sync.Once

var searchTransactionsByBlock = "https://apilist.tronscanapi.com/api/transaction?sort=-timestamp&block=%d&limit=200"

type TransactionInfoList struct {
	List []*TransactionInfo `json:"data"`
}
type TransactionInfo struct {
	Block         int64    `json:"block"`
	Hash          string   `json:"hash"`
	TimeStamp     int64    `json:"timestamp"`
	OwnerAddress  string   `json:"ownerAddress"`
	ToAddressList []string `json:"toAddressList"`
	ToAddress     string   `json:"toAddress"`
	ContractType  int64    `json:"contractType"`
	Confirmed     bool     `json:"confirmed"`
	Revert        bool     `json:"revert"`
	ContractData  struct {
		Amount        int64  `json:"amount"`
		Owner_address string `json:"owner_address"`
		To_address    string `json:"to_address"`
	} `json:"contractData"`
	SmartCalls  string `json:"SmartCalls"`
	Events      string `json:"Events"`
	Id          string `json:"id"`
	Data        string `json:"data"`
	Fee         string `json:"fee"`
	ContractRet string `json:"contractRet"`
	Result      string `json:"result"`
	Amount      string `json:"amount"`
	Cost        struct {
		Net_fee              int64 `json:"net_fee"`
		Energy_penalty_total int64 `json:"energy_penalty_total"`
		Energy_usage         int64 `json:"energy_usage"`
		Fee                  int64 `json:"fee"`
		Energy_fee           int64 `json:"energy_fee"`
		Energy_usage_total   int64 `json:"energy_usage_total"`
		Origin_energy_usage  int64 `json:"origin_energy_usage"`
		Net_usage            int64 `json:"net_usage"`
	} `json:"cost"`
	TokenInfo struct {
		TokenID      string `json:"tokenId"`
		TokenAbbr    string `json:"tokenAbbr"`
		TokenName    string `json:"tokenName"`
		TokenDecimal int64  `json:"tokenDecimal"`
		TokenCanShow int64  `json:"tokenCanShow"`
		TokenType    string `json:"tokenType"`
		TokenLogo    string `json:"tokenLogo"`
		TokenLevel   string `json:"tokenLevel"`
		Vip          bool   `json:"vip"`
	} `json:"tokenInfo"`
	TokenType string `json:"tokenType"`
}

// type TransactionInfoList struct {
// 	List []*TransactionInfo `json:"data"`
// }
// type TransactionInfo struct {
// 	Id                               string                 `json:"id"`
// 	Fee                              string                 `json:"fee"`
// 	BlockNumber                      int64                  `json:"blockNumber"`
// 	BlockTimeStamp                   int64                  `json:"blockTimeStamp"`
// 	ContractResult                   []string               `json:"contractResult"`
// 	Contract_address                 string                 `json:"contract_address"`
// 	Receipt                          ResourceReceipt        `json:"receipt"`
// 	Log                              []*Log                 `json:"log"`
// 	Result                           string                 `json:"result"`
// 	ResMessage                       string                 `json:"resMessage"`
// 	AssetIssueID                     string                 `json:"assetIssueID"`
// 	Withdraw_amount                  int64                  `json:"withdraw_amount"`
// 	Unfreeze_amount                  int64                  `json:"unfreeze_amount"`
// 	Internal_transactions            []*InternalTransaction `json:"internal_transactions"`
// 	Exchange_received_amount         int64                  `json:"exchange_received_amount"`
// 	Exchange_inject_another_amount   int64                  `json:"exchange_inject_another_amount"`
// 	Exchange_withdraw_another_amount int64                  `json:"exchange_withdraw_another_amount"`
// 	Exchange_id                      int64                  `json:"exchange_id"`
// 	Shielded_transaction_fee         int64                  `json:"shielded_transaction_fee"`
// 	OrderID                          string                 `json:"orderId"`
// 	OrderDetails                     []*MarketOrderDetail   `json:"orderDetails"`
// 	PackingFee                       int64                  `json:"packingFee"`
// 	Withdraw_expire_amount           int64                  `json:"withdraw_expire_amount"`
// 	Cancel_unfreezeV2_amount         map[string]int64       `json:"cancel_unfreezeV2_amount"`
// }
// type InternalTransaction struct {
// 	Hash               string `json:"hash"`
// 	Caller_address     string `json:"caller_address"`
// 	TransferTo_address string `json:"transferTo_address"`
// 	CallValueInfo      []struct {
// 		CallValue int64  `json:"callValue"`
// 		TokenId   string `json:"tokenId"`
// 	} `json:"callValueInfo"`
// 	Note     string `json:"note"`
// 	Rejected bool   `json:"rejected"`
// 	Extra    string `json:"extra"`
// }

// type MarketOrderDetail struct {
// 	MarkerOrderId    string `json:"markerOrderId"`
// 	TakerOrderId     string `json:"takerOrderId"`
// 	FillSellQuantity int64  `json:"fillSellQuantity"`
// 	FillBuyQuantity  int64  `json:"fillBuyQuantity"`
// }
// type Log struct {
// 	Address string   `json:"address"`
// 	Topics  []string `json:"topics"`
// 	Data    string   `json:"data"`
// }
// type ResourceReceipt struct {
// 	Energy_usage         int64 `json:"energy_usage"`
// 	Energy_fee           int64 `json:"energy_fee"`
// 	Origin_energy_usage  int64 `json:"origin_energy_usage"`
// 	Energy_usage_total   int64 `json:"energy_usage_total"`
// 	Net_usage            int64 `json:"net_usage"`
// 	Net_fee              int64 `json:"net_fee"`
// 	Result               int64 `json:"result"`
// 	Energy_penalty_total int64 `json:"energy_penalty_total"`
// }

var singleHeader = http.Header{}

var client = http.Client{}

var apiKey = "5ed25c5c-d4d6-45c2-a573-1121340bce3a"

func init() {
	singleHeader.Add("TRON-PRO-API-KEY", apiKey)
}

func GetTransactionByBlock(number int64) (*TransactionInfoList, error) {
reset:
	req, err := http.NewRequest("GET", fmt.Sprintf(searchTransactionsByBlock, number), nil)
	if err != nil {
		return nil, err
	}
	req.Header = singleHeader

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.Status != "200 OK" {
		resp.Body.Close()
		goto reset
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var transactionsList = new(TransactionInfoList)
	err = json.Unmarshal(body, transactionsList)
	if err != nil {
		return nil, err
	}
	// for i := 0; i < len(transactionsList.List); i++ {
	// 	fmt.Println(transactionsList.List[i])
	// }
	return transactionsList, nil
}

//	func GetTransactionByBlock(number int) error {
//		var serializeNum = &NumberJson{
//			Num: number,
//		}
//		valueByt, err := json.Marshal(serializeNum)
//		if err != nil {
//			log.Printf("serialize the blocknumber %d error:%s", number, err.Error())
//			return err
//		}
//		resp, err := tronClient.Post(searchBlockByNumberURL, "application/json;charset=utf-8", bytes.NewReader(valueByt))
//		if err != nil {
//			log.Printf("get transactions from %d block failed:%s", number, err.Error())
//			return err
//		}
//		respBytes, err := io.ReadAll(resp.Body)
//		if err != nil {
//			log.Printf("read the response from search trasactions from %d block method failed:%s", number, err.Error())
//			return err
//		}
//		return nil
//	}
var producer *kafka.Writer
var kafkaOnce sync.Once

func getKafkaWriter() *kafka.Writer {
	kafkaOnce.Do(func() {
		producer = kafka.NewWriter(kafka.WriterConfig{
			Brokers:      []string{"bigdata004:9092", "bigdata003:9092", "bigdata002:9092"},
			Topic:        "tron-transaction",
			Balancer:     &kafka.LeastBytes{},
			RequiredAcks: 1,
		})
	})
	return producer
}
