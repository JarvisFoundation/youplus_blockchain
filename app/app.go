package app

import (
	"./db"
	"./engine"
)

//App has router and db instances
type App struct{}

//Init initializes the app with predefined configuration
func (app App) Init() {
	//var blockchainService = new(engine.Blockchain)
	//var powService = new(engine.PowService)

	var dbService = new(db.LDB)
	var blockchainService = new(engine.Blockchain)
	blockchainService.NewBlockChainLevel(dbService)
	/*data := dbService.Get([]byte("key"))
	if data != nil {
		fmt.Println("existing")
	} else {
		fmt.Println("new")
		dbService.Put([]byte("key"), []byte("test data2"))
	}

	//data := dbService.Get([]byte("key"))
	//fmt.Println(data)

	/*bc := blockchainService.NewBlockchain()
	bc.AddBlock("Send 1 JRVS to Ivan")
	bc.AddBlock("Send 2 more JRVS to Ivan")
	bc.AddBlock("Send 3 more JRVS to Ivory")
	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := powService.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}*/
}
