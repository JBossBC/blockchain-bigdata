package Model


type Fee interface{
	GetBaseFee()*BaseFee
	GetNetFee()Fee
	GetContractFee()Fee
	GetExtFee()Fee
}

type BaseFee struct{
	Currency string
	Number float64
	Info   string
}

func(base *BaseFee)GetBaseFee()*BaseFee{
	return base
}

func(base *BaseFee)GetNetFee()Fee{
	return nil
}
func(base *BaseFee)GetContractFee()Fee{
	return nil
}
func(base *BaseFee)GetExtFee()Fee{
	return nil
}