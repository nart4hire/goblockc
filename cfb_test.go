/*
Copyright © 2024 Nathanael Santoso, Gede Prasidha Bhawarnawa, Felicia Sutandijo <business@nathancs.dev, 13520004@std.stei.itb.ac.id, feliciasutandijo@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package goblockc_test

import (
	"encoding/hex"
	"testing"

	. "github.com/nart4hire/goblockc"
)

func TestCFB(t *testing.T) {
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

	stream, err := NewCFB(gbc, iv, Encrypt)

	if err != nil {
		t.Error("Unexpected error")
	}

	ciphertext := make([]byte, len(plaintext))
	copy(ciphertext, plaintext)
	stream.XORKeyStream(ciphertext, ciphertext)
	t.Log(hex.EncodeToString(ciphertext))

	stream, err = NewCFB(gbc, iv, Decrypt)

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