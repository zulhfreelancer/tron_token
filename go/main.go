package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "math/big"
  "os"
  "strings"
  "time"

  "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

  "github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

type Config struct {
  RpcServer string `json:"rpc_server"`
  TokenAddress string `json:"token_address"`
  SenderAccount string `json:"sender_account"`
  SenderAccountPassword string `json:"sender_account_password"`
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

  /**
  * Check balance
  */
  balBefore, _ := contract.BalanceOf(&bind.CallOpts{}, common.HexToAddress(conf.SenderAccount))
  fmt.Println("before", balBefore)

  /**
  * Transfer
  */
  file, err := os.Open("private_key")
    if err != nil {
      log.Fatal("Failed to read file: %v", err)
    }
    defer file.Close()
  _file, err := ioutil.ReadAll(file)
  privateKey := string(_file)
  auth, err := bind.NewTransactor(
    strings.NewReader(privateKey),
    conf.SenderAccountPassword,
  )
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	tx, err := contract.Transfer(
    auth,
    common.HexToAddress("0x0000000000000000000000000000000000000000"),
    big.NewInt(1),
  )
	if err != nil {
		log.Fatalf("Failed to request token transfer: %v", err)
	}
	fmt.Printf("Transfer pending: 0x%x\n", tx.Hash())

  // wait for tx to get mined (simulation only)
  time.Sleep(30 * time.Second)

  /**
  * Check balance
  */
  balAfter, _ := contract.BalanceOf(&bind.CallOpts{}, common.HexToAddress(conf.SenderAccount))
  fmt.Println("after", balAfter)
}
