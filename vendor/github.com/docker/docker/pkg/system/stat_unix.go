// +build !windows

package system

import (
	"syscall"
)

// StatT type contains status of a file. It contains metadata
// like permission, owner, group, size, etc about a file.
type StatT struct ***REMOVED***
	mode uint32
	uid  uint32
	gid  uint32
	rdev uint64
	size int64
	mtim syscall.Timespec
***REMOVED***

// Mode returns file's permission mode.
func (s StatT) Mode() uint32 ***REMOVED***
	return s.mode
***REMOVED***

// UID returns file's user id of owner.
func (s StatT) UID() uint32 ***REMOVED***
	return s.uid
***REMOVED***

// GID returns file's group id of owner.
func (s StatT) GID() uint32 ***REMOVED***
	return s.gid
***REMOVED***

// Rdev returns file's device ID (if it's special file).
func (s StatT) Rdev() uint64 ***REMOVED***
	return s.rdev
***REMOVED***

// Size returns file's size.
func (s StatT) Size() int64 ***REMOVED***
	return s.size
***REMOVED***

// Mtim returns file's last modification time.
func (s StatT) Mtim() syscall.Timespec ***REMOVED***
	return s.mtim
***REMOVED***

// IsDir reports whether s describes a directory.
func (s StatT) IsDir() bool ***REMOVED***
	return s.mode&syscall.S_IFDIR != 0
***REMOVED***

// Stat takes a path to a file and returns
// a system.StatT type pertaining to that file.
//
// Throws an error if the file does not exist
func Stat(path string) (*StatT, error) ***REMOVED***
	s := &syscall.Stat_t***REMOVED******REMOVED***
	if err := syscall.Stat(path, s); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return fromStatT(s)
***REMOVED***
