package main

import (
	"github.com/ethereum/go-ethereum/common"
	"swapcore/Core"
	"swapcore/lib"
)

//var BSCClient *ethclient.Client

//func init() {
//var err error
//BSCClient, err = ethclient.Dial(lib.BSCConfig.NodeAddress)
//lib.CheckError(err)
//}

func main() {
	handler := Core.PoolHandler{
		PoolAddress: common.HexToAddress(lib.BSCConfig.DeployerAddress),
		NodeAddress: lib.BSCConfig.NodeAddress,
	}
	handler.Subscribe()
	defer handler.Con.Close()
}
