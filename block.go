package main

import (
	"time"
)

type Block struct {
	// block header:
	//
	Timestamp int64
	PrevHash  []byte
	Hash      []byte
	Nonce     int //counter for POW
	// head end

	// block body:
	Data []byte
}

func NewBlock(data string, prevHash []byte) *Block {

	block := &Block{
		Timestamp: time.Now().Unix(),
		PrevHash:  prevHash,
		Hash:      []byte{},
		Data:      []byte(data),
		Nonce:     0,
	}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Nonce = nonce
	block.Hash = hash

	return block
}

func NewGenesis() *Block {
	return NewBlock("genesis block", []byte{})
}
