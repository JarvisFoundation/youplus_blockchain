package engine

import (
	"encoding/hex"
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

//FindUnspentTransactions find all the unspent outputs which are not input to any transaction
func (bc *Blockchain) FindUnspentTransactions(address string) []Transaction {
	var unspentTXs []Transaction
	spentTXO := make(map[string][]int)
	bci := bc.Iterator()
	for {
		block := bci.Next()
		for _, tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)
		Outputs:
			for outIdx, out := range tx.Vout {
				if spentTXO[txID] != nil {
					for _, spentOut := range spentTXO[txID] {
						if spentOut == outIdx {
							continue Outputs
						}
					}
				}

				if out.UnlockInput(address) {
					unspentTXs = append(unspentTXs, *tx)
				}
			}

			if tx.IsCoinbase() == false {
				for _, in := range tx.Vin {
					if in.UnlockOutput(address) {
						inTxID := hex.EncodeToString(in.Txid)
						spentTXO[inTxID] = append(spentTXO[inTxID], in.Vout)
					}
				}
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
	return unspentTXs
}

//Next Get the block pointed by tip of the chain
func (i *BlockchainIterator) Next() *Block {
	var block *Block
	tipBlock := i.ldb.Get(i.currentHash)
	block = DeserializeBlock(tipBlock)
	i.currentHash = block.PrevBlockHash
	return block
}
