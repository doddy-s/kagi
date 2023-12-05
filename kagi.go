package kagi

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

type Kagi struct {
	// secretKey is the key used to encrypt and decrypt the data
	// 32 bytes long (AES-256) key is a must.
	secretKey string
}

// New returns a new Kagi instance
// secretKey is the key used to encrypt and decrypt the data
// 32 bytes long (AES-256) key is a must.
func New(secretKey string) *Kagi {
	return &Kagi{secretKey: secretKey}
}

// Encrypt encrypts the plainText and returns the cipherText
func (k *Kagi) Encrypt(plainText string) string {
	aes, err := aes.NewCipher([]byte(k.secretKey))
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		panic(err)
	}

	// We need a 12-byte nonce for GCM (modifiable if you use cipher.NewGCMWithNonceSize())
	// A nonce should always be randomly generated for every encryption.
	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		panic(err)
	}

	// cipherText here is actually nonce+ciphertext
	// So that when we decrypt, just knowing the nonce size
	// is enough to separate it from the ciphertext.
	cipherText := gcm.Seal(nonce, nonce, []byte(plainText), nil)

	return string(cipherText)
}

// Decrypt decrypts the cipherText and returns the plainText
func (k *Kagi) Decrypt(cipherText string) string {
	aes, err := aes.NewCipher([]byte(k.secretKey))
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		panic(err)
	}

	// The nonce size is 12 bytes by default
	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		panic(err)
	}

	// Extract the nonce from the ciphertext
	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]

	// Decrypt the ciphertext using the AES
	plainText, err := gcm.Open(nil, []byte(nonce), []byte(cipherText), nil)
	if err != nil {
		panic(err)
	}

	return string(plainText)
}
