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
	ldb *db.LDB
}

// BlockchainIterator Iterates over blocks of blockchain
type BlockchainIterator struct {
	currentHash []byte
	ldb         *db.LDB
}

// AddBlock add block to blockchain
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte
	lastHash = bc.ldb.Get([]byte("1"))
	newBlock := NewBlock(data, lastHash)
	bc.ldb.Put(newBlock.Hash, newBlock.SerializeBlock())
	bc.ldb.Put([]byte("1"), newBlock.Hash)
	bc.tip = newBlock.Hash
}

// NewBlockchain creates a new Blockchain with genesis Block
func (bc *Blockchain) NewBlockchain(ldb *db.LDB) *Blockchain {
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
