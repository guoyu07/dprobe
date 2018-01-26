// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import (
	"syscall"
	"unsafe"
)

func setTimespec(sec, nsec int64) Timespec ***REMOVED***
	return Timespec***REMOVED***Sec: int32(sec), Nsec: int32(nsec)***REMOVED***
***REMOVED***

func setTimeval(sec, usec int64) Timeval ***REMOVED***
	return Timeval***REMOVED***Sec: int32(sec), Usec: int32(usec)***REMOVED***
***REMOVED***

//sysnb	gettimeofday(tp *Timeval) (sec int32, usec int32, err error)
func Gettimeofday(tv *Timeval) (err error) ***REMOVED***
	// The tv passed to gettimeofday must be non-nil
	// but is otherwise unused. The answers come back
	// in the two registers.
	sec, usec, err := gettimeofday(tv)
	tv.Sec = int32(sec)
	tv.Usec = int32(usec)
	return err
***REMOVED***

func SetKevent(k *Kevent_t, fd, mode, flags int) ***REMOVED***
	k.Ident = uint32(fd)
	k.Filter = int16(mode)
	k.Flags = uint16(flags)
***REMOVED***

func (iov *Iovec) SetLen(length int) ***REMOVED***
	iov.Len = uint32(length)
***REMOVED***

func (msghdr *Msghdr) SetControllen(length int) ***REMOVED***
	msghdr.Controllen = uint32(length)
***REMOVED***

func (cmsg *Cmsghdr) SetLen(length int) ***REMOVED***
	cmsg.Len = uint32(length)
***REMOVED***

func sendfile(outfd int, infd int, offset *int64, count int) (written int, err error) ***REMOVED***
	var length = uint64(count)

	_, _, e1 := Syscall9(SYS_SENDFILE, uintptr(infd), uintptr(outfd), uintptr(*offset), uintptr(*offset>>32), uintptr(unsafe.Pointer(&length)), 0, 0, 0, 0)

	written = int(length)

	if e1 != 0 ***REMOVED***
		err = e1
	***REMOVED***
	return
***REMOVED***

func Syscall9(num, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err syscall.Errno) // sic
