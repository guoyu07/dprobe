// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin dragonfly freebsd linux netbsd openbsd solaris

package unix

func itoa(val int) string ***REMOVED*** // do it here rather than with fmt to avoid dependency
	if val < 0 ***REMOVED***
		return "-" + uitoa(uint(-val))
	***REMOVED***
	return uitoa(uint(val))
***REMOVED***

func uitoa(val uint) string ***REMOVED***
	var buf [32]byte // big enough for int64
	i := len(buf) - 1
	for val >= 10 ***REMOVED***
		buf[i] = byte(val%10 + '0')
		i--
		val /= 10
	***REMOVED***
	buf[i] = byte(val + '0')
	return string(buf[i:])
***REMOVED***
