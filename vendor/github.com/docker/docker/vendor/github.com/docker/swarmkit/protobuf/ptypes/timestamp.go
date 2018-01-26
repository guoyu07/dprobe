package ptypes

import (
	"time"

	gogotypes "github.com/gogo/protobuf/types"
)

// MustTimestampProto converts time.Time to a google.protobuf.Timestamp proto.
// It panics if input timestamp is invalid.
func MustTimestampProto(t time.Time) *gogotypes.Timestamp ***REMOVED***
	ts, err := gogotypes.TimestampProto(t)
	if err != nil ***REMOVED***
		panic(err.Error())
	***REMOVED***
	return ts
***REMOVED***
