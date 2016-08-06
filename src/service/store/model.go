package store

import "github.com/boltdb/bolt"

// Datastore is the interface for a store
type Datastore interface {
	All() ([]*Entry, error)
	Put(e *Entry) error
	Delete(e *Entry) error
}

// Store holds the Bolt database and our Crypter. It should satisfy the
// Datastore interface
type Store struct {
	DB      *bolt.DB
	Crypter *Crypter
}

// Entry is the format of a database entry
type Entry struct {
	ID    []byte `json:"key"`
	Key   []byte `json:"identifier"`
	Value []byte `json:"value"`
}

// LoginRequest is the request we receive when a user enters their master
// password into the front end
type LoginRequest struct {
	Master string `json:"master"`
}
