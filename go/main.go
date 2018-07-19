package main

import (
  "fmt"
  "log"

  "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
  "github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

type Config struct {
  RpcServer string `json:"rpc_server"`
  TokenAddress string `json:"token_address"`
  TestAccount string `json:"test_account"`
}

func main() {
  config.Load(file.NewSource(file.WithPath("config.json")))
  var conf Config
  config.Scan(&conf)

  conn, err := ethclient.Dial(conf.RpcServer)
  if err != nil {
    log.Fatalf("Failed to connect to Ethereum node: %v", err)
  }

  contract, err := NewTronToken(common.HexToAddress(conf.TokenAddress), conn)
  if err != nil {
    log.Fatalf("Failed to create contract instance: %v", err)
  }

  bal, _ := contract.BalanceOf(&bind.CallOpts{}, common.HexToAddress(conf.TestAccount))
  fmt.Println(bal)
}
