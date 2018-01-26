// Copyright 2015 CoreOS, Inc.
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

package capnslog

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var pid = os.Getpid()

type GlogFormatter struct ***REMOVED***
	StringFormatter
***REMOVED***

func NewGlogFormatter(w io.Writer) *GlogFormatter ***REMOVED***
	g := &GlogFormatter***REMOVED******REMOVED***
	g.w = bufio.NewWriter(w)
	return g
***REMOVED***

func (g GlogFormatter) Format(pkg string, level LogLevel, depth int, entries ...interface***REMOVED******REMOVED***) ***REMOVED***
	g.w.Write(GlogHeader(level, depth+1))
	g.StringFormatter.Format(pkg, level, depth+1, entries...)
***REMOVED***

func GlogHeader(level LogLevel, depth int) []byte ***REMOVED***
	// Lmmdd hh:mm:ss.uuuuuu threadid file:line]
	now := time.Now().UTC()
	_, file, line, ok := runtime.Caller(depth) // It's always the same number of frames to the user's call.
	if !ok ***REMOVED***
		file = "???"
		line = 1
	***REMOVED*** else ***REMOVED***
		slash := strings.LastIndex(file, "/")
		if slash >= 0 ***REMOVED***
			file = file[slash+1:]
		***REMOVED***
	***REMOVED***
	if line < 0 ***REMOVED***
		line = 0 // not a real line number
	***REMOVED***
	buf := &bytes.Buffer***REMOVED******REMOVED***
	buf.Grow(30)
	_, month, day := now.Date()
	hour, minute, second := now.Clock()
	buf.WriteString(level.Char())
	twoDigits(buf, int(month))
	twoDigits(buf, day)
	buf.WriteByte(' ')
	twoDigits(buf, hour)
	buf.WriteByte(':')
	twoDigits(buf, minute)
	buf.WriteByte(':')
	twoDigits(buf, second)
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(now.Nanosecond() / 1000))
	buf.WriteByte('Z')
	buf.WriteByte(' ')
	buf.WriteString(strconv.Itoa(pid))
	buf.WriteByte(' ')
	buf.WriteString(file)
	buf.WriteByte(':')
	buf.WriteString(strconv.Itoa(line))
	buf.WriteByte(']')
	buf.WriteByte(' ')
	return buf.Bytes()
***REMOVED***

const digits = "0123456789"

func twoDigits(b *bytes.Buffer, d int) ***REMOVED***
	c2 := digits[d%10]
	d /= 10
	c1 := digits[d%10]
	b.WriteByte(c1)
	b.WriteByte(c2)
***REMOVED***
