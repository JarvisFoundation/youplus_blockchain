package engine

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

const subsidy = 10

//TransactionService ...
type TransactionService struct{}

//TXOutput transaction output
type TXOutput struct {
	Value        int
	ScriptPubKey string
}

//TXInput transaction input
type TXInput struct {
	Txid      []byte
	Vout      int
	ScriptSig string
}

// Transaction transaction base
type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

//NewCoinbaseTX first transaction(coinbase) of a block
func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to %s", to)
	}
	tx := Transaction{nil, []TXInput{TXInput{[]byte{}, -1, data}}, []TXOutput{TXOutput{subsidy, to}}}
	tx.SetTransactionID()
	return &tx
}

//SetTransactionID setting the transaction id
func (tx *Transaction) SetTransactionID() {
	var encoded bytes.Buffer
	var hash [32]byte
	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		fmt.Println(err.Error())
	}
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}
