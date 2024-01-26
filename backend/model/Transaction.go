package Model

import(
	"time"
	"math/big"
)


type TxType int


func(txType TxType)String()string{
	 switch txType{
	 	case Normal:
			return "普通转账"
		case CallContract:
			return "合约调用"	
		case CreateContract:
			return "合约创建"	
		case DestructionContract:
			return 	"合约销毁"
		default:
			return "未知交易"	
	 }
}

const(
	Normal TxType = 1
	CallContract TxType =2
	CreateContract TxType = 3
	DestructionContract TxType =4
)

type Transactions []*Transaction
type Transaction struct{
	Inner TxData
	Time time.Time
	//交易的Hash值
	Hash string
	Size int64
	// 交易的发起者
	From string
	//交易最终的执行状态
	Status bool
	// 交易所在区块的位置
	TransactionIndex int
	// 交易所在的区块Hash
	BlockHash string
	// 交易所在的区块高度
	BlockHeight int64
	//交易相关的费用
	TransactionFee Fee
}


type TxData interface{
	TxType() TxType
	Nonce()uint64
	Value()*big.Int
	Data() []byte
	To()string
	Fee() *Fee
}

