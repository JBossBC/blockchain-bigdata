package rpc

import Model "backend/model"

type CrossBridge interface {
	GetFromBlockchain() *Model.BlockChain
	GetToBlockChain() *Model.BlockChain
	GetContractAddress() *Model.Address
}

type BaseBridge struct {
	From    *Model.BlockChain
	Address *Model.Address
	To      *Model.BlockChain
}

func (cross *BaseBridge) GetFromBlockchain() *Model.BlockChain {
	return cross.From
}

func (cross *BaseBridge) GetToBlockChain() *Model.BlockChain {
	return cross.To
}
func (cross *BaseBridge) GetContractAddress() *Model.Address {
	return cross.Address
}

type CrossBridgeServer struct {
}

func (bridgeServer *CrossBridgeServer) GetTargetBlockChain(from *Model.BlockChain, address string) *Model.BlockChain {
	return nil
}

func (bridgeServer *CrossBridgeServer) IsConnective(from *Model.BlockChain, to *Model.BaseBlock) *Model.Address {
	return nil
}
