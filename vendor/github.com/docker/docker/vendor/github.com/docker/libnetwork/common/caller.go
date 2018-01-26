package common

import (
	"runtime"
	"strings"
)

func callerInfo(i int) string ***REMOVED***
	ptr, _, _, ok := runtime.Caller(i)
	fName := "unknown"
	if ok ***REMOVED***
		f := runtime.FuncForPC(ptr)
		if f != nil ***REMOVED***
			// f.Name() is like: github.com/docker/libnetwork/common.MethodName
			tmp := strings.Split(f.Name(), ".")
			if len(tmp) > 0 ***REMOVED***
				fName = tmp[len(tmp)-1]
			***REMOVED***
		***REMOVED***
	***REMOVED***

	return fName
***REMOVED***

// CallerName returns the name of the function at the specified level
// level == 0 means current method name
func CallerName(level int) string ***REMOVED***
	return callerInfo(2 + level)
***REMOVED***
