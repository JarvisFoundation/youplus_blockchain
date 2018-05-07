package engine

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"

	"../utils"
)

const targetBits = 24

var (
	maxNonce = math.MaxInt64
)

//PowService ...
type PowService struct{}

//UtilService ...
type UtilService struct{}

// ProofOfWork struct to create pow
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

//NewProofOfWork Create Proof of Work
func (pw *PowService) NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	pow := &ProofOfWork{b, target}
	return pow
}

//PrepareData Prepare blockchain data
func (pow *ProofOfWork) PrepareData(nonce int) []byte {
	var utilService = new(utils.HelperService)
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			utilService.IntToHex(pow.block.Timestamp),
			utilService.IntToHex(int64(targetBits)),
			utilService.IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

//Run Prepare blockchain data
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	fmt.Println("Mining block:", pow.block.Data)
	for nonce < maxNonce {
		data := pow.PrepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	return nonce, hash[:]
}

//Validate Validate proof of work
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := pow.PrepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}
