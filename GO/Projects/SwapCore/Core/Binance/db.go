package Binance

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"swapcore/Core/DB"
	"swapcore/Core/lib"
)

func SavePoolAddress(address common.Address) {
	clinet := DB.Client
	c, err := clinet.Network.Create().SetNetworkID(56).SetName("BinanceSmartChain").SetSymbol("BSC").Save(context.Background())
	lib.CheckError(err)
	print("Network created successfully", c)
}
