// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

// This code is shared between the main code generator and the test code.

import (
	"flag"
	"log"
	"strconv"
	"strings"

	"golang.org/x/text/internal/gen"
	"golang.org/x/text/internal/ucd"
)

var (
	outputFile = flag.String("out", "tables.go", "output file")
)

var typeMap = map[string]elem***REMOVED***
	"A":  tagAmbiguous,
	"N":  tagNeutral,
	"Na": tagNarrow,
	"W":  tagWide,
	"F":  tagFullwidth,
	"H":  tagHalfwidth,
***REMOVED***

// getWidthData calls f for every entry for which it is defined.
//
// f may be called multiple times for the same rune. The last call to f is the
// correct value. f is not called for all runes. The default tag type is
// Neutral.
func getWidthData(f func(r rune, tag elem, alt rune)) ***REMOVED***
	// Set the default values for Unified Ideographs. In line with Annex 11,
	// we encode full ranges instead of the defined runes in Unified_Ideograph.
	for _, b := range []struct***REMOVED*** lo, hi rune ***REMOVED******REMOVED***
		***REMOVED***0x4E00, 0x9FFF***REMOVED***,   // the CJK Unified Ideographs block,
		***REMOVED***0x3400, 0x4DBF***REMOVED***,   // the CJK Unified Ideographs Externsion A block,
		***REMOVED***0xF900, 0xFAFF***REMOVED***,   // the CJK Compatibility Ideographs block,
		***REMOVED***0x20000, 0x2FFFF***REMOVED***, // the Supplementary Ideographic Plane,
		***REMOVED***0x30000, 0x3FFFF***REMOVED***, // the Tertiary Ideographic Plane,
	***REMOVED*** ***REMOVED***
		for r := b.lo; r <= b.hi; r++ ***REMOVED***
			f(r, tagWide, 0)
		***REMOVED***
	***REMOVED***

	inverse := map[rune]rune***REMOVED******REMOVED***
	maps := map[string]bool***REMOVED***
		"<wide>":   true,
		"<narrow>": true,
	***REMOVED***

	// We cannot reuse package norm's decomposition, as we need an unexpanded
	// decomposition. We make use of the opportunity to verify that the
	// decomposition type is as expected.
	ucd.Parse(gen.OpenUCDFile("UnicodeData.txt"), func(p *ucd.Parser) ***REMOVED***
		r := p.Rune(0)
		s := strings.SplitN(p.String(ucd.DecompMapping), " ", 2)
		if !maps[s[0]] ***REMOVED***
			return
		***REMOVED***
		x, err := strconv.ParseUint(s[1], 16, 32)
		if err != nil ***REMOVED***
			log.Fatalf("Error parsing rune %q", s[1])
		***REMOVED***
		if inverse[r] != 0 || inverse[rune(x)] != 0 ***REMOVED***
			log.Fatalf("Circular dependency in mapping between %U and %U", r, x)
		***REMOVED***
		inverse[r] = rune(x)
		inverse[rune(x)] = r
	***REMOVED***)

	// <rune range>;<type>
	ucd.Parse(gen.OpenUCDFile("EastAsianWidth.txt"), func(p *ucd.Parser) ***REMOVED***
		tag, ok := typeMap[p.String(1)]
		if !ok ***REMOVED***
			log.Fatalf("Unknown width type %q", p.String(1))
		***REMOVED***
		r := p.Rune(0)
		alt, ok := inverse[r]
		if tag == tagFullwidth || tag == tagHalfwidth && r != wonSign ***REMOVED***
			tag |= tagNeedsFold
			if !ok ***REMOVED***
				log.Fatalf("Narrow or wide rune %U has no decomposition", r)
			***REMOVED***
		***REMOVED***
		f(r, tag, alt)
	***REMOVED***)
***REMOVED***