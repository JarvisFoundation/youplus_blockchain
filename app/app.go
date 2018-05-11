package app

import (
	"fmt"
	"strconv"

	"./db"
	"./engine"
)

//App has router and db instances
type App struct{}

//Init initializes the app with predefined configuration
func (app App) Init() {
	var dbService = new(db.LDB)
	var blockchainService = new(engine.Blockchain)
	var powService = new(engine.PowService)
	bc := blockchainService.NewBlockchain(dbService)
	//bc.AddBlock("Send 1 JRVS to Ivan")
	//bc.AddBlock("Send 2 more JRVS to Ivan")

	bci := bc.Iterator()
	for {
		block := bci.Next()
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := powService.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
