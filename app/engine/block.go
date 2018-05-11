package engine

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"strconv"
	"time"
)

//BlockService ...
type BlockService struct{}

// Block keeps block headers
type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// NewBlock creates and returns Block
func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	var powService = new(PowService)
	block := &Block{time.Now().Unix(), transactions, prevBlockHash, []byte{}, 0}
	pow := powService.NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

// SetHash calculates and sets block hash
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

//HashTransactions Hash of all transactions on this block
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte
	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}

// SerializeBlock Serialization of block data to be stored in database
func (b *Block) SerializeBlock() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		fmt.Println("Serialization Error: " + err.Error())
	}
	return result.Bytes()
}

// DeserializeBlock Deserialize block data in byte array
func DeserializeBlock(d []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		fmt.Println("Deserialization Error: " + err.Error())
	}
	return &block
}
