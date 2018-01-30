// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.9

package blake2s

import (
	"crypto"
	"hash"
)

func init() ***REMOVED***
	newHash256 := func() hash.Hash ***REMOVED***
		h, _ := New256(nil)
		return h
	***REMOVED***

	crypto.RegisterHash(crypto.BLAKE2s_256, newHash256)
***REMOVED***