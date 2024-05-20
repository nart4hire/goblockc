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
	"reflect"
	"testing"
)

func TestKeyScheduleBuffer(t *testing.T) {
	ciphertext := []byte("abcdefghijklmnop")
	key := []byte("abcdefghijklmnop")
	gbc, err := NewBlock(key)

	GBC, ok := gbc.(*GoBlockC)
	if !ok {
		t.Error("Failed to cast to GoBlockC")
	}
	keySchedule1 := make([][]byte, 16)
	copy(keySchedule1, GBC.keySchedule)


	plaintext := make([]byte, 16)
	copy(plaintext, ciphertext)
	gbc.Decrypt(plaintext, plaintext)

	if err != nil {
		t.Error("Unexpected error")
	}

	GBC, ok = gbc.(*GoBlockC)
	if !ok {
		t.Error("Failed to cast to GoBlockC")
	}
	keySchedule2 := make([][]byte, 16)
	copy(keySchedule2, GBC.keySchedule)

	t.Log(keySchedule1)
	t.Log(keySchedule2)

	if !reflect.DeepEqual(keySchedule1, keySchedule2) {
		t.Error("Key schedule buffer is not equal")
	}
}