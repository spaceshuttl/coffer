package store

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	padText   = []byte("0")
	masterKey string
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

		key = pad([]byte(keyString), 32)
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

	logrus.WithFields(logrus.Fields{
		"value": p,
	}).Debug("encrypting value")

	p = pad(p, -1)

	ciphertext := make([]byte, aes.BlockSize+len(p))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(c.block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], p)

	// blockmode := cipher.NewCBCEncrypter(c.block, iv)
	henc := hex.EncodeToString(ciphertext)

	// TODO: authenticated encryption
	// block, err := aes.NewCipher([]byte(c.key))
	// if err != nil {
	// 	panic(err.Error())
	// }
	//
	// nonce := make([]byte, 12)
	// if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
	// 	panic(err.Error())
	// }
	//
	// aesgcm, err := cipher.NewGCM(block)
	// if err != nil {
	// 	panic(err.Error())
	// }
	//
	// ciphertext := aesgcm.Seal(nil, nonce, []byte(p), nil)
	// fmt.Printf("%x\n", ciphertext)
	//
	// henc := hex.EncodeToString(ciphertext)

	return henc, nil
}

// Decrypt will use the crypter's keys and other values and descrypt p
func (c *Crypter) Decrypt(cipherhex []byte) ([]byte, error) {
	var (
		hexcipher = make([]byte, hex.DecodedLen(len(cipherhex)))
	)

	logrus.WithFields(logrus.Fields{
		"value": cipherhex,
	}).Debug("decrypting value")

	hex.Decode(hexcipher, cipherhex)

	_, err := hex.Decode(hexcipher, cipherhex)
	if err != nil {
		return nil, err
	}

	if len(hexcipher) < aes.BlockSize {
		panic("ciphertext too short")
	}

	iv := hexcipher[:aes.BlockSize]
	plaintext := hexcipher[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(plaintext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(c.block, iv)
	mode.CryptBlocks(plaintext, plaintext)

	s := string(plaintext)

	return unpad(s), nil
}

// Pad providers a left padding function.
func pad(text []byte, length int) []byte {
	if length <= 0 {
		length = aes.BlockSize - (len(text) % aes.BlockSize)
	}

	for i := 0; i < length; i++ {
		text = append(text, padText...)
	}
	return text
}

// Unpad will remove the padding
func unpad(text string) []byte {
	return []byte(strings.Replace(text, string(padText), "", -1))
}

// Errors
var (
	ErrNoKey = errors.New("no key given to crypter")
)
