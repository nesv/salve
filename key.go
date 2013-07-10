package main

/*
This file, key.go, implements any structs and functions related to working with
the hashing and locating of keys on the partition ring.
*/

import (
	"errors"
	"encoding/binary"
	"crypto/sha1"
	"io"
)

var (
	ErrBadHashSize = errors.New("hash is the wrong size")
)

func BlurKeyHash(hash []byte) (i uint64, err error) {
	if len(hash) != 20 {
		err = ErrBadHashSize
		return
	}

	a, _ := binary.Uvarint(hash[0:8])
	b, _ := binary.Uvarint(hash[8:16])
	c, _ := binary.Uvarint(hash[16:20])

	i = a ^ b ^ c
	return
}

