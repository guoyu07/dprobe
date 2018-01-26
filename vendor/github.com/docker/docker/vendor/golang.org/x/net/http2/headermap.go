// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http2

import (
	"net/http"
	"strings"
)

var (
	commonLowerHeader = map[string]string***REMOVED******REMOVED*** // Go-Canonical-Case -> lower-case
	commonCanonHeader = map[string]string***REMOVED******REMOVED*** // lower-case -> Go-Canonical-Case
)

func init() ***REMOVED***
	for _, v := range []string***REMOVED***
		"accept",
		"accept-charset",
		"accept-encoding",
		"accept-language",
		"accept-ranges",
		"age",
		"access-control-allow-origin",
		"allow",
		"authorization",
		"cache-control",
		"content-disposition",
		"content-encoding",
		"content-language",
		"content-length",
		"content-location",
		"content-range",
		"content-type",
		"cookie",
		"date",
		"etag",
		"expect",
		"expires",
		"from",
		"host",
		"if-match",
		"if-modified-since",
		"if-none-match",
		"if-unmodified-since",
		"last-modified",
		"link",
		"location",
		"max-forwards",
		"proxy-authenticate",
		"proxy-authorization",
		"range",
		"referer",
		"refresh",
		"retry-after",
		"server",
		"set-cookie",
		"strict-transport-security",
		"trailer",
		"transfer-encoding",
		"user-agent",
		"vary",
		"via",
		"www-authenticate",
	***REMOVED*** ***REMOVED***
		chk := http.CanonicalHeaderKey(v)
		commonLowerHeader[chk] = v
		commonCanonHeader[v] = chk
	***REMOVED***
***REMOVED***

func lowerHeader(v string) string ***REMOVED***
	if s, ok := commonLowerHeader[v]; ok ***REMOVED***
		return s
	***REMOVED***
	return strings.ToLower(v)
***REMOVED***
