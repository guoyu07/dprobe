// Copyright 2015 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fileutil

import (
	"os"
	"syscall"
	"time"
)

func TryLockFile(path string, flag int, perm os.FileMode) (*LockedFile, error) ***REMOVED***
	if err := os.Chmod(path, syscall.DMEXCL|PrivateFileMode); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	f, err := os.Open(path, flag, perm)
	if err != nil ***REMOVED***
		return nil, ErrLocked
	***REMOVED***
	return &LockedFile***REMOVED***f***REMOVED***, nil
***REMOVED***

func LockFile(path string, flag int, perm os.FileMode) (*LockedFile, error) ***REMOVED***
	if err := os.Chmod(path, syscall.DMEXCL|PrivateFileMode); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	for ***REMOVED***
		f, err := os.OpenFile(path, flag, perm)
		if err == nil ***REMOVED***
			return &LockedFile***REMOVED***f***REMOVED***, nil
		***REMOVED***
		time.Sleep(10 * time.Millisecond)
	***REMOVED***
***REMOVED***
