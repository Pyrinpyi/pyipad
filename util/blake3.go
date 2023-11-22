// Copyright (c) 2013-2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package util

import (
	"lukechampine.com/blake3"
)

// HashBlake2b calculates the hash blake2b(b).
func HashBlake3(buf []byte) []byte {
	var err error

	hasher := blake3.New(32, nil)
	_, err = hasher.Write(buf)
	if err != nil {
		return []byte{}
	}
	hashedBuf := hasher.Sum(nil)
	return hashedBuf
}
