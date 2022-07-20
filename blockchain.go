package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

type BlockChain struct {
	// hash of the latest block.
	latestH []byte
	db      *bolt.DB
}

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

func NewBlockChain() *BlockChain {
	var latestH []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	// if there are no blocks in DB, create a new one.
	err = db.Update(func(tx *bolt.Tx) error {
		bBucket := tx.Bucket([]byte(blocksBucket))
		if bBucket == nil {
			fmt.Println("no existing blocks in db, creating a new one.")
			genesis := NewGenesis()

			// create new blocks bucket:
			bBucket, err = tx.CreateBucket([]byte(blocksBucket))

			if err != nil {
				log.Panic(err)
			}

			// add (hash, block.serialization) into db:
			err = bBucket.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}

			// put (latest, genesis.hash) into db:
			err = bBucket.Put([]byte("latest"), genesis.Hash)
			if err != nil {
				log.Panic(err)
			}
			latestH = genesis.Hash
		} else {
			latestH = bBucket.Get([]byte("latest"))
		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	return &BlockChain{latestH, db}
}

func (bc *BlockChain) AddBlock(data string) {
	var latestHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error {
		bBucket := tx.Bucket([]byte(blocksBucket))
		latestHash = bBucket.Get([]byte("latest"))
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(data, latestHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		bBucket := tx.Bucket([]byte(blocksBucket))
		err := bBucket.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = bBucket.Put([]byte("latest"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		bc.latestH = newBlock.Hash

		return nil
	})
}

type ChainIterator struct {
	curHash []byte
	db      *bolt.DB
}

func (bc *BlockChain) Iterator() *ChainIterator {
	return &ChainIterator{bc.latestH, bc.db}
}

func (iter *ChainIterator) Next() *Block {
	var b *Block

	err := iter.db.View(func(tx *bolt.Tx) error {
		bBucket := tx.Bucket([]byte(blocksBucket))
		serializedBlock := bBucket.Get(iter.curHash)
		b = DeserializeBlock(serializedBlock)
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	iter.curHash = b.PrevHash

	return b
}
