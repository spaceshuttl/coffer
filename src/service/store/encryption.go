package store

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	mrand "math/rand"
	"strings"
	"time"
)

var (
	padText = "0"
)

// Crypter the the top level crypter used for encrypting and decrypting store values
type Crypter struct {
	key   string
	block cipher.Block
}

// InitaliaseCrypter will return a new crypter with the key attached
func InitaliaseCrypter(key string) (*Crypter, error) {

	// HACK(replace this with perm usage of a master key)
	if key == "" {
		key = RandStringBytesMaskImprSrc(32)
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	return &Crypter{
		key:   key,
		block: block,
	}, nil
}

// Encrypt will use the crypter's keys and other values and encrypt p
func (c *Crypter) Encrypt(p string) (string, error) {

	p = pad(p)

	ciphertext := make([]byte, aes.BlockSize+len(p))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(c.block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], []byte(p))

	// blockmode := cipher.NewCBCEncrypter(c.block, iv)
	henc := hex.EncodeToString(ciphertext)

	return henc, nil
}

// Decrypt will use the crypter's keys and other values and descrypt p
func (c *Crypter) Decrypt(cipherhext string) (string, error) {

	ciphertext, err := hex.DecodeString(cipherhext)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	plaintext := ciphertext[aes.BlockSize:]

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
func pad(text string) string {
	length := aes.BlockSize - (len(text) % aes.BlockSize)
	for i := 0; i < length; i++ {
		text += padText
	}
	return text
}

// Unpad will remove the padding
func unpad(text string) string {
	return strings.Replace(text, padText, "", -1)
}

/*
 *	Credit: http://stackoverflow.com/a/31832326
 */

// RandStringBytesMaskImprSrc returns a random string of x length
func RandStringBytesMaskImprSrc(n int) string {
	var (
		src           = mrand.NewSource(time.Now().UnixNano())
		letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		letterIdxBits = uint(6)              // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)

	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & int64(letterIdxMask)); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
