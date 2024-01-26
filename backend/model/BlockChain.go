package Model



type BlockChain interface{
	GetChainName()string
	GetVersion() string
	GetCurrency()string
}


type BaseBlockchain struct{
	ChainName string
	Version string
	OriginCurrency string
}

func(base *BaseBlockchain)GetChainName()string{
	return base.ChainName
}
func(base *BaseBlockchain)GetVersion()string{
	return base.Version
}
func(base *BaseBlockchain)GetCurrency()string{
	return base.OriginCurrency
}