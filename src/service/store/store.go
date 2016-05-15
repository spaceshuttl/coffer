package store

import (
	"encoding/json"
	"os"

	"github.com/boltdb/bolt"
	"github.com/sirupsen/logrus"
)

var (
	path, _ = os.Getwd()

	bucket = []byte("store")
)

// Datastore is the interface for a store
type Datastore interface {
	All() ([]*Entry, error)
	Put(e *Entry) error
	Delete(id *Entry) error
}

// Store holds the Bolt database
type Store struct {
	DB *bolt.DB
}

// Entry is the format of a database entry
type Entry struct {
	ID    string `json:"key"`
	Key   string `json:"identifier"`
	Value string `json:"value"`
}

/**
 *	Store functions:
 */

// Start will load the database file ready for transactions
func Start() (*Store, error) {
	db, err := bolt.Open("store.bolt", 0666, nil)
	if err != nil {
		return nil, err
	}

	// Create our master bucket
	err = db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &Store{
		DB: db,
	}, nil
}

// All will return all entries from the database
func (s *Store) All() ([]*Entry, error) {
	entries := []*Entry{}

	// Read the store
	err := s.DB.View(func(tx *bolt.Tx) error {

		// Open our master bucket
		bucket := tx.Bucket(bucket)
		// Range over all key/values in our master bucket
		err := bucket.ForEach(func(k, v []byte) error {
			var entry Entry

			// Open an account bucket
			logrus.Debugf("Opening bucket %s", k)
			entry.ID = string(k)

			bkt := bucket.Bucket(k)
			err := bkt.ForEach(func(k, v []byte) error {

				switch string(k) {
				case "key":
					entry.Key = string(v)
				case "value":
					entry.Value = string(v)
				}

				return nil
			})

			// Add our found entry to the slice
			entries = append(entries, &entry)

			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return err
		}

		return nil
	})

	// Check if our transaction errored
	if err != nil {
		return nil, err
	}

	return entries, nil
}

// Put will place an entry into the store
func (s *Store) Put(e *Entry) error {
	// Update the store
	err := s.DB.Update(func(tx *bolt.Tx) error {

		// Open our bucket
		masterBucket, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}

		// Create our key-specific bucket
		bucket, err := masterBucket.CreateBucketIfNotExists(e.bucketID())
		if err != nil {
			return err
		}

		err = bucket.Put([]byte("key"), toStore(e.Key))
		if err != nil {
			return err
		}

		err = bucket.Put([]byte("value"), toStore(e.Value))
		if err != nil {
			return err
		}

		return nil
	})

	// Check if the transaction errored
	if err != nil {
		return err
	}

	// Return A-OK on that transaction
	return nil
}

// Delete will remove an entry from the store
func (s *Store) Delete(e *Entry) error {
	err := s.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucket)

		// Delete the entry bucket
		logrus.Debugf("deleting bucket %s", e.ID)
		err := bucket.DeleteBucket(toStore(e.ID))
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

/**
 *  Helpers:
 */

// toStore converts a string to a byte slice
func toStore(e string) []byte {
	return []byte(e)
}

// bucketID returns the ID for an entries' bucket
func (e *Entry) bucketID() []byte {
	return []byte(e.ID)
}

// Marshal returns the string values of an entry
func (e *Entry) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

// Unmarshal will unmarshal a byte array onto an entry type
func Unmarshal(d []byte, v interface{}) error {
	return json.Unmarshal(d, v)
}
