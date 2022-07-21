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

		fmt.Println("****************************************************************************")
		fmt.Printf("prev hash: %x\n", b.PrevHash)
		fmt.Printf("Data: %s\n", b.Data)
		fmt.Printf("Hash: %x\n", b.Hash)
		pow := NewProofOfWork(b)
		fmt.Printf("Proof of work: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Printf("****************************************************************************\n\n")

		if len(b.PrevHash) == 0 {
			break
		}
	}
}
