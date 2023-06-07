package Binance

import (
	"github.com/ethereum/go-ethereum/core/types"
)

func Run(eventMsg types.Log) {
	SavePoolAddress(Decode(eventMsg).GetAddress())
}
