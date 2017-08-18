package bolt

import (
	"fmt"
	"os"
	"os/user"

	"github.com/boltdb/bolt"
	"github.com/sirupsen/logrus"
)

var (
	masterBucket = []byte("store")
	dbFileName   = "store.bolt"

	currentUser, _ = user.Current()
	configDir      = fmt.Sprintf("%s/.coffer/", currentUser.HomeDir)
)

// Start will load the database file ready for transactions
func Start() (*Store, error) {

	if !DBExists() {
		err := InitialiseStore()
		if err != nil {
			return nil, err
		}
	}

	// Create our store
	logrus.Debugf("opening store at %s", configDir+"dbFileName")
	db, err := bolt.Open(configDir+"dbFileName", 0666, nil)
	if err != nil {
		return nil, err
	}

	crypter, err := InitaliaseCrypter("some password")
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

// AddCrypter will set the crypter to call methods on
func (s *Store) AddCrypter(c *Crypter) {
	s.Crypter = c
}

// All will return all entries from the database
func (s *Store) All() ([]*Entry, error) {
	var entries = make([]*Entry, 0)

	// Read the store
	logrus.Debugf("Opening database...")
	if err := s.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(masterBucket)

		// Range over all key/values in our master bucket
		return bucket.ForEach(func(k, v []byte) error {
			var entry Entry

			// Open an account bucket
			logrus.Debugf("Opening bucket %s", k)
			entry.ID = k

			bkt := bucket.Bucket(k)
			if err := bkt.ForEach(func(k, v []byte) error {

				switch string(k) {
				case "key":
					entry.Key = v
				case "value":
					entry.Value = v
				}

				return nil
			}); err != nil {
				return err
			}

			// Add our found entry to the slice
			entries = append(entries, &entry)
			return nil
		})
	}); err != nil {
		return nil, err
	}

	for i, entry := range entries {
		plaintext, err := s.Crypter.Decrypt(entry.Value)
		if err != nil {
			return nil, err
		}
		entries[i].Value = plaintext
	}

	// to prevent 'cannot read property map of null'
	if entries == nil {
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
	e.Value = []byte(cipherValue)

	// Update the store
	return s.DB.Update(func(tx *bolt.Tx) error {

		// Open our bucket
		bucket := tx.Bucket(masterBucket)

		// Create our key-specific bucket
		bucket, err := bucket.CreateBucketIfNotExists(e.ID)
		if err != nil {
			return err
		}

		if err := bucket.Put([]byte("key"), e.Key); err != nil {
			return err
		}

		return bucket.Put([]byte("value"), e.Value)
	})
}

// Delete will remove an entry from the store
func (s *Store) Delete(e *Entry) error {
	return s.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(masterBucket)

		// Delete the entry bucket
		logrus.Debugf("deleting bucket %s", e.ID)
		return bucket.DeleteBucket(e.ID)
	})
}

// DBExists checks if an existing db file exists
func DBExists() bool {
	_, err := os.Open(configDir + dbFileName)
	if err != nil {
		return false
	}
	return true
}

// InitialiseStore will create the config dir and db file
func InitialiseStore() error {

	_, err := os.Stat(configDir)
	if err == nil {
		os.Mkdir(configDir, 0744)
	}

	_, err = os.Create(configDir + dbFileName)
	return err
}
