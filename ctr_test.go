package goblockc_test

import (
	"encoding/hex"
	"testing"

	. "github.com/nart4hire/goblockc"
)

func TestCTR(t *testing.T) {
	
	plaintext := []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	t.Log(hex.EncodeToString(plaintext))
	key := []byte("abcdefghijklmnop")
	t.Log(hex.EncodeToString(key))
	iv := []byte("abcdefghijklmnop")
	t.Log(hex.EncodeToString(iv))

	gbc, err := NewBlock(key)

	if err != nil {
		t.Error("Unexpected error")
	}

	stream, err := NewCTR(gbc, iv)

	if err != nil {
		t.Error("Unexpected error")
	}

	ciphertext := make([]byte, len(plaintext))
	copy(ciphertext, plaintext)
	stream.XORKeyStream(ciphertext, ciphertext)
	t.Log(hex.EncodeToString(ciphertext))

	stream, err = NewCTR(gbc, iv)

	if err != nil {
		t.Error("Unexpected error")
	}

	if hex.EncodeToString(plaintext) == hex.EncodeToString(ciphertext) {
		t.Error("Text was not encrypted properly")
		t.Log(hex.EncodeToString(plaintext), hex.EncodeToString(ciphertext))
	}

	decrypted := make([]byte, len(plaintext))
	copy(decrypted, ciphertext)
	stream.XORKeyStream(decrypted, decrypted)
	t.Log(hex.EncodeToString(decrypted))

	if hex.EncodeToString(plaintext) != hex.EncodeToString(decrypted) {
		t.Error("Decrypted text is not equal to original text")
		t.Log(hex.EncodeToString(plaintext), hex.EncodeToString(decrypted))
	}
}