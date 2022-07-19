package main

import (
	"log"

	"github.com/boltdb/bolt"
)

type BlockChain struct {
	// hash of the latest block.
	tip []byte
	db  *bolt.DB
}

func NewBlockChain() *BlockChain {
	//return &BlockChain{[]*Block{NewGenesis()}}

	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	// TODO: db

	return &BlockChain{tip, db}
}

//func (chain *BlockChain) AddBlock(data string) {
//	newestBlock := chain.blocks[len(chain.blocks)-1]
//	newBlock := NewBlock(data, newestBlock.Hash)
//	chain.blocks = append(chain.blocks, newBlock)
//}
