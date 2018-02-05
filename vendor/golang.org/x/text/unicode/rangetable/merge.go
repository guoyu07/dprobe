// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rangetable

import (
	"unicode"
)

// atEnd is used to mark a completed iteration.
const atEnd = unicode.MaxRune + 1

// Merge returns a new RangeTable that is the union of the given tables.
// It can also be used to compact user-created RangeTables. The entries in
// R16 and R32 for any given RangeTable should be sorted and non-overlapping.
//
// A lookup in the resulting table can be several times faster than using In
// directly on the ranges. Merge is an expensive operation, however, and only
// makes sense if one intends to use the result for more than a couple of
// hundred lookups.
func Merge(ranges ...*unicode.RangeTable) *unicode.RangeTable ***REMOVED***
	rt := &unicode.RangeTable***REMOVED******REMOVED***
	if len(ranges) == 0 ***REMOVED***
		return rt
	***REMOVED***

	iter := tablesIter(make([]tableIndex, len(ranges)))

	for i, t := range ranges ***REMOVED***
		iter[i] = tableIndex***REMOVED***t, 0, atEnd***REMOVED***
		if len(t.R16) > 0 ***REMOVED***
			iter[i].next = rune(t.R16[0].Lo)
		***REMOVED***
	***REMOVED***

	if r0 := iter.next16(); r0.Stride != 0 ***REMOVED***
		for ***REMOVED***
			r1 := iter.next16()
			if r1.Stride == 0 ***REMOVED***
				rt.R16 = append(rt.R16, r0)
				break
			***REMOVED***
			stride := r1.Lo - r0.Hi
			if (r1.Lo == r1.Hi || stride == r1.Stride) && (r0.Lo == r0.Hi || stride == r0.Stride) ***REMOVED***
				// Fully merge the next range into the previous one.
				r0.Hi, r0.Stride = r1.Hi, stride
				continue
			***REMOVED*** else if stride == r0.Stride ***REMOVED***
				// Move the first element of r1 to r0. This may eliminate an
				// entry.
				r0.Hi = r1.Lo
				r0.Stride = stride
				r1.Lo = r1.Lo + r1.Stride
				if r1.Lo > r1.Hi ***REMOVED***
					continue
				***REMOVED***
			***REMOVED***
			rt.R16 = append(rt.R16, r0)
			r0 = r1
		***REMOVED***
	***REMOVED***

	for i, t := range ranges ***REMOVED***
		iter[i] = tableIndex***REMOVED***t, 0, atEnd***REMOVED***
		if len(t.R32) > 0 ***REMOVED***
			iter[i].next = rune(t.R32[0].Lo)
		***REMOVED***
	***REMOVED***

	if r0 := iter.next32(); r0.Stride != 0 ***REMOVED***
		for ***REMOVED***
			r1 := iter.next32()
			if r1.Stride == 0 ***REMOVED***
				rt.R32 = append(rt.R32, r0)
				break
			***REMOVED***
			stride := r1.Lo - r0.Hi
			if (r1.Lo == r1.Hi || stride == r1.Stride) && (r0.Lo == r0.Hi || stride == r0.Stride) ***REMOVED***
				// Fully merge the next range into the previous one.
				r0.Hi, r0.Stride = r1.Hi, stride
				continue
			***REMOVED*** else if stride == r0.Stride ***REMOVED***
				// Move the first element of r1 to r0. This may eliminate an
				// entry.
				r0.Hi = r1.Lo
				r1.Lo = r1.Lo + r1.Stride
				if r1.Lo > r1.Hi ***REMOVED***
					continue
				***REMOVED***
			***REMOVED***
			rt.R32 = append(rt.R32, r0)
			r0 = r1
		***REMOVED***
	***REMOVED***

	for i := 0; i < len(rt.R16) && rt.R16[i].Hi <= unicode.MaxLatin1; i++ ***REMOVED***
		rt.LatinOffset = i + 1
	***REMOVED***

	return rt
***REMOVED***

type tableIndex struct ***REMOVED***
	t    *unicode.RangeTable
	p    uint32
	next rune
***REMOVED***

type tablesIter []tableIndex

// sortIter does an insertion sort using the next field of tableIndex. Insertion
// sort is a good sorting algorithm for this case.
func sortIter(t []tableIndex) ***REMOVED***
	for i := range t ***REMOVED***
		for j := i; j > 0 && t[j-1].next > t[j].next; j-- ***REMOVED***
			t[j], t[j-1] = t[j-1], t[j]
		***REMOVED***
	***REMOVED***
***REMOVED***

// next16 finds the ranged to be added to the table. If ranges overlap between
// multiple tables it clips the result to a non-overlapping range if the
// elements are not fully subsumed. It returns a zero range if there are no more
// ranges.
func (ti tablesIter) next16() unicode.Range16 ***REMOVED***
	sortIter(ti)

	t0 := ti[0]
	if t0.next == atEnd ***REMOVED***
		return unicode.Range16***REMOVED******REMOVED***
	***REMOVED***
	r0 := t0.t.R16[t0.p]
	r0.Lo = uint16(t0.next)

	// We restrict the Hi of the current range if it overlaps with another range.
	for i := range ti ***REMOVED***
		tn := ti[i]
		// Since our tableIndices are sorted by next, we can break if the there
		// is no overlap. The first value of a next range can always be merged
		// into the current one, so we can break in case of equality as well.
		if rune(r0.Hi) <= tn.next ***REMOVED***
			break
		***REMOVED***
		rn := tn.t.R16[tn.p]
		rn.Lo = uint16(tn.next)

		// Limit r0.Hi based on next ranges in list, but allow it to overlap
		// with ranges as long as it subsumes it.
		m := (rn.Lo - r0.Lo) % r0.Stride
		if m == 0 && (rn.Stride == r0.Stride || rn.Lo == rn.Hi) ***REMOVED***
			// Overlap, take the min of the two Hi values: for simplicity's sake
			// we only process one range at a time.
			if r0.Hi > rn.Hi ***REMOVED***
				r0.Hi = rn.Hi
			***REMOVED***
		***REMOVED*** else ***REMOVED***
			// Not a compatible stride. Set to the last possible value before
			// rn.Lo, but ensure there is at least one value.
			if x := rn.Lo - m; r0.Lo <= x ***REMOVED***
				r0.Hi = x
			***REMOVED***
			break
		***REMOVED***
	***REMOVED***

	// Update the next values for each table.
	for i := range ti ***REMOVED***
		tn := &ti[i]
		if rune(r0.Hi) < tn.next ***REMOVED***
			break
		***REMOVED***
		rn := tn.t.R16[tn.p]
		stride := rune(rn.Stride)
		tn.next += stride * (1 + ((rune(r0.Hi) - tn.next) / stride))
		if rune(rn.Hi) < tn.next ***REMOVED***
			if tn.p++; int(tn.p) == len(tn.t.R16) ***REMOVED***
				tn.next = atEnd
			***REMOVED*** else ***REMOVED***
				tn.next = rune(tn.t.R16[tn.p].Lo)
			***REMOVED***
		***REMOVED***
	***REMOVED***

	if r0.Lo == r0.Hi ***REMOVED***
		r0.Stride = 1
	***REMOVED***

	return r0
***REMOVED***

// next32 finds the ranged to be added to the table. If ranges overlap between
// multiple tables it clips the result to a non-overlapping range if the
// elements are not fully subsumed. It returns a zero range if there are no more
// ranges.
func (ti tablesIter) next32() unicode.Range32 ***REMOVED***
	sortIter(ti)

	t0 := ti[0]
	if t0.next == atEnd ***REMOVED***
		return unicode.Range32***REMOVED******REMOVED***
	***REMOVED***
	r0 := t0.t.R32[t0.p]
	r0.Lo = uint32(t0.next)

	// We restrict the Hi of the current range if it overlaps with another range.
	for i := range ti ***REMOVED***
		tn := ti[i]
		// Since our tableIndices are sorted by next, we can break if the there
		// is no overlap. The first value of a next range can always be merged
		// into the current one, so we can break in case of equality as well.
		if rune(r0.Hi) <= tn.next ***REMOVED***
			break
		***REMOVED***
		rn := tn.t.R32[tn.p]
		rn.Lo = uint32(tn.next)

		// Limit r0.Hi based on next ranges in list, but allow it to overlap
		// with ranges as long as it subsumes it.
		m := (rn.Lo - r0.Lo) % r0.Stride
		if m == 0 && (rn.Stride == r0.Stride || rn.Lo == rn.Hi) ***REMOVED***
			// Overlap, take the min of the two Hi values: for simplicity's sake
			// we only process one range at a time.
			if r0.Hi > rn.Hi ***REMOVED***
				r0.Hi = rn.Hi
			***REMOVED***
		***REMOVED*** else ***REMOVED***
			// Not a compatible stride. Set to the last possible value before
			// rn.Lo, but ensure there is at least one value.
			if x := rn.Lo - m; r0.Lo <= x ***REMOVED***
				r0.Hi = x
			***REMOVED***
			break
		***REMOVED***
	***REMOVED***

	// Update the next values for each table.
	for i := range ti ***REMOVED***
		tn := &ti[i]
		if rune(r0.Hi) < tn.next ***REMOVED***
			break
		***REMOVED***
		rn := tn.t.R32[tn.p]
		stride := rune(rn.Stride)
		tn.next += stride * (1 + ((rune(r0.Hi) - tn.next) / stride))
		if rune(rn.Hi) < tn.next ***REMOVED***
			if tn.p++; int(tn.p) == len(tn.t.R32) ***REMOVED***
				tn.next = atEnd
			***REMOVED*** else ***REMOVED***
				tn.next = rune(tn.t.R32[tn.p].Lo)
			***REMOVED***
		***REMOVED***
	***REMOVED***

	if r0.Lo == r0.Hi ***REMOVED***
		r0.Stride = 1
	***REMOVED***

	return r0
***REMOVED***