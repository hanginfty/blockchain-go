package main

type BlockChain struct {
	blocks []*Block
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
