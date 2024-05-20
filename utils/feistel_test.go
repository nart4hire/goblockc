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
package utils_test

import (
	"testing"

	. "github.com/nart4hire/goblockc/utils"
)

func TestFeistel(t *testing.T) {
	forward, err := Feistel(0x7d7d7d7d7d7d7d7d, 0xffffffffffffffff, true)
	if err != nil {
		t.Error("Unexpected error")
	}

	if forward != 0x0 {
		t.Error("Feistel is not correct")
	}
}

func TestFeistelInvertibility(t *testing.T) {
	forward, err := Feistel(0x0102030405060708, 0x090A0B0C0D0E0F10, true)
	if err != nil {
		t.Error("Unexpected error")
	}

	if backward, err := Feistel(forward, 0x090A0B0C0D0E0F10, false); err != nil || backward != 0x0102030405060708 {
		t.Error("Feistel is not invertible")
	}
}

func TestFeistelInvertibility2(t *testing.T) {
	forward, err := Feistel(0x0102030405060708, 0x090A0B0C0D0E0F10, true)
	if err != nil {
		t.Error("Unexpected error")
	}

	forward2, err := Feistel(forward, 0x090A0B0C0D0E0F10, true)
	if err != nil {
		t.Error("Unexpected error")
	}

	backward2, err := Feistel(forward2, 0x090A0B0C0D0E0F10, false)
	if err != nil {
		t.Error("Unexpected error")
	}

	if backward, err := Feistel(backward2, 0x090A0B0C0D0E0F10, false); err != nil || backward != 0x0102030405060708 {
		t.Error("Feistel is not invertible")
	}
}

func TestPermute(t *testing.T) {
	left := Permute(0x0102030405060708, true)

	if left != 0x0102030406070805 {
		t.Error("Left is not correct")
		t.Log(left)
	}
}

func TestPermuteInvertibility(t *testing.T) {
	right := Permute(0x0102030405060708, false)
	left := Permute(right, true)

	if left != 0x0102030405060708 {
		t.Error("Left is not correct")
	}
}
