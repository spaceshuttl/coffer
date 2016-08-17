package store

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"math/rand"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	padText     byte = '0'
	masterKey   string
	letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!#$%&()*+,-./:;<=>?@[]^_`{|}~"
)

// Crypter the the top level crypter used for encrypting and decrypting store values
type Crypter struct {
	key   []byte
	block cipher.Block
}

// InitaliaseCrypter will return a new crypter with the key attached
func InitaliaseCrypter(keyString string) (*Crypter, error) {
	var key = []byte(keyString)

	if keyString == "" {
		return nil, ErrNoKey
	}

	if len(keyString) != 32 {
		logrus.WithFields(logrus.Fields{
			"original": keyString,
		}).Debug("padding crypter key")

		key = pad([]byte(keyString))
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return &Crypter{
		key:   key,
		block: block,
	}, nil
}

// Encrypt will use the crypter's keys and other values and encrypt p
func (c *Crypter) Encrypt(p []byte) (string, error) {
	var (
		dst []byte
	)

	mode, err := cipher.NewGCM(c.block)
	if err != nil {
		return "", err
	}

	dst = mode.Seal(
		dst,
		randStr(mode.NonceSize()),
		p,
		nil,
	)

	logrus.WithFields(logrus.Fields{
		"original": string(p),
	}).Debug("encrypting value")

	return string(dst), nil
}

// Decrypt will use the crypter's keys and other values and descrypt p
func (c *Crypter) Decrypt(ciphertext []byte) ([]byte, error) {
	var (
		dst           []byte
		ciphertextHex []byte
	)

	logrus.Debug("decrypting value")
	hex.Decode(ciphertextHex, ciphertext)

	mode, err := cipher.NewGCM(c.block)
	if err != nil {
		return nil, err
	}

	dst, err = mode.Open(
		dst,
		randStr(mode.NonceSize()),
		ciphertextHex,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return dst, nil
}

// Pad providers a left padding function.
func pad(text []byte) []byte {
	var (
		// Hard coded AES-256
		length  = 32
		newText = make([]byte, length)
	)

	for i := 0; i < length; i++ {
		if i < len(text) {
			newText[i] = text[i]
			continue
		}
		newText[i] = padText
	}
	return newText
}

// Unpad will remove the padding
func unpad(text string) []byte {
	return []byte(strings.Replace(text, string(padText), "", -1))
}

func randStr(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}

// Errors
var (
	ErrNoKey = errors.New("no key given to crypter")
)
