package kagi

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	uuid := "123e4567-e89b-12d3-a456-426655440000"

	k := New("12345678901234567890123456789012")
	cipherText := k.Encrypt(uuid)

	if cipherText == uuid {
		t.Errorf("Expected cipherText to be different from uuid")
	}
}

func TestDecrypt(t *testing.T) {
	uuid := "123e4567-e89b-12d3-a456-426655440000"

	k := New("12345678901234567890123456789012")
	cipherText := k.Encrypt(uuid)

	if cipherText == uuid {
		t.Errorf("Expected cipherText to be different from uuid")
	}

	plainText := k.Decrypt(cipherText)

	if plainText != uuid {
		t.Errorf("Expected plainText to be equal to uuid")
	}
}
