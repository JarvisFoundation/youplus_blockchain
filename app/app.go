package app

import (
	"fmt"

	"./engine"
)

//App has router and db instances
type App struct{}

//Init initializes the app with predefined configuration
func (app App) Init() {
	var blockchainService = new(engine.Blockchain)
	bc := blockchainService.NewBlockchain()
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")
	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
