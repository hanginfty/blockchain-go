package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
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

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)
	if err != nil {
		fmt.Printf("block %v serialize failed\n", b.Hash)
	}

	return res.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)

	if err != nil {
		fmt.Printf("deserialization failed: %v\n", d)
	}

	return &block
}
