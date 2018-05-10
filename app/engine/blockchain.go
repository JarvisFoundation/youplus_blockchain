package engine

import (
	"fmt"

	"../db"
)

//BlockchainService ...
type BlockchainService struct{}

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	tip    []byte
	Blocks []*Block
}

const dbFile = "jarvis.db"
const blocksBucket = "blocks"

// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// NewBlockchain creates a new Blockchain with genesis Block
func (bc *Blockchain) NewBlockchain() {
	//return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// NewBlockChainLevel creates a new Blockchain with genesis Block
func (bc *Blockchain) NewBlockChainLevel() {
	var tip []byte
	var dbService = new(db.LDB)
	b := dbService.Get([]byte("1"))
	if b == nil {
		fmt.Println("Create Blockchain: Genesis Block")
		genesis := NewGenesisBlock()
		dbService.Put(genesis.Hash, genesis.SerializeBlock())
		dbService.Put([]byte("l"), genesis.Hash)
		tip = genesis.Hash
	} else {
		fmt.Println("blockchain present:")
		tip = dbService.Get([]byte("1"))
		fmt.Println(tip)
	}
	//bc := Blockchain{tip, db}
	//return &bc
}
