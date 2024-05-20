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

// Manual implementation of Cipher Block Chaining (CFB) as described in NIST SP 800-38A, pp 10-11

import (
	"crypto/cipher"
	"errors"

	"github.com/nart4hire/goblockc/utils"
)

type cfb struct {
	mode		Mode
	block		cipher.Block
	used		int
	next		[]byte
	out			[]byte
}

func NewCFB(block cipher.Block, initVector []byte, mode Mode) (cipher.Stream, error) {
	if len(initVector) != block.BlockSize() {
		return nil, errors.New("IV length must equal block size")
	}

	stream := &cfb{
		mode:		mode,
		block:		block,
		used:		block.BlockSize(),
		next:		make([]byte, block.BlockSize()),
		out:		make([]byte, block.BlockSize()),
	}
	
	copy(stream.next, initVector)

	return stream, nil
}

func (c *cfb) XORKeyStream(dst, src []byte) {
	if len(dst) < len(src) {
		panic("Output buffer smaller than input")
	}

	result := make([]byte, len(dst))
	accum := 0

	for len(src) > 0 {
		if c.used == len(c.out) {
			c.block.Encrypt(c.out, c.next)
			c.used = 0
		}

		if c.mode == Decrypt {
			copy(c.next[c.used:], src)
		}

		n := utils.XORBytes(result[accum:], src, c.out[c.used:])

		if c.mode == Encrypt {
			copy(c.next[c.used:], result[accum:])
		}

		src = src[n:]
		c.used += n
		accum += n
	}

	copy(dst, result)
}