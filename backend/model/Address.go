package Model
import(
	"math/big"
)

type Address interface{
	GetBalance()*big.Int
	GetAddress()string
	GetCurNonce()uint64
	GetOnwerContract()[]string
}