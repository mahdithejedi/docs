package Binance

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"net/http"
	"strings"
	"swapcore/Core/lib"
)

type Event struct {
	token0      common.Address
	token1      common.Address
	fee         uint
	tickSpacing int
	address     common.Address
}

func (e *Event) GetAddress() common.Address {
	return e.address
}

var EventType string = "PoolCreated"
var bscScanUrl string = "https://api.bscscan.com/api?module=contract&action=getabi&address=%s&apikey=%s"

func getABI() string {
	url := fmt.Sprintf(bscScanUrl, BSCConfig.GetDeployerAddress(), BSCConfig.GetAPIKEY())
	resp, err := http.Get(url)
	lib.CheckError(err)
	defer func() {
		lib.CheckError(resp.Body.Close())
	}()

	var response struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Result  string `json:"result"`
	}
	lib.CheckError(json.NewDecoder(resp.Body).Decode(&response))

	if response.Status != "1" {
		lib.CheckError(errors.New(fmt.Sprintf("error getting contract ABI: %s", response.Message)))
	}

	return response.Result
}

func Decode(eventMsg types.Log) *Event {
	contractABI, err := abi.JSON(strings.NewReader(getABI()))
	if err != nil {
		lib.CheckError(err)
	}
	event := Event{}
	lib.CheckError(contractABI.UnpackIntoInterface(&event, EventType, eventMsg.Data))
	//fmt.Println("Decoded data is", event.From, "to", event.To, "With amount", event.Value)
	return &event
}
