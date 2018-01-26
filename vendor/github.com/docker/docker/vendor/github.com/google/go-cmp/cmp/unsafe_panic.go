// Copyright 2017, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.md file.

// +build appengine js

package cmp

import "reflect"

const supportAllowUnexported = false

func unsafeRetrieveField(reflect.Value, reflect.StructField) reflect.Value ***REMOVED***
	panic("unsafeRetrieveField is not implemented")
***REMOVED***
