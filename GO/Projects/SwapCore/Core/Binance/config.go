package Binance

import (
	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/ini.v1"
	"swapcore/Core/lib"
)

type Config struct {
	nodeAddress     string
	deployerAddress string
	apiKey          string
}

func (c *Config) GetDeployerAddress() common.Address {
	return common.HexToAddress(c.deployerAddress)
}

func (c *Config) GetNodeAddress() string {
	return c.nodeAddress
}

func (c *Config) GetAPIKEY() string {
	return c.apiKey
}

var BSCConfig Config

func init() {
	cfg, err := ini.Load("./config.ini")
	lib.CheckError(err)
	section := cfg.Section("BSC")
	BSCConfig.nodeAddress = section.Key("Node").String()
	BSCConfig.deployerAddress = section.Key("Deployer").String()
	BSCConfig.apiKey = section.Key("API_KEY").String()
}
