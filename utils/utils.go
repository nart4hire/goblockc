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
	"bytes"
	"encoding/binary"
	"errors"
)

func GetSubkeys(key []byte) ([][]byte, error) {
	if len(key) != 16 {
		return nil, errors.New("key must be 16 bytes long")
	}

	subkeys := make([][]byte, 16)
	for i := range 16 {
		subkeys[i] = Rotate(key, i, true)[:8]
	}

	return subkeys, nil
}

func Rotate(nums []byte, n int, left bool) []byte {
	n = n % len(nums)
	if left {
		return append(nums[n:len(nums):len(nums)], nums[:n]...)
	} else {
		return append(nums[len(nums)-n:len(nums):len(nums)], nums[:len(nums)-n]...)
	}
}

func BytesToUInt64(b []byte) (uint64, error) {
	if len(b) != 8 {
		return 0, errors.New("byte must be 8 bytes long")
	}

	var res uint64
	err := binary.Read(bytes.NewReader(b[:8]), binary.BigEndian, &res)
	if err != nil {
		return 0, err
	}

	return res, nil
}
