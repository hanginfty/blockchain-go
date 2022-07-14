package main

import "math/big"

const targetBits = 16

type BlockChain struct {
	blocks []*Block
}

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	return &ProofOfWork{
		block:  b,
		target: target,
	}
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{newGenesisBlock()}}
}

func newGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func (chain *BlockChain) AddBlock(data string) {
	newestBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := NewBlock(data, newestBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}
