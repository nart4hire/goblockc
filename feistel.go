package goblockc

import (
	"encoding/binary"
	"slices"

	"github.com/nart4hire/goblockc/utils"
)

func Feistel(right uint64, key uint64, forward bool) (uint64, error) {
	if forward {
		sbox := utils.GetSBox()
		bytes := make([]byte, 8)
		binary.BigEndian.PutUint64(bytes, right)

		for i := range 8 {
			bytes[i] = sbox[bytes[i]]
		}

		substitutedRight, err := utils.BytesToUInt64(bytes)
		if err != nil {
			return 0, err
		}

		return substitutedRight ^ key, nil
	} else {
		unXORedRight := right ^ key

		invsbox := utils.GetInvSBox()
		bytes := make([]byte, 8)
		binary.BigEndian.PutUint64(bytes, unXORedRight)

		for i := range 8 {
			bytes[i] = invsbox[bytes[i]]
		}

		unsubstitutedRight, err := utils.BytesToUInt64(bytes)
		if err != nil {
			return 0, err
		}

		return unsubstitutedRight, nil
	}
}

func Permute(left uint64, right uint64, forward bool) (uint64, uint64) {
	if forward {
		leftBytes := make([]byte, 8)
		rightBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(leftBytes, left)
		binary.BigEndian.PutUint64(rightBytes, right)

		leftBytesOne := make([]byte, 4)
		copy(leftBytesOne, leftBytes[:4])
		leftBytesTwo := make([]byte, 4)
		copy(leftBytesTwo, leftBytes[4:])
		rightBytesOne := make([]byte, 4)
		copy(rightBytesOne, rightBytes[:4])
		rightBytesTwo := make([]byte, 4)
		copy(rightBytesTwo, rightBytes[4:])

		result := slices.Concat(leftBytesOne, utils.Rotate(leftBytesTwo, 1, forward), utils.Rotate(rightBytesOne, 2, forward), utils.Rotate(rightBytesTwo, 3, forward))
		return binary.BigEndian.Uint64(result[:8]), binary.BigEndian.Uint64(result[8:])
	} else {
		leftBytes := make([]byte, 8)
		rightBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(leftBytes, left)
		binary.BigEndian.PutUint64(rightBytes, right)

		leftBytesOne := make([]byte, 4)
		copy(leftBytesOne, leftBytes[:4])
		leftBytesTwo := make([]byte, 4)
		copy(leftBytesTwo, leftBytes[4:])
		rightBytesOne := make([]byte, 4)
		copy(rightBytesOne, rightBytes[:4])
		rightBytesTwo := make([]byte, 4)
		copy(rightBytesTwo, rightBytes[4:])

		result := slices.Concat(leftBytesOne, utils.Rotate(leftBytesTwo, 1, forward), utils.Rotate(rightBytesOne, 2, forward), utils.Rotate(rightBytesTwo, 3, forward))
		return binary.BigEndian.Uint64(result[:8]), binary.BigEndian.Uint64(result[8:])
	}
}
