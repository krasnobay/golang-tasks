package main

import (
	"fmt"
)

func main() {
	bc := NewBlockchain()

	bc.AddBlock("transfer 1")
	bc.AddBlock("transfer 2")
	bc.AddBlock("transfer 3")

	for _, block := range bc.blocks {
		fmt.Printf(
			"Prev Hash: %x\nData: %s\nHash: %x\n\n",
			block.PrevBlockHash,
			block.Data,
			block.Hash,
		)
	}
}
