// Copyright 2017, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.md file.

// Package diff implements an algorithm for producing edit-scripts.
// The edit-script is a sequence of operations needed to transform one list
// of symbols into another (or vice-versa). The edits allowed are insertions,
// deletions, and modifications. The summation of all edits is called the
// Levenshtein distance as this problem is well-known in computer science.
//
// This package prioritizes performance over accuracy. That is, the run time
// is more important than obtaining a minimal Levenshtein distance.
package diff

// EditType represents a single operation within an edit-script.
type EditType uint8

const (
	// Identity indicates that a symbol pair is identical in both list X and Y.
	Identity EditType = iota
	// UniqueX indicates that a symbol only exists in X and not Y.
	UniqueX
	// UniqueY indicates that a symbol only exists in Y and not X.
	UniqueY
	// Modified indicates that a symbol pair is a modification of each other.
	Modified
)

// EditScript represents the series of differences between two lists.
type EditScript []EditType

// String returns a human-readable string representing the edit-script where
// Identity, UniqueX, UniqueY, and Modified are represented by the
// '.', 'X', 'Y', and 'M' characters, respectively.
func (es EditScript) String() string ***REMOVED***
	b := make([]byte, len(es))
	for i, e := range es ***REMOVED***
		switch e ***REMOVED***
		case Identity:
			b[i] = '.'
		case UniqueX:
			b[i] = 'X'
		case UniqueY:
			b[i] = 'Y'
		case Modified:
			b[i] = 'M'
		default:
			panic("invalid edit-type")
		***REMOVED***
	***REMOVED***
	return string(b)
***REMOVED***

// stats returns a histogram of the number of each type of edit operation.
func (es EditScript) stats() (s struct***REMOVED*** NI, NX, NY, NM int ***REMOVED***) ***REMOVED***
	for _, e := range es ***REMOVED***
		switch e ***REMOVED***
		case Identity:
			s.NI++
		case UniqueX:
			s.NX++
		case UniqueY:
			s.NY++
		case Modified:
			s.NM++
		default:
			panic("invalid edit-type")
		***REMOVED***
	***REMOVED***
	return
***REMOVED***

// Dist is the Levenshtein distance and is guaranteed to be 0 if and only if
// lists X and Y are equal.
func (es EditScript) Dist() int ***REMOVED*** return len(es) - es.stats().NI ***REMOVED***

// LenX is the length of the X list.
func (es EditScript) LenX() int ***REMOVED*** return len(es) - es.stats().NY ***REMOVED***

// LenY is the length of the Y list.
func (es EditScript) LenY() int ***REMOVED*** return len(es) - es.stats().NX ***REMOVED***

// EqualFunc reports whether the symbols at indexes ix and iy are equal.
// When called by Difference, the index is guaranteed to be within nx and ny.
type EqualFunc func(ix int, iy int) Result

// Result is the result of comparison.
// NSame is the number of sub-elements that are equal.
// NDiff is the number of sub-elements that are not equal.
type Result struct***REMOVED*** NSame, NDiff int ***REMOVED***

// Equal indicates whether the symbols are equal. Two symbols are equal
// if and only if NDiff == 0. If Equal, then they are also Similar.
func (r Result) Equal() bool ***REMOVED*** return r.NDiff == 0 ***REMOVED***

// Similar indicates whether two symbols are similar and may be represented
// by using the Modified type. As a special case, we consider binary comparisons
// (i.e., those that return Result***REMOVED***1, 0***REMOVED*** or Result***REMOVED***0, 1***REMOVED***) to be similar.
//
// The exact ratio of NSame to NDiff to determine similarity may change.
func (r Result) Similar() bool ***REMOVED***
	// Use NSame+1 to offset NSame so that binary comparisons are similar.
	return r.NSame+1 >= r.NDiff
***REMOVED***

// Difference reports whether two lists of lengths nx and ny are equal
// given the definition of equality provided as f.
//
// This function may return a edit-script, which is a sequence of operations
// needed to convert one list into the other. If non-nil, the following
// invariants for the edit-script are maintained:
//	• eq == (es.Dist()==0)
//	• nx == es.LenX()
//	• ny == es.LenY()
//
// This algorithm is not guaranteed to be an optimal solution (i.e., one that
// produces an edit-script with a minimal Levenshtein distance). This algorithm
// favors performance over optimality. The exact output is not guaranteed to
// be stable and may change over time.
func Difference(nx, ny int, f EqualFunc) (eq bool, es EditScript) ***REMOVED***
	es = searchGraph(nx, ny, f)
	st := es.stats()
	eq = len(es) == st.NI
	if !eq && st.NI < (nx+ny)/4 ***REMOVED***
		return eq, nil // Edit-script more distracting than helpful
	***REMOVED***
	return eq, es
***REMOVED***

func searchGraph(nx, ny int, f EqualFunc) EditScript ***REMOVED***
	// This algorithm is based on traversing what is known as an "edit-graph".
	// See Figure 1 from "An O(ND) Difference Algorithm and Its Variations"
	// by Eugene W. Myers. Since D can be as large as N itself, this is
	// effectively O(N^2). Unlike the algorithm from that paper, we are not
	// interested in the optimal path, but at least some "decent" path.
	//
	// For example, let X and Y be lists of symbols:
	//	X = [A B C A B B A]
	//	Y = [C B A B A C]
	//
	// The edit-graph can be drawn as the following:
	//	   A B C A B B A
	//	  ┌─────────────┐
	//	C │_|_|\|_|_|_|_│ 0
	//	B │_|\|_|_|\|\|_│ 1
	//	A │\|_|_|\|_|_|\│ 2
	//	B │_|\|_|_|\|\|_│ 3
	//	A │\|_|_|\|_|_|\│ 4
	//	C │ | |\| | | | │ 5
	//	  └─────────────┘ 6
	//	   0 1 2 3 4 5 6 7
	//
	// List X is written along the horizontal axis, while list Y is written
	// along the vertical axis. At any point on this grid, if the symbol in
	// list X matches the corresponding symbol in list Y, then a '\' is drawn.
	// The goal of any minimal edit-script algorithm is to find a path from the
	// top-left corner to the bottom-right corner, while traveling through the
	// fewest horizontal or vertical edges.
	// A horizontal edge is equivalent to inserting a symbol from list X.
	// A vertical edge is equivalent to inserting a symbol from list Y.
	// A diagonal edge is equivalent to a matching symbol between both X and Y.

	// Invariants:
	//	• 0 ≤ fwdPath.X ≤ (fwdFrontier.X, revFrontier.X) ≤ revPath.X ≤ nx
	//	• 0 ≤ fwdPath.Y ≤ (fwdFrontier.Y, revFrontier.Y) ≤ revPath.Y ≤ ny
	//
	// In general:
	//	• fwdFrontier.X < revFrontier.X
	//	• fwdFrontier.Y < revFrontier.Y
	// Unless, it is time for the algorithm to terminate.
	fwdPath := path***REMOVED***+1, point***REMOVED***0, 0***REMOVED***, make(EditScript, 0, (nx+ny)/2)***REMOVED***
	revPath := path***REMOVED***-1, point***REMOVED***nx, ny***REMOVED***, make(EditScript, 0)***REMOVED***
	fwdFrontier := fwdPath.point // Forward search frontier
	revFrontier := revPath.point // Reverse search frontier

	// Search budget bounds the cost of searching for better paths.
	// The longest sequence of non-matching symbols that can be tolerated is
	// approximately the square-root of the search budget.
	searchBudget := 4 * (nx + ny) // O(n)

	// The algorithm below is a greedy, meet-in-the-middle algorithm for
	// computing sub-optimal edit-scripts between two lists.
	//
	// The algorithm is approximately as follows:
	//	• Searching for differences switches back-and-forth between
	//	a search that starts at the beginning (the top-left corner), and
	//	a search that starts at the end (the bottom-right corner). The goal of
	//	the search is connect with the search from the opposite corner.
	//	• As we search, we build a path in a greedy manner, where the first
	//	match seen is added to the path (this is sub-optimal, but provides a
	//	decent result in practice). When matches are found, we try the next pair
	//	of symbols in the lists and follow all matches as far as possible.
	//	• When searching for matches, we search along a diagonal going through
	//	through the "frontier" point. If no matches are found, we advance the
	//	frontier towards the opposite corner.
	//	• This algorithm terminates when either the X coordinates or the
	//	Y coordinates of the forward and reverse frontier points ever intersect.
	//
	// This algorithm is correct even if searching only in the forward direction
	// or in the reverse direction. We do both because it is commonly observed
	// that two lists commonly differ because elements were added to the front
	// or end of the other list.
	//
	// Running the tests with the "debug" build tag prints a visualization of
	// the algorithm running in real-time. This is educational for understanding
	// how the algorithm works. See debug_enable.go.
	f = debug.Begin(nx, ny, f, &fwdPath.es, &revPath.es)
	for ***REMOVED***
		// Forward search from the beginning.
		if fwdFrontier.X >= revFrontier.X || fwdFrontier.Y >= revFrontier.Y || searchBudget == 0 ***REMOVED***
			break
		***REMOVED***
		for stop1, stop2, i := false, false, 0; !(stop1 && stop2) && searchBudget > 0; i++ ***REMOVED***
			// Search in a diagonal pattern for a match.
			z := zigzag(i)
			p := point***REMOVED***fwdFrontier.X + z, fwdFrontier.Y - z***REMOVED***
			switch ***REMOVED***
			case p.X >= revPath.X || p.Y < fwdPath.Y:
				stop1 = true // Hit top-right corner
			case p.Y >= revPath.Y || p.X < fwdPath.X:
				stop2 = true // Hit bottom-left corner
			case f(p.X, p.Y).Equal():
				// Match found, so connect the path to this point.
				fwdPath.connect(p, f)
				fwdPath.append(Identity)
				// Follow sequence of matches as far as possible.
				for fwdPath.X < revPath.X && fwdPath.Y < revPath.Y ***REMOVED***
					if !f(fwdPath.X, fwdPath.Y).Equal() ***REMOVED***
						break
					***REMOVED***
					fwdPath.append(Identity)
				***REMOVED***
				fwdFrontier = fwdPath.point
				stop1, stop2 = true, true
			default:
				searchBudget-- // Match not found
			***REMOVED***
			debug.Update()
		***REMOVED***
		// Advance the frontier towards reverse point.
		if revPath.X-fwdFrontier.X >= revPath.Y-fwdFrontier.Y ***REMOVED***
			fwdFrontier.X++
		***REMOVED*** else ***REMOVED***
			fwdFrontier.Y++
		***REMOVED***

		// Reverse search from the end.
		if fwdFrontier.X >= revFrontier.X || fwdFrontier.Y >= revFrontier.Y || searchBudget == 0 ***REMOVED***
			break
		***REMOVED***
		for stop1, stop2, i := false, false, 0; !(stop1 && stop2) && searchBudget > 0; i++ ***REMOVED***
			// Search in a diagonal pattern for a match.
			z := zigzag(i)
			p := point***REMOVED***revFrontier.X - z, revFrontier.Y + z***REMOVED***
			switch ***REMOVED***
			case fwdPath.X >= p.X || revPath.Y < p.Y:
				stop1 = true // Hit bottom-left corner
			case fwdPath.Y >= p.Y || revPath.X < p.X:
				stop2 = true // Hit top-right corner
			case f(p.X-1, p.Y-1).Equal():
				// Match found, so connect the path to this point.
				revPath.connect(p, f)
				revPath.append(Identity)
				// Follow sequence of matches as far as possible.
				for fwdPath.X < revPath.X && fwdPath.Y < revPath.Y ***REMOVED***
					if !f(revPath.X-1, revPath.Y-1).Equal() ***REMOVED***
						break
					***REMOVED***
					revPath.append(Identity)
				***REMOVED***
				revFrontier = revPath.point
				stop1, stop2 = true, true
			default:
				searchBudget-- // Match not found
			***REMOVED***
			debug.Update()
		***REMOVED***
		// Advance the frontier towards forward point.
		if revFrontier.X-fwdPath.X >= revFrontier.Y-fwdPath.Y ***REMOVED***
			revFrontier.X--
		***REMOVED*** else ***REMOVED***
			revFrontier.Y--
		***REMOVED***
	***REMOVED***

	// Join the forward and reverse paths and then append the reverse path.
	fwdPath.connect(revPath.point, f)
	for i := len(revPath.es) - 1; i >= 0; i-- ***REMOVED***
		t := revPath.es[i]
		revPath.es = revPath.es[:i]
		fwdPath.append(t)
	***REMOVED***
	debug.Finish()
	return fwdPath.es
***REMOVED***

type path struct ***REMOVED***
	dir   int // +1 if forward, -1 if reverse
	point     // Leading point of the EditScript path
	es    EditScript
***REMOVED***

// connect appends any necessary Identity, Modified, UniqueX, or UniqueY types
// to the edit-script to connect p.point to dst.
func (p *path) connect(dst point, f EqualFunc) ***REMOVED***
	if p.dir > 0 ***REMOVED***
		// Connect in forward direction.
		for dst.X > p.X && dst.Y > p.Y ***REMOVED***
			switch r := f(p.X, p.Y); ***REMOVED***
			case r.Equal():
				p.append(Identity)
			case r.Similar():
				p.append(Modified)
			case dst.X-p.X >= dst.Y-p.Y:
				p.append(UniqueX)
			default:
				p.append(UniqueY)
			***REMOVED***
		***REMOVED***
		for dst.X > p.X ***REMOVED***
			p.append(UniqueX)
		***REMOVED***
		for dst.Y > p.Y ***REMOVED***
			p.append(UniqueY)
		***REMOVED***
	***REMOVED*** else ***REMOVED***
		// Connect in reverse direction.
		for p.X > dst.X && p.Y > dst.Y ***REMOVED***
			switch r := f(p.X-1, p.Y-1); ***REMOVED***
			case r.Equal():
				p.append(Identity)
			case r.Similar():
				p.append(Modified)
			case p.Y-dst.Y >= p.X-dst.X:
				p.append(UniqueY)
			default:
				p.append(UniqueX)
			***REMOVED***
		***REMOVED***
		for p.X > dst.X ***REMOVED***
			p.append(UniqueX)
		***REMOVED***
		for p.Y > dst.Y ***REMOVED***
			p.append(UniqueY)
		***REMOVED***
	***REMOVED***
***REMOVED***

func (p *path) append(t EditType) ***REMOVED***
	p.es = append(p.es, t)
	switch t ***REMOVED***
	case Identity, Modified:
		p.add(p.dir, p.dir)
	case UniqueX:
		p.add(p.dir, 0)
	case UniqueY:
		p.add(0, p.dir)
	***REMOVED***
	debug.Update()
***REMOVED***

type point struct***REMOVED*** X, Y int ***REMOVED***

func (p *point) add(dx, dy int) ***REMOVED*** p.X += dx; p.Y += dy ***REMOVED***

// zigzag maps a consecutive sequence of integers to a zig-zag sequence.
//	[0 1 2 3 4 5 ...] => [0 -1 +1 -2 +2 ...]
func zigzag(x int) int ***REMOVED***
	if x&1 != 0 ***REMOVED***
		x = ^x
	***REMOVED***
	return x >> 1
***REMOVED***
