// +build solaris,!cgo

//
// Implementing the functions below requires cgo support.  Non-cgo stubs
// versions are defined below to enable cross-compilation of source code
// that depends on these functions, but the resultant cross-compiled
// binaries cannot actually be used.  If the stub function(s) below are
// actually invoked they will cause the calling process to exit.
//

package devices

func getmajor(dev uint64) uint64 ***REMOVED***
	panic("getmajor() support requires cgo.")
***REMOVED***

func getminor(dev uint64) uint64 ***REMOVED***
	panic("getminor() support requires cgo.")
***REMOVED***

func makedev(major int, minor int) int ***REMOVED***
	panic("makedev() support requires cgo.")
***REMOVED***
