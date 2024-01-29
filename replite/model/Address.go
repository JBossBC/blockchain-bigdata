package Model

import (
	"math/big"
)

type AddressType int

const (
	Contract AddressType = 1
	EOA      AddressType = 2
)

type AddressStatus int

const (
	//30天内转过账
	Active AddressStatus = 1
	// 不存在(从未执行过交易)
	NotExist AddressStatus = 2
	// 不活跃(上一次交易在30天之外)
	NotActive AddressStatus = 3
	// 被销毁(适用于合约)
	Destruction AddressStatus = 4
)

type Address interface {
	GetAddressType() AddressType
	GetBalance() *big.Int
	GetAddress() string
	GetCurNonce() uint64
	GetOwnerContract() []string
	GetStatus() AddressStatus
	GetLabels() []string
}

type BaseEOA struct {
	Address   string
	Balance   *big.Int
	Nonce     uint64
	Contracts []string
	Labels    []string
}

func (eoa *BaseEOA) GetAddressType() AddressType {
	return EOA
}
func (eoa *BaseEOA) GetStatus() AddressStatus {
	return NotExist
}
func (eoa *BaseEOA) GetBalance() *big.Int {
	return eoa.Balance
}
func (eoa *BaseEOA) GetCurNonce() uint64 {
	return eoa.Nonce
}
func (eoa *BaseEOA) GetOwnerContract() []string {
	return eoa.Contracts
}
func (eoa *BaseEOA) GetAddress() string {
	return eoa.Address
}
func (eoa *BaseEOA) GetLabels() []string {
	return eoa.Labels
}

type BaseContract struct {
	Address   string
	Balance   *big.Int
	Nonce     uint64
	Contracts []string
	Labels    []string
}

func (contract *BaseContract) GetAddressType() AddressType {
	return Contract
}
func (contract *BaseContract) GetStatus() AddressStatus {
	return NotExist
}
func (contract *BaseContract) GetBalance() *big.Int {
	return contract.Balance
}
func (contract *BaseContract) GetCurNonce() uint64 {
	return contract.Nonce
}
func (contract *BaseContract) GetOwnerContract() []string {
	return contract.Contracts
}
func (contract *BaseContract) GetAddress() string {
	return contract.Address
}
func (contract *BaseContract) GetLabels() []string {
	return contract.Labels
}
