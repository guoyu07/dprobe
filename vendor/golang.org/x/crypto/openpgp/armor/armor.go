// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package armor implements OpenPGP ASCII Armor, see RFC 4880. OpenPGP Armor is
// very similar to PEM except that it has an additional CRC checksum.
package armor // import "golang.org/x/crypto/openpgp/armor"

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"golang.org/x/crypto/openpgp/errors"
	"io"
)

// A Block represents an OpenPGP armored structure.
//
// The encoded form is:
//    -----BEGIN Type-----
//    Headers
//
//    base64-encoded Bytes
//    '=' base64 encoded checksum
//    -----END Type-----
// where Headers is a possibly empty sequence of Key: Value lines.
//
// Since the armored data can be very large, this package presents a streaming
// interface.
type Block struct ***REMOVED***
	Type    string            // The type, taken from the preamble (i.e. "PGP SIGNATURE").
	Header  map[string]string // Optional headers.
	Body    io.Reader         // A Reader from which the contents can be read
	lReader lineReader
	oReader openpgpReader
***REMOVED***

var ArmorCorrupt error = errors.StructuralError("armor invalid")

const crc24Init = 0xb704ce
const crc24Poly = 0x1864cfb
const crc24Mask = 0xffffff

// crc24 calculates the OpenPGP checksum as specified in RFC 4880, section 6.1
func crc24(crc uint32, d []byte) uint32 ***REMOVED***
	for _, b := range d ***REMOVED***
		crc ^= uint32(b) << 16
		for i := 0; i < 8; i++ ***REMOVED***
			crc <<= 1
			if crc&0x1000000 != 0 ***REMOVED***
				crc ^= crc24Poly
			***REMOVED***
		***REMOVED***
	***REMOVED***
	return crc
***REMOVED***

var armorStart = []byte("-----BEGIN ")
var armorEnd = []byte("-----END ")
var armorEndOfLine = []byte("-----")

// lineReader wraps a line based reader. It watches for the end of an armor
// block and records the expected CRC value.
type lineReader struct ***REMOVED***
	in  *bufio.Reader
	buf []byte
	eof bool
	crc uint32
***REMOVED***

func (l *lineReader) Read(p []byte) (n int, err error) ***REMOVED***
	if l.eof ***REMOVED***
		return 0, io.EOF
	***REMOVED***

	if len(l.buf) > 0 ***REMOVED***
		n = copy(p, l.buf)
		l.buf = l.buf[n:]
		return
	***REMOVED***

	line, isPrefix, err := l.in.ReadLine()
	if err != nil ***REMOVED***
		return
	***REMOVED***
	if isPrefix ***REMOVED***
		return 0, ArmorCorrupt
	***REMOVED***

	if len(line) == 5 && line[0] == '=' ***REMOVED***
		// This is the checksum line
		var expectedBytes [3]byte
		var m int
		m, err = base64.StdEncoding.Decode(expectedBytes[0:], line[1:])
		if m != 3 || err != nil ***REMOVED***
			return
		***REMOVED***
		l.crc = uint32(expectedBytes[0])<<16 |
			uint32(expectedBytes[1])<<8 |
			uint32(expectedBytes[2])

		line, _, err = l.in.ReadLine()
		if err != nil && err != io.EOF ***REMOVED***
			return
		***REMOVED***
		if !bytes.HasPrefix(line, armorEnd) ***REMOVED***
			return 0, ArmorCorrupt
		***REMOVED***

		l.eof = true
		return 0, io.EOF
	***REMOVED***

	if len(line) > 96 ***REMOVED***
		return 0, ArmorCorrupt
	***REMOVED***

	n = copy(p, line)
	bytesToSave := len(line) - n
	if bytesToSave > 0 ***REMOVED***
		if cap(l.buf) < bytesToSave ***REMOVED***
			l.buf = make([]byte, 0, bytesToSave)
		***REMOVED***
		l.buf = l.buf[0:bytesToSave]
		copy(l.buf, line[n:])
	***REMOVED***

	return
***REMOVED***

// openpgpReader passes Read calls to the underlying base64 decoder, but keeps
// a running CRC of the resulting data and checks the CRC against the value
// found by the lineReader at EOF.
type openpgpReader struct ***REMOVED***
	lReader    *lineReader
	b64Reader  io.Reader
	currentCRC uint32
***REMOVED***

func (r *openpgpReader) Read(p []byte) (n int, err error) ***REMOVED***
	n, err = r.b64Reader.Read(p)
	r.currentCRC = crc24(r.currentCRC, p[:n])

	if err == io.EOF ***REMOVED***
		if r.lReader.crc != uint32(r.currentCRC&crc24Mask) ***REMOVED***
			return 0, ArmorCorrupt
		***REMOVED***
	***REMOVED***

	return
***REMOVED***

// Decode reads a PGP armored block from the given Reader. It will ignore
// leading garbage. If it doesn't find a block, it will return nil, io.EOF. The
// given Reader is not usable after calling this function: an arbitrary amount
// of data may have been read past the end of the block.
func Decode(in io.Reader) (p *Block, err error) ***REMOVED***
	r := bufio.NewReaderSize(in, 100)
	var line []byte
	ignoreNext := false

TryNextBlock:
	p = nil

	// Skip leading garbage
	for ***REMOVED***
		ignoreThis := ignoreNext
		line, ignoreNext, err = r.ReadLine()
		if err != nil ***REMOVED***
			return
		***REMOVED***
		if ignoreNext || ignoreThis ***REMOVED***
			continue
		***REMOVED***
		line = bytes.TrimSpace(line)
		if len(line) > len(armorStart)+len(armorEndOfLine) && bytes.HasPrefix(line, armorStart) ***REMOVED***
			break
		***REMOVED***
	***REMOVED***

	p = new(Block)
	p.Type = string(line[len(armorStart) : len(line)-len(armorEndOfLine)])
	p.Header = make(map[string]string)
	nextIsContinuation := false
	var lastKey string

	// Read headers
	for ***REMOVED***
		isContinuation := nextIsContinuation
		line, nextIsContinuation, err = r.ReadLine()
		if err != nil ***REMOVED***
			p = nil
			return
		***REMOVED***
		if isContinuation ***REMOVED***
			p.Header[lastKey] += string(line)
			continue
		***REMOVED***
		line = bytes.TrimSpace(line)
		if len(line) == 0 ***REMOVED***
			break
		***REMOVED***

		i := bytes.Index(line, []byte(": "))
		if i == -1 ***REMOVED***
			goto TryNextBlock
		***REMOVED***
		lastKey = string(line[:i])
		p.Header[lastKey] = string(line[i+2:])
	***REMOVED***

	p.lReader.in = r
	p.oReader.currentCRC = crc24Init
	p.oReader.lReader = &p.lReader
	p.oReader.b64Reader = base64.NewDecoder(base64.StdEncoding, &p.lReader)
	p.Body = &p.oReader

	return
***REMOVED***