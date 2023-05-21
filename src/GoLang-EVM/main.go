package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve the Infura ID from the environment variables
	infuraID := os.Getenv("INFURA_ID")
	if infuraID == "" {
		log.Fatal("INFURA_ID not found in .env file")
	}

	// Retrieve the QuickNode URL from the environment variables
	quickNodeURL := os.Getenv("QUICKNODE_URL")
	if quickNodeURL == "" {
		log.Fatal("QUICKNODE_URL not found in .env file")
	}

	// Connect to the Ethereum mainnet
	// client, err := ethclient.Dial("https://mainnet.infura.io/v3/" + infuraID)
	client, err := ethclient.Dial(quickNodeURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("We have a connection")

	// Fetch the latest block number
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Latest block number: %d\n", blockNumber)

	// Fetch and display information about the latest block
	block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Latest block details:\n")
	fmt.Printf("Block number: %d\n", block.Number().Uint64())
	fmt.Printf("Timestamp: %s\n", time.Unix(int64(block.Time()), 0).String())
	fmt.Printf("Gas limit: %d\n", block.GasLimit())
	fmt.Printf("Miner address: %s\n", block.Coinbase().Hex())

	// Fetch and display details of the latest transactions in the block
	transactions := block.Transactions()
	fmt.Printf("Latest transactions in the block:\n")
	for _, tx := range transactions {
		txHash := tx.Hash().Hex()
		from, err := client.TransactionSender(context.Background(), tx, block.Hash(), uint(block.Number().Uint64()))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Transaction Hash: %s\n", txHash)
		fmt.Printf("From: %s\n", from.Hex())
		if tx.To() != nil {
			fmt.Printf("To: %s\n", tx.To().Hex())
		} else {
			fmt.Println("To: <nil>")
		}
		fmt.Printf("Value: %s\n", tx.Value().String())
		fmt.Printf("Gas Limit: %d\n", tx.Gas())

		time.Sleep(1 * time.Second) // Add a delay of 1 second
	}

	// Fetch the suggested gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Suggested gas price: %s\n", gasPrice.String())

	// Fetch the network ID
	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Network ID: %d\n", networkID.Uint64())

	// Fetch the client version
	clientVersion := "go-ethereum"
	fmt.Printf("Client Version: %s\n", clientVersion)

	// Fetch the balance of a specific Ethereum address
	address := common.HexToAddress("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Balance: %s\n", balance.String())

	// Subscribe to new block events
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Subscribed to new block events")

	// Handle new block events
	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case header := <-headers:
				fmt.Printf("New block detected! Block number: %d\n", header.Number.Uint64())
				// Additional logic to handle new block events
			}
		}
	}()

	// Keep the program running indefinitely
	select {}
}
