package main

import (
	"fmt"
	"strconv"

	"github.com/just-benedict-it/blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First block after Genesis")
	chain.AddBlock("second block after Genesis")
	chain.AddBlock("third block after Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash : %x\n", block.PrevHash)
		fmt.Printf("Data : %s\n", block.Data)
		fmt.Printf("Hash : %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

}
