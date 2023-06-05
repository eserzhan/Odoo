package boltDb

import (
	"errors"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/eserzhan/tgBott/pkg/repository"
)

type BoltRepositry struct {
	db *bolt.DB
}

func NewBolt(db *bolt.DB) *BoltRepositry {
	return &BoltRepositry{db: db}
}

func (br *BoltRepositry) Save(chatId int64, token string, bucket repository.Bucket) error {
	return br.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket))

		return bucket.Put(intToByte(chatId), []byte(token))
	})
}

func (br *BoltRepositry) Get(chatId int64, bucket repository.Bucket) (string, error) {
	var token string

	err := br.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket))

		token = string(bucket.Get(intToByte(chatId)))

		return nil 
	})

	if token == "" {
		return "", errors.New("not found")
	}

	return token, err
}

func intToByte(chatId int64) []byte {
	return []byte(strconv.FormatInt(chatId, 10))
}