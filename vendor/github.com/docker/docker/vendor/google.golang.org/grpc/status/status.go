/*
 *
 * Copyright 2017, Google Inc.
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 *     * Redistributions of source code must retain the above copyright
 * notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above
 * copyright notice, this list of conditions and the following disclaimer
 * in the documentation and/or other materials provided with the
 * distribution.
 *     * Neither the name of Google Inc. nor the names of its
 * contributors may be used to endorse or promote products derived from
 * this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 * A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 * OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 * LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 * DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 * THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 */

// Package status implements errors returned by gRPC.  These errors are
// serialized and transmitted on the wire between server and client, and allow
// for additional data to be transmitted via the Details field in the status
// proto.  gRPC service handlers should return an error created by this
// package, and gRPC clients should expect a corresponding error to be
// returned from the RPC call.
//
// This package upholds the invariants that a non-nil error may not
// contain an OK code, and an OK code must result in a nil error.
package status

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
)

// statusError is an alias of a status proto.  It implements error and Status,
// and a nil statusError should never be returned by this package.
type statusError spb.Status

func (se *statusError) Error() string ***REMOVED***
	p := (*spb.Status)(se)
	return fmt.Sprintf("rpc error: code = %s desc = %s", codes.Code(p.GetCode()), p.GetMessage())
***REMOVED***

func (se *statusError) status() *Status ***REMOVED***
	return &Status***REMOVED***s: (*spb.Status)(se)***REMOVED***
***REMOVED***

// Status represents an RPC status code, message, and details.  It is immutable
// and should be created with New, Newf, or FromProto.
type Status struct ***REMOVED***
	s *spb.Status
***REMOVED***

// Code returns the status code contained in s.
func (s *Status) Code() codes.Code ***REMOVED***
	if s == nil || s.s == nil ***REMOVED***
		return codes.OK
	***REMOVED***
	return codes.Code(s.s.Code)
***REMOVED***

// Message returns the message contained in s.
func (s *Status) Message() string ***REMOVED***
	if s == nil || s.s == nil ***REMOVED***
		return ""
	***REMOVED***
	return s.s.Message
***REMOVED***

// Proto returns s's status as an spb.Status proto message.
func (s *Status) Proto() *spb.Status ***REMOVED***
	if s == nil ***REMOVED***
		return nil
	***REMOVED***
	return proto.Clone(s.s).(*spb.Status)
***REMOVED***

// Err returns an immutable error representing s; returns nil if s.Code() is
// OK.
func (s *Status) Err() error ***REMOVED***
	if s.Code() == codes.OK ***REMOVED***
		return nil
	***REMOVED***
	return (*statusError)(s.s)
***REMOVED***

// New returns a Status representing c and msg.
func New(c codes.Code, msg string) *Status ***REMOVED***
	return &Status***REMOVED***s: &spb.Status***REMOVED***Code: int32(c), Message: msg***REMOVED******REMOVED***
***REMOVED***

// Newf returns New(c, fmt.Sprintf(format, a...)).
func Newf(c codes.Code, format string, a ...interface***REMOVED******REMOVED***) *Status ***REMOVED***
	return New(c, fmt.Sprintf(format, a...))
***REMOVED***

// Error returns an error representing c and msg.  If c is OK, returns nil.
func Error(c codes.Code, msg string) error ***REMOVED***
	return New(c, msg).Err()
***REMOVED***

// Errorf returns Error(c, fmt.Sprintf(format, a...)).
func Errorf(c codes.Code, format string, a ...interface***REMOVED******REMOVED***) error ***REMOVED***
	return Error(c, fmt.Sprintf(format, a...))
***REMOVED***

// ErrorProto returns an error representing s.  If s.Code is OK, returns nil.
func ErrorProto(s *spb.Status) error ***REMOVED***
	return FromProto(s).Err()
***REMOVED***

// FromProto returns a Status representing s.
func FromProto(s *spb.Status) *Status ***REMOVED***
	return &Status***REMOVED***s: proto.Clone(s).(*spb.Status)***REMOVED***
***REMOVED***

// FromError returns a Status representing err if it was produced from this
// package, otherwise it returns nil, false.
func FromError(err error) (s *Status, ok bool) ***REMOVED***
	if err == nil ***REMOVED***
		return &Status***REMOVED***s: &spb.Status***REMOVED***Code: int32(codes.OK)***REMOVED******REMOVED***, true
	***REMOVED***
	if s, ok := err.(*statusError); ok ***REMOVED***
		return s.status(), true
	***REMOVED***
	return nil, false
***REMOVED***
