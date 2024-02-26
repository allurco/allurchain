package entity

import (
	"github.com/allurco/allurchain/internal/helpers"
	"github.com/dgraph-io/badger"
)

const (
	dbPath = "../../tmp/blocks"
)

type Blockchain struct {
	LastHash []byte
	Database *badger.DB
}

func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte
	err := bc.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		helpers.Handle(err)
		err1 := item.Value(func(val []byte) error {
			lastHash = append([]byte{}, val...)
			return nil
		})
		helpers.Handle(err1)
		return err
	})

	helpers.Handle(err)
	newBlock := CreateBlock(data, lastHash)
	err = bc.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		helpers.Handle(err)
		err1 := txn.Set([]byte("lh"), newBlock.Hash)
		helpers.Handle(err1)
		bc.LastHash = newBlock.Hash
		return err
	})

	helpers.Handle(err)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockchain() *Blockchain {
	var lastHash []byte
	opts := badger.DefaultOptions(dbPath)
	opts.Dir = dbPath
	opts.ValueDir = dbPath

	db, err := badger.Open(opts)
	helpers.Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			genesis := Genesis()
			println("Genesis Created")
			err := txn.Set(genesis.Hash, genesis.Serialize())
			helpers.Handle(err)
			err = txn.Set([]byte("lh"), genesis.Hash)
			lastHash = genesis.Hash
			return err
		} else {
			item, err := txn.Get([]byte("lh"))
			helpers.Handle(err)
			err1 := item.Value(func(value []byte) error {
				lastHash = append([]byte{}, value...)
				return nil
			})
			helpers.Handle(err1)
			return err
		}
	})

	helpers.Handle(err)
	blockchain := Blockchain{lastHash, db}
	return &blockchain
}
