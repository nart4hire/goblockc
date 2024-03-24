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
package utils

import (
	"encoding/binary"
	"encoding/hex"
	"strconv"
	"testing"
)

func TestGetSubkeys(t *testing.T) {
	key := []byte("abcdefghijklmnopqrstuvwxyz")
	subkeys, err := GetSubkeys(key[:16])

	if err != nil {
		t.Error("Unexpected error")
	}

	for i := range 16 {
		if len(subkeys[i]) != 8 {
			t.Error("Subkey length is not 8")
		}
		t.Log(subkeys[i])
	}
}

func TestLeftRotate(t *testing.T) {
	nums := []byte{0x00, 0x01, 0x02, 0x03, 0x04}
	for i := range 5 {
		rotated := Rotate(nums, i, true)
		t.Log(rotated)

		if len(rotated) != len(nums) {
			t.Error("Rotated length is not equal to original length")
		}

		for j := range len(rotated) {
			if rotated[j] != nums[(j+i)%len(nums)] {
				t.Error("Rotated improperly")
			}
		}
	}
}

func TestRightRotate(t *testing.T) {
	nums := []byte{0x00, 0x01, 0x02, 0x03, 0x04}
	for i := range 5 {
		rotated := Rotate(nums, i, false)
		t.Log(rotated)

		if len(rotated) != len(nums) {
			t.Error("Rotated length is not equal to original length")
		}

		for j := range len(rotated) {
			if nums[j] != rotated[(j+i)%len(nums)] {
				t.Error("Rotated improperly")
			}
		}
	}
}

func TestRotateInvertibility(t *testing.T) {
	nums := []byte{0x00, 0x01, 0x02, 0x03, 0x04}
	rotated := Rotate(nums, 2, true)
	rotated = Rotate(rotated, 2, false)

	if len(rotated) != len(nums) {
		t.Error("Rotated length is not equal to original length")
	}

	for i := range len(rotated) {
		if rotated[i] != nums[i] {
			t.Error("Rotated improperly")
		}
	}
}

func TestBytesToUInt64(t *testing.T) {
	b := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}
	num, err := BytesToUInt64(b)

	if err != nil {
		t.Error("Unexpected error")
	}

	expected, err := strconv.ParseUint(hex.EncodeToString(b), 16, 64)

	if err != nil {
		t.Error("Unexpected error")
		t.Log(err)
	}

	if num != expected {
		t.Error("Number is not correct")
	}
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, num)
	t.Log(num, expected)
	t.Log(hex.EncodeToString(bytes))
}
