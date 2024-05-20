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
	"bytes"
	"crypto/cipher"
	"encoding/binary"
	"errors"
	"slices"

	"github.com/nart4hire/goblockc/utils"
)

const BlockSize = 16

type Mode int

const (
	Encrypt Mode = iota
	Decrypt
)

type GoBlockC struct {
	key []byte
	keySchedule [][]byte
}


func NewBlock(key []byte) (cipher.Block, error) {
	if len(key) != 16 {
		return nil, errors.New("key must be 16 bytes long")
	}

	ks, err := utils.GetSubkeys(key)

	if err != nil {
		return nil, err
	}

	c := GoBlockC{
		key:			bytes.Clone(key),
		keySchedule:	ks,
	}

	return &c, nil
}

func (gbc *GoBlockC) BlockSize() int {
	return BlockSize
}

func (gbc *GoBlockC) Encrypt(dst, src []byte) {
	out, err := gbc.Parse(src, Encrypt)
	if err != nil {
		panic(err)
	}

	copy(dst, out)
}

func (gbc *GoBlockC) Decrypt(dst, src []byte) {
	out, err := gbc.Parse(src, Decrypt)
	if err != nil {
		panic(err)
	}
	copy(dst, out)
}

func (gbc *GoBlockC) Parse(block []byte, mode Mode) ([]byte, error) {
	if len(block) != 16 {
		return nil, errors.New("block must be 16 bytes long")
	}

	left, err := utils.BytesToUInt64(block[:8])
	if err != nil {
		return nil, err
	}

	right, err := utils.BytesToUInt64(block[8:])
	if err != nil {
		return nil, err
	}

	keySchedule := make([][]byte, 16)
	copy(keySchedule, gbc.keySchedule)

	if mode == Decrypt {
		slices.Reverse(keySchedule)
	}

	for i := range 16 {
		subkey, err := utils.BytesToUInt64(keySchedule[i])
		if err != nil {
			return nil, err
		}

		feistel, err := utils.Feistel(right, subkey, true)
		if err != nil {
			return nil, err
		}

		left, right = right, left^feistel
	}
	left, right = right, left
	leftBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(leftBytes, left)

	rightBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(rightBytes, right)

	return slices.Concat(leftBytes, rightBytes), nil
}
