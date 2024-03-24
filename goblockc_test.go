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
package goblockc

import (
	"encoding/hex"
	"testing"
)

func TestParse(t *testing.T) {
	plaintext := []byte("abcdefghijklmnopqrstuvwxyz")
	key := []byte("abcdefghijklmnopqrstuvwxyz")

	ciphertext, err := parse(plaintext[:16], key[:16], true)

	if err != nil {
		t.Error("Unexpected error")
	}

	decrypted, err := parse(ciphertext, key[:16], false)

	if err != nil {
		t.Error("Unexpected error")
	}

	if hex.EncodeToString(plaintext[:16]) != hex.EncodeToString(decrypted) {
		t.Error("Decrypted text is not equal to original text")
		t.Log(hex.EncodeToString(plaintext[:16]), hex.EncodeToString(decrypted))
	}
}
