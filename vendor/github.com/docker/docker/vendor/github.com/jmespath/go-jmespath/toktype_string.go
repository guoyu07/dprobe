// generated by stringer -type=tokType; DO NOT EDIT

package jmespath

import "fmt"

const _tokType_name = "tUnknowntStartDottFiltertFlattentLparentRparentLbrackettRbrackettLbracetRbracetOrtPipetNumbertUnquotedIdentifiertQuotedIdentifiertCommatColontLTtLTEtGTtGTEtEQtNEtJSONLiteraltStringLiteraltCurrenttExpreftAndtNottEOF"

var _tokType_index = [...]uint8***REMOVED***0, 8, 13, 17, 24, 32, 39, 46, 55, 64, 71, 78, 81, 86, 93, 112, 129, 135, 141, 144, 148, 151, 155, 158, 161, 173, 187, 195, 202, 206, 210, 214***REMOVED***

func (i tokType) String() string ***REMOVED***
	if i < 0 || i >= tokType(len(_tokType_index)-1) ***REMOVED***
		return fmt.Sprintf("tokType(%d)", i)
	***REMOVED***
	return _tokType_name[_tokType_index[i]:_tokType_index[i+1]]
***REMOVED***
