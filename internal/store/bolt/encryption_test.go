package bolt

import "testing"

var (
	key        = "something that satisifies 32...."
	cipherText string
)

func TestEncrypt(t *testing.T) {
	crypter, err := InitaliaseCrypter(key)
	if err != nil {
		t.Error(err)
	}
	// Test for 32 byte strings
	c, err := crypter.Encrypt([]byte("something that satisifies 32...."))
	if err != nil {
		t.Error(err)
	}

	// Test for sub-32 bytes with padding
	_, err = crypter.Encrypt([]byte("this is odd"))
	if err != nil {
		t.Error(err)
	}

	/*
	 *	Test the key generation
	 */

	_, err = InitaliaseCrypter("")
	if err == nil {
		t.Errorf("expected %v got %v", err, nil)
	}

	// Push out cipher to a varible to decrypt
	cipherText = c
}

func TestDecrypt(t *testing.T) {
	crypter, err := InitaliaseCrypter(key)
	if err != nil {
		t.Error(err)
	}

	plaintext, err := crypter.Decrypt([]byte(cipherText))
	if err != nil {
		t.Error(err)
	}

	expected := "something that satisifies 32...."
	if string(plaintext) != expected {
		t.Errorf("error expected %s got %s", expected, plaintext)
	}

}

func TestPad(t *testing.T) {
	in := []byte("memes")
	resp := pad(in)
	if string(resp) != "memes000000000000000000000000000" {
		t.Errorf("got incorrect padding length of %v", len(resp))
	}
}
