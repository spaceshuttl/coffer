package store

import "github.com/boltdb/bolt"

// Datastore is the interface for a store
type Datastore interface {
	All() ([]*Entry, error)
	Put(e *Entry) error
	Delete(id *Entry) error
}

// Store holds the Bolt database
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

// BucketID returns a byte slice of an entry's ID
// func (e *Entry) BucketID() []byte {
// 	return []byte(e.ID)
// }
//
// // EntryValue returns a byte slice of an entry's Value
// func (e *Entry) EntryValue() []byte {
// 	return []byte(e.Value)
// }
//
// // EntryKey returns a byte slice of an entry's Key
// func (e *Entry) EntryKey() []byte {
// 	return []byte(e.Key)
// }

// LoginRequest is the request we receive when a user enters their master
// password into the front end
type LoginRequest struct {
	Master string `json:"master"`
}
