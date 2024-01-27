package Model

type LayerLevel int

const (
	Layer_0 LayerLevel = 0
	Layer_1 LayerLevel = 1
	Layer_2 LayerLevel = 2
	Layer_3 LayerLevel = 3
)

type BlockChain interface {
	GetChainName() string
	GetVersion() string
	GetCurrency() string
	//获取最底层依赖链,或者跨链桥网络方案
	GetLow_Level() BlockChain
}

type BaseBlockchain struct {
	ChainName      string
	Version        string
	OriginCurrency string
	Layer          LayerLevel
}

func (base *BaseBlockchain) GetChainName() string {
	return base.ChainName
}
func (base *BaseBlockchain) GetVersion() string {
	return base.Version
}
func (base *BaseBlockchain) GetCurrency() string {
	return base.OriginCurrency
}
func (base *BaseBlockchain) GetLow_Level() BlockChain {
	return nil
}
