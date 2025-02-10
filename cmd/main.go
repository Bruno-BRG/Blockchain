package main

import (
	"Blockchain/internal/API"
	"Blockchain/internal/service"
)

func main() {
	// Populate the blockchain with sample blocks.
	bc := service.BlockchainInstance
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	// Start the server to listen for requests continuously.
	API.Run()
}
