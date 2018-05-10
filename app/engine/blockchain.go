package engine

import (
	"fmt"

	"../db"
)

//BlockchainService ...
type BlockchainService struct{}

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	tip []byte
	//Blocks []*Block
}

// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data string) {
	//prevBlock := bc.Blocks[len(bc.Blocks)-1]
	//newBlock := NewBlock(data, prevBlock.Hash)
	//bc.Blocks = append(bc.Blocks, newBlock)
}

// NewBlockchain creates a new Blockchain with genesis Block
func (bc *Blockchain) NewBlockchain() *Blockchain {
	var tip []byte
	return &Blockchain{tip}
}

// NewBlockChainLevel creates a new Blockchain with genesis Block
func (bc *Blockchain) NewBlockChainLevel(ldb *db.LDB) *Blockchain {
	var tip []byte
	ref := ldb.Get([]byte("1"))
	if ref == nil {
		fmt.Println("Creating Blockchain")
		genesis := NewGenesisBlock()
		ldb.Put(genesis.Hash, genesis.SerializeBlock())
		ldb.Put([]byte("1"), genesis.Hash)
		tip = genesis.Hash
	} else {
		fmt.Println("Blockchain present")
		tip = ldb.Get([]byte("1"))
	}
	bc2 := Blockchain{tip}
	return &bc2
}
