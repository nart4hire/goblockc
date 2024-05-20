package goblockc

import (
	"crypto/cipher"
	"errors"

	"github.com/nart4hire/goblockc/utils"
)

type ctr struct {
	block 		cipher.Block
	used		int
	counter		[]byte
	out			[]byte
}

func NewCTR(block cipher.Block, initVector []byte) (cipher.Stream, error) {
	if len(initVector) != block.BlockSize() {
		return nil, errors.New("IV length must equal block size")
	}

	stream := &ctr{
		block:		block,
		used:		block.BlockSize(),
		counter:	make([]byte, block.BlockSize()),
		out:		make([]byte, block.BlockSize()),
	}

	copy(stream.counter, initVector)

	return stream, nil
}

func (c *ctr) XORKeyStream(dst, src []byte) {
	if len(dst) < len(src) {
		panic("Output buffer smaller than input")
	}

	result := make([]byte, len(dst))
	accum := 0

	for len(src) > 0 {
		if c.used == len(c.out) {
			c.block.Encrypt(c.out, c.counter)
			c.used = 0
			
			for i := len(c.counter) - 1; i >= 0; i-- {
				c.counter[i]++
				if c.counter[i] != 0 {
					break
				}
			}
		}
		n := utils.XORBytes(result[accum:], src, c.out[c.used:])
		src = src[n:]
		c.used += n
		accum += n
	}

	copy(dst, result)
}