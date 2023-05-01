package main

import (
	"swapcore/Core"
	"swapcore/Core/Binance"
)

//var BSCClient *ethclient.Client

//func init() {
//var err error
//BSCClient, err = ethclient.Dial(lib.BSCConfig.NodeAddress)
//lib.CheckError(err)
//}

//var handler Core.PoolHandler

func watcher() {
	handler := Core.PoolHandler{
		PoolAddress: Binance.BSCConfig.GetDeployerAddress(),
		NodeAddress: Binance.BSCConfig.GetNodeAddress(),
	}
	handler.Subscribe()
	defer handler.Con.Close()
}

func main() {
	watcher()
	//defer func() { handler.Con.Close() }()
}
