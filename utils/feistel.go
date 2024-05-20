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
	"slices"
)

func Feistel(right uint64, key uint64, forward bool) (uint64, error) {
	if forward {
		sbox := GetSBox()
		bytes := make([]byte, 8)
		binary.BigEndian.PutUint64(bytes, right)

		for i := range 8 {
			bytes[i] = sbox[bytes[i]]
		}

		substitutedRight, err := BytesToUInt64(bytes)
		if err != nil {
			return 0, err
		}

		permutedRight := Permute(substitutedRight, forward)

		return permutedRight ^ key, nil
	} else {

		unXORedRight := right ^ key

		unpermutedRight := Permute(unXORedRight, forward)

		invsbox := GetInvSBox()
		bytes := make([]byte, 8)
		binary.BigEndian.PutUint64(bytes, unpermutedRight)

		for i := range 8 {
			bytes[i] = invsbox[bytes[i]]
		}

		unsubstitutedRight, err := BytesToUInt64(bytes)
		if err != nil {
			return 0, err
		}


		return unsubstitutedRight, nil
	}
}

func Permute(right uint64, forward bool) uint64 {
	rightBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(rightBytes, right)

	rightBytesOne := make([]byte, 4)
	copy(rightBytesOne, rightBytes[:4])
	rightBytesTwo := make([]byte, 4)
	copy(rightBytesTwo, rightBytes[4:])

	result := slices.Concat(rightBytesOne, Rotate(rightBytesTwo, 1, forward))
	return binary.BigEndian.Uint64(result)
}
