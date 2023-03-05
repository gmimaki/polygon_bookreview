package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	client, err := ethclient.Dial(os.Getenv("API_URL"))
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := client.ChainID(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ChainId:", chainId)

	// Get the balance of an account
	account := common.HexToAddress(os.Getenv("WALLET_ADDRESS"))
	balance, err := client.BalanceAt(ctx, account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Account balance:", balance)

	// Get the latest known block
	block, err := client.BlockByNumber(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Latest block:", block.Number().Uint64())
}
