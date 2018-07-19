package main

import (
  "fmt"
  "log"

  "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient" // help us connect to node
)

func main() {
  conn, err := ethclient.Dial("http://127.0.0.1:7545")
  if err != nil {
    log.Fatalf("Failed to connect to Ethereum node: %v", err)
  }

  contract, err := NewTronToken(common.HexToAddress("0x2c2b9c9a4a25e24b174f26114e8926a9f2128fe4"), conn)
  if err != nil {
    log.Fatalf("Failed to create contract instance: %v", err)
  }

  bal, _ := contract.BalanceOf(&bind.CallOpts{}, common.HexToAddress("0x627306090abaB3A6e1400e9345bC60c78a8BEf57"))
  fmt.Println(bal)
}
