// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package secretbox encrypts and authenticates small messages.

Secretbox uses XSalsa20 and Poly1305 to encrypt and authenticate messages with
secret-key cryptography. The length of messages is not hidden.

It is the caller's responsibility to ensure the uniqueness of nonces—for
example, by using nonce 1 for the first message, nonce 2 for the second
message, etc. Nonces are long enough that randomly generated nonces have
negligible risk of collision.

This package is interoperable with NaCl: https://nacl.cr.yp.to/secretbox.html.
*/
package secretbox // import "golang.org/x/crypto/nacl/secretbox"

import (
	"golang.org/x/crypto/poly1305"
	"golang.org/x/crypto/salsa20/salsa"
)

// Overhead is the number of bytes of overhead when boxing a message.
const Overhead = poly1305.TagSize

// setup produces a sub-key and Salsa20 counter given a nonce and key.
func setup(subKey *[32]byte, counter *[16]byte, nonce *[24]byte, key *[32]byte) ***REMOVED***
	// We use XSalsa20 for encryption so first we need to generate a
	// key and nonce with HSalsa20.
	var hNonce [16]byte
	copy(hNonce[:], nonce[:])
	salsa.HSalsa20(subKey, &hNonce, key, &salsa.Sigma)

	// The final 8 bytes of the original nonce form the new nonce.
	copy(counter[:], nonce[16:])
***REMOVED***

// sliceForAppend takes a slice and a requested number of bytes. It returns a
// slice with the contents of the given slice followed by that many bytes and a
// second slice that aliases into it and contains only the extra bytes. If the
// original slice has sufficient capacity then no allocation is performed.
func sliceForAppend(in []byte, n int) (head, tail []byte) ***REMOVED***
	if total := len(in) + n; cap(in) >= total ***REMOVED***
		head = in[:total]
	***REMOVED*** else ***REMOVED***
		head = make([]byte, total)
		copy(head, in)
	***REMOVED***
	tail = head[len(in):]
	return
***REMOVED***

// Seal appends an encrypted and authenticated copy of message to out, which
// must not overlap message. The key and nonce pair must be unique for each
// distinct message and the output will be Overhead bytes longer than message.
func Seal(out, message []byte, nonce *[24]byte, key *[32]byte) []byte ***REMOVED***
	var subKey [32]byte
	var counter [16]byte
	setup(&subKey, &counter, nonce, key)

	// The Poly1305 key is generated by encrypting 32 bytes of zeros. Since
	// Salsa20 works with 64-byte blocks, we also generate 32 bytes of
	// keystream as a side effect.
	var firstBlock [64]byte
	salsa.XORKeyStream(firstBlock[:], firstBlock[:], &counter, &subKey)

	var poly1305Key [32]byte
	copy(poly1305Key[:], firstBlock[:])

	ret, out := sliceForAppend(out, len(message)+poly1305.TagSize)

	// We XOR up to 32 bytes of message with the keystream generated from
	// the first block.
	firstMessageBlock := message
	if len(firstMessageBlock) > 32 ***REMOVED***
		firstMessageBlock = firstMessageBlock[:32]
	***REMOVED***

	tagOut := out
	out = out[poly1305.TagSize:]
	for i, x := range firstMessageBlock ***REMOVED***
		out[i] = firstBlock[32+i] ^ x
	***REMOVED***
	message = message[len(firstMessageBlock):]
	ciphertext := out
	out = out[len(firstMessageBlock):]

	// Now encrypt the rest.
	counter[8] = 1
	salsa.XORKeyStream(out, message, &counter, &subKey)

	var tag [poly1305.TagSize]byte
	poly1305.Sum(&tag, ciphertext, &poly1305Key)
	copy(tagOut, tag[:])

	return ret
***REMOVED***

// Open authenticates and decrypts a box produced by Seal and appends the
// message to out, which must not overlap box. The output will be Overhead
// bytes smaller than box.
func Open(out []byte, box []byte, nonce *[24]byte, key *[32]byte) ([]byte, bool) ***REMOVED***
	if len(box) < Overhead ***REMOVED***
		return nil, false
	***REMOVED***

	var subKey [32]byte
	var counter [16]byte
	setup(&subKey, &counter, nonce, key)

	// The Poly1305 key is generated by encrypting 32 bytes of zeros. Since
	// Salsa20 works with 64-byte blocks, we also generate 32 bytes of
	// keystream as a side effect.
	var firstBlock [64]byte
	salsa.XORKeyStream(firstBlock[:], firstBlock[:], &counter, &subKey)

	var poly1305Key [32]byte
	copy(poly1305Key[:], firstBlock[:])
	var tag [poly1305.TagSize]byte
	copy(tag[:], box)

	if !poly1305.Verify(&tag, box[poly1305.TagSize:], &poly1305Key) ***REMOVED***
		return nil, false
	***REMOVED***

	ret, out := sliceForAppend(out, len(box)-Overhead)

	// We XOR up to 32 bytes of box with the keystream generated from
	// the first block.
	box = box[Overhead:]
	firstMessageBlock := box
	if len(firstMessageBlock) > 32 ***REMOVED***
		firstMessageBlock = firstMessageBlock[:32]
	***REMOVED***
	for i, x := range firstMessageBlock ***REMOVED***
		out[i] = firstBlock[32+i] ^ x
	***REMOVED***

	box = box[len(firstMessageBlock):]
	out = out[len(firstMessageBlock):]

	// Now decrypt the rest.
	counter[8] = 1
	salsa.XORKeyStream(out, box, &counter, &subKey)

	return ret, true
***REMOVED***
