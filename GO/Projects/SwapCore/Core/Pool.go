package Core

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"net/rpc"
	"swapcore/lib"
)

type PoolHandler struct {
	PoolAddress       common.Address
	NodeAddress       string
	Con               *ethclient.Client ``
	pureRpcConnection *rpc.Client
}

func (p *PoolHandler) getConnection() *ethclient.Client {
	if p.Con != nil {
		return p.Con
	}
	client, err := ethclient.Dial(p.NodeAddress)
	lib.CheckError(err)
	p.Con = client
	return client
}

func (p *PoolHandler) getPureConnection() *rpc.Client {
	if p.pureRpcConnection != nil {
		return p.pureRpcConnection
	}
	var err error
	p.pureRpcConnection, err = rpc.DialHTTP("tpc", p.NodeAddress)
	lib.CheckError(err)
	return p.pureRpcConnection
}

func (p *PoolHandler) Subscribe() {
	sub, log := p.getSubscription()
	for {
		select {
		case subErr := <-sub.Err():
			lib.CheckError(subErr)
		case msg := <-log:
			fmt.Println(msg)

		}
	}

}

func (p *PoolHandler) getSubscription() (ethereum.Subscription, chan types.Log) {
	query := p.getQuery()
	logs := make(chan types.Log)
	sub, err := p.getConnection().SubscribeFilterLogs(context.Background(), query, logs)
	lib.CheckError(err)
	return sub, logs
}

func (p *PoolHandler) getQuery() ethereum.FilterQuery {
	return ethereum.FilterQuery{
		Addresses: []common.Address{p.PoolAddress},
	}

}
