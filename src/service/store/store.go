package store

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"

	"github.com/boltdb/bolt"
	"github.com/sirupsen/logrus"
)

var (
	path, _ = os.Getwd()

	masterBucket = []byte("store")
	configDir    string
)

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
	ID    string `json:"key"`
	Key   string `json:"identifier"`
	Value string `json:"value"`
}

// LoginRequest is the request we receive when a user enters their master
// password into the front end
type LoginRequest struct {
	Master string `json:"master"`
}

/**
 *	Store functions:
 */

// Start will load the database file ready for transactions
func Start() (*Store, error) {
	// Initialise our crypter
	crypter, err := InitaliaseCrypter("one really really secure key... ")
	if err != nil {
		return nil, err
	}

	// Create our config directory
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	configDir = fmt.Sprintf("%s/.coffer/", usr.HomeDir)

	// Create our store

	logrus.Debugf("opening store at %s", configDir+"store.bolt")
	db, err := bolt.Open(configDir+"store.bolt", 0666, nil)
	if err != nil {
		return nil, err
	}

	// Create our master bucket
	if err = db.Update(func(tx *bolt.Tx) error {
		if _, err = tx.CreateBucketIfNotExists(masterBucket); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &Store{
		DB:      db,
		Crypter: crypter,
	}, nil

}

// IsNewInstall will return whether a
func (s *Store) IsNewInstall() bool {
	return true
}

// Initialise will create the config dir,
// func (s *Store) Initialise() error {
// 	stat, err := os.Stat(configDir)
// 	if err != nil {
// 		return err
// 	}
//
// 	err := os.Mkdir(configDir, 0744)
// 	if err != nil {
// 		return err
// 	}
//
// 	err = s.DB.Update(func(tx *bolt.Tx) error {
// 		bkt, err := tx.CreateBucketIfNotExists([]byte("meta"))
// 		if err != nil {
// 			return err
// 		}
//
// 		err = bkt.Put([]byte("initd"), []byte("true"))
// 		if err != nil {
// 			return err
// 		}
//
// 		return nil
// 	})
//
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }

// All will return all entries from the database
func (s *Store) All() ([]*Entry, error) {
	var entries []*Entry

	// Read the store
	logrus.Debugf("Opening database...")
	err := s.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(masterBucket)

		// Range over all key/values in our master bucket
		return bucket.ForEach(func(k, v []byte) error {
			var entry Entry

			// Open an account bucket
			logrus.Debugf("Opening bucket %s", k)
			entry.ID = string(k)

			bkt := bucket.Bucket(k)
			if err := bkt.ForEach(func(k, v []byte) error {

				switch string(k) {
				case "key":
					entry.Key = string(v)
				case "value":
					entry.Value = string(v)
				}

				return nil
			}); err != nil {
				return err
			}

			// Add our found entry to the slice
			entries = append(entries, &entry)
			return nil
		})
	})

	// Check if our transaction errored
	if err != nil {
		return nil, err
	}

	for i, entry := range entries {
		plaintext, err := s.Crypter.Decrypt(entry.Value)
		if err != nil {
			return nil, err
		}
		entries[i].Value = plaintext
	}

	return entries, nil
}

// Put will place an entry into the store
func (s *Store) Put(e *Entry) error {
	// Encrypt our values
	cipherValue, err := s.Crypter.Encrypt(e.Value)
	if err != nil {
		return err
	}

	// HACK(mnzt): we shouldn't really be changing values on the entry struct
	e.Value = cipherValue

	// Update the store
	return s.DB.Update(func(tx *bolt.Tx) error {

		// Open our bucket
		bucket := tx.Bucket(masterBucket)

		// Create our key-specific bucket
		bucket, err := bucket.CreateBucketIfNotExists(e.bucketID())
		if err != nil {
			return err
		}

		err = bucket.Put([]byte("key"), toStore(e.Key))
		if err != nil {
			return err
		}

		return bucket.Put([]byte("value"), toStore(e.Value))

	})
}

// Delete will remove an entry from the store
func (s *Store) Delete(e *Entry) error {
	return s.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(masterBucket)

		// Delete the entry bucket
		logrus.Debugf("deleting bucket %s", e.ID)
		return bucket.DeleteBucket(e.bucketID())
	})
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
func (e *Entry) Unmarshal(d []byte, v interface{}) error {
	return json.Unmarshal(d, v)
}
