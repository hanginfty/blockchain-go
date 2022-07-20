package main

import (
	"fmt"
	"strconv"
)

func (cli *CLI) addBlock(data string) {
	cli.bc.AddBlock(data)
	fmt.Println("add block success")
}

func (cli *CLI) printChain() {
	iter := cli.bc.Iterator()

	for {
		b := iter.Next()

		fmt.Printf("prev hash: %x\n", b.PrevHash)
		fmt.Printf("Data: %s\n", b.Data)
		fmt.Printf("Hash: %x\n", b.Hash)
		pow := NewProofOfWork(b)
		fmt.Printf("proof of work: %s\n\n", strconv.FormatBool(pow.Validate()))

		if len(b.PrevHash) == 0 {
			break
		}
	}
}
