package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Index        int    // position of the block within the chain
	Timestamp    string // time when the block was created
	HashValue    string // hash of the block, formed by combinging the data and the previous block
	Data         string
	PreviousHash string
}

// CalculateHash function will form the cache using data and the previous hash value. We will use sha256 hashing algorithim to form this hash.
func CalculateHash(b Block) string {
	record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PreviousHash
	hashingAlgo := sha256.New()
	hashingAlgo.Write([]byte(record))
	hashedValue := hashingAlgo.Sum(nil)
	return hex.EncodeToString(hashedValue)
}

// genrateBlock will generate a new block using data, calucate the hash and add it to the block
func generateBlock(oldBlock Block, data string) (Block, error) {
	newBlock := Block{
		Index:        oldBlock.Index + 1,
		Data:         data,
		PreviousHash: oldBlock.HashValue,
		Timestamp:    time.Now().String(),
	}
	newBlock.HashValue = CalculateHash(newBlock)
	return newBlock, nil
}

// isValidBlock will tell us if the blockchain is valid or has been tempered
func isValidBlock(previousBlock, currentBlock Block) bool {

	if previousBlock.Index+1 != currentBlock.Index {
		return false
	}
	if currentBlock.PreviousHash != previousBlock.HashValue {
		return false
	}
	if CalculateHash(currentBlock) != currentBlock.HashValue {
		return false
	}
	return true
}

func main() {
	genesisBlock := Block{
		Index:        0,
		Data:         "Genesis is a block",
		Timestamp:    time.Now().String(),
		PreviousHash: "",
		HashValue:    "",
	}
	genesisBlock.HashValue = CalculateHash(genesisBlock)

	secondBlock, _ := generateBlock(genesisBlock, "Second Block Data")
	fmt.Println(secondBlock)

	fmt.Println(isValidBlock(secondBlock, genesisBlock))
}
