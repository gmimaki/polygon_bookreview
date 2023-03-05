package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	client, err := ethclient.Dial("https://polygon-mumbai.g.alchemy.com/v2/rJVPNM0wTvjXHH8R1-Y3cyAYvVo5Qi-C")
	if err != nil {
		panic(err)
	}

	fmt.Println(client)
}
