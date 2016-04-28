package store

import (
	"errors"

	"github.com/boltdb/bolt"
)

// Store is a wrapper around Bolt's database
type Store struct {
	db *bolt.DB
}

// Field contains data to interface with the database
type Field struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Init will return an initialised database store
func Init() (*Store, error) {
	db, err := bolt.Open("./store.bolt", 0666, nil)
	if err != nil {
		return nil, err
	}

	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("store"))
		return nil
	})

	return &Store{
		db: db,
	}, nil
}

// Get will retrieve a value from the store
func (s *Store) Get(key string) (*Field, error) {
	var value string

	bKey := []byte(key)
	if bKey == nil {
		return nil, ErrEmptyKey
	}

	err := s.db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("store"))
		// HACK(mnzt): expensive on memory
		v := b.Get(bKey)
		if v == nil {
			return ErrNotFound
		}
		value = string(v)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &Field{
		Key:   key,
		Value: value,
	}, nil
}

// Put will place a field into the database
func (s *Store) Put(in *Field) error {
	if in == nil {
		return ErrEmptyKey
	}
	err := s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("store"))
		if err := b.Put([]byte(in.Key), []byte(in.Value)); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

var (
	ErrEmptyKey = errors.New("error empty key provided")
	ErrNotFound = errors.New("error key/value not found")
)
