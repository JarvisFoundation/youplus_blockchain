package engine

import (
	"fmt"

	"../db"
)

const genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

//BlockchainService ...
type BlockchainService struct{}

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	tip []byte
	ldb *db.LDB
}

// BlockchainIterator Iterates over blocks of blockchain
type BlockchainIterator struct {
	currentHash []byte
	ldb         *db.LDB
}

// NewBlockchain creates a new Blockchain with genesis Block
func (bc *Blockchain) NewBlockchain(address string, ldb *db.LDB) *Blockchain {
	var tip []byte
	ref := ldb.Get([]byte("1"))
	if ref == nil {
		fmt.Println("Creating Blockchain")
		cbtx := NewCoinbaseTX(address, genesisCoinbaseData)
		genesis := NewGenesisBlock(cbtx)
		ldb.Put(genesis.Hash, genesis.SerializeBlock())
		ldb.Put([]byte("1"), genesis.Hash)
		tip = genesis.Hash
	} else {
		fmt.Println("Blockchain present")
		tip = ldb.Get([]byte("1"))
	}
	return &Blockchain{tip, ldb}
}

//Iterator helper to iterate over blockchain
func (bc *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{bc.tip, bc.ldb}
}

//Next Get the block pointed by tip of the chain
func (i *BlockchainIterator) Next() *Block {
	var block *Block
	tipBlock := i.ldb.Get(i.currentHash)
	block = DeserializeBlock(tipBlock)
	i.currentHash = block.PrevBlockHash
	return block
}
