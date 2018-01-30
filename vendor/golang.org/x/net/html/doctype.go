// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

import (
	"strings"
)

// parseDoctype parses the data from a DoctypeToken into a name,
// public identifier, and system identifier. It returns a Node whose Type
// is DoctypeNode, whose Data is the name, and which has attributes
// named "system" and "public" for the two identifiers if they were present.
// quirks is whether the document should be parsed in "quirks mode".
func parseDoctype(s string) (n *Node, quirks bool) ***REMOVED***
	n = &Node***REMOVED***Type: DoctypeNode***REMOVED***

	// Find the name.
	space := strings.IndexAny(s, whitespace)
	if space == -1 ***REMOVED***
		space = len(s)
	***REMOVED***
	n.Data = s[:space]
	// The comparison to "html" is case-sensitive.
	if n.Data != "html" ***REMOVED***
		quirks = true
	***REMOVED***
	n.Data = strings.ToLower(n.Data)
	s = strings.TrimLeft(s[space:], whitespace)

	if len(s) < 6 ***REMOVED***
		// It can't start with "PUBLIC" or "SYSTEM".
		// Ignore the rest of the string.
		return n, quirks || s != ""
	***REMOVED***

	key := strings.ToLower(s[:6])
	s = s[6:]
	for key == "public" || key == "system" ***REMOVED***
		s = strings.TrimLeft(s, whitespace)
		if s == "" ***REMOVED***
			break
		***REMOVED***
		quote := s[0]
		if quote != '"' && quote != '\'' ***REMOVED***
			break
		***REMOVED***
		s = s[1:]
		q := strings.IndexRune(s, rune(quote))
		var id string
		if q == -1 ***REMOVED***
			id = s
			s = ""
		***REMOVED*** else ***REMOVED***
			id = s[:q]
			s = s[q+1:]
		***REMOVED***
		n.Attr = append(n.Attr, Attribute***REMOVED***Key: key, Val: id***REMOVED***)
		if key == "public" ***REMOVED***
			key = "system"
		***REMOVED*** else ***REMOVED***
			key = ""
		***REMOVED***
	***REMOVED***

	if key != "" || s != "" ***REMOVED***
		quirks = true
	***REMOVED*** else if len(n.Attr) > 0 ***REMOVED***
		if n.Attr[0].Key == "public" ***REMOVED***
			public := strings.ToLower(n.Attr[0].Val)
			switch public ***REMOVED***
			case "-//w3o//dtd w3 html strict 3.0//en//", "-/w3d/dtd html 4.0 transitional/en", "html":
				quirks = true
			default:
				for _, q := range quirkyIDs ***REMOVED***
					if strings.HasPrefix(public, q) ***REMOVED***
						quirks = true
						break
					***REMOVED***
				***REMOVED***
			***REMOVED***
			// The following two public IDs only cause quirks mode if there is no system ID.
			if len(n.Attr) == 1 && (strings.HasPrefix(public, "-//w3c//dtd html 4.01 frameset//") ||
				strings.HasPrefix(public, "-//w3c//dtd html 4.01 transitional//")) ***REMOVED***
				quirks = true
			***REMOVED***
		***REMOVED***
		if lastAttr := n.Attr[len(n.Attr)-1]; lastAttr.Key == "system" &&
			strings.ToLower(lastAttr.Val) == "http://www.ibm.com/data/dtd/v11/ibmxhtml1-transitional.dtd" ***REMOVED***
			quirks = true
		***REMOVED***
	***REMOVED***

	return n, quirks
***REMOVED***

// quirkyIDs is a list of public doctype identifiers that cause a document
// to be interpreted in quirks mode. The identifiers should be in lower case.
var quirkyIDs = []string***REMOVED***
	"+//silmaril//dtd html pro v0r11 19970101//",
	"-//advasoft ltd//dtd html 3.0 aswedit + extensions//",
	"-//as//dtd html 3.0 aswedit + extensions//",
	"-//ietf//dtd html 2.0 level 1//",
	"-//ietf//dtd html 2.0 level 2//",
	"-//ietf//dtd html 2.0 strict level 1//",
	"-//ietf//dtd html 2.0 strict level 2//",
	"-//ietf//dtd html 2.0 strict//",
	"-//ietf//dtd html 2.0//",
	"-//ietf//dtd html 2.1e//",
	"-//ietf//dtd html 3.0//",
	"-//ietf//dtd html 3.2 final//",
	"-//ietf//dtd html 3.2//",
	"-//ietf//dtd html 3//",
	"-//ietf//dtd html level 0//",
	"-//ietf//dtd html level 1//",
	"-//ietf//dtd html level 2//",
	"-//ietf//dtd html level 3//",
	"-//ietf//dtd html strict level 0//",
	"-//ietf//dtd html strict level 1//",
	"-//ietf//dtd html strict level 2//",
	"-//ietf//dtd html strict level 3//",
	"-//ietf//dtd html strict//",
	"-//ietf//dtd html//",
	"-//metrius//dtd metrius presentational//",
	"-//microsoft//dtd internet explorer 2.0 html strict//",
	"-//microsoft//dtd internet explorer 2.0 html//",
	"-//microsoft//dtd internet explorer 2.0 tables//",
	"-//microsoft//dtd internet explorer 3.0 html strict//",
	"-//microsoft//dtd internet explorer 3.0 html//",
	"-//microsoft//dtd internet explorer 3.0 tables//",
	"-//netscape comm. corp.//dtd html//",
	"-//netscape comm. corp.//dtd strict html//",
	"-//o'reilly and associates//dtd html 2.0//",
	"-//o'reilly and associates//dtd html extended 1.0//",
	"-//o'reilly and associates//dtd html extended relaxed 1.0//",
	"-//softquad software//dtd hotmetal pro 6.0::19990601::extensions to html 4.0//",
	"-//softquad//dtd hotmetal pro 4.0::19971010::extensions to html 4.0//",
	"-//spyglass//dtd html 2.0 extended//",
	"-//sq//dtd html 2.0 hotmetal + extensions//",
	"-//sun microsystems corp.//dtd hotjava html//",
	"-//sun microsystems corp.//dtd hotjava strict html//",
	"-//w3c//dtd html 3 1995-03-24//",
	"-//w3c//dtd html 3.2 draft//",
	"-//w3c//dtd html 3.2 final//",
	"-//w3c//dtd html 3.2//",
	"-//w3c//dtd html 3.2s draft//",
	"-//w3c//dtd html 4.0 frameset//",
	"-//w3c//dtd html 4.0 transitional//",
	"-//w3c//dtd html experimental 19960712//",
	"-//w3c//dtd html experimental 970421//",
	"-//w3c//dtd w3 html//",
	"-//w3o//dtd w3 html 3.0//",
	"-//webtechs//dtd mozilla html 2.0//",
	"-//webtechs//dtd mozilla html//",
***REMOVED***