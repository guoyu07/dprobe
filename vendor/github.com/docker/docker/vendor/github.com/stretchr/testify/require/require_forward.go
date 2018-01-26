/*
* CODE GENERATED AUTOMATICALLY WITH github.com/stretchr/testify/_codegen
* THIS FILE MUST NOT BE EDITED BY HAND
 */

package require

import (
	assert "github.com/stretchr/testify/assert"
	http "net/http"
	url "net/url"
	time "time"
)

// Condition uses a Comparison to assert a complex condition.
func (a *Assertions) Condition(comp assert.Comparison, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	Condition(a.t, comp, msgAndArgs...)
***REMOVED***

// Contains asserts that the specified string, list(array, slice...) or map contains the
// specified substring or element.
//
//    a.Contains("Hello World", "World", "But 'Hello World' does contain 'World'")
//    a.Contains(["Hello", "World"], "World", "But ["Hello", "World"] does contain 'World'")
//    a.Contains(***REMOVED***"Hello": "World"***REMOVED***, "Hello", "But ***REMOVED***'Hello': 'World'***REMOVED*** does contain 'Hello'")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Contains(s interface***REMOVED******REMOVED***, contains interface***REMOVED******REMOVED***, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	Contains(a.t, s, contains, msgAndArgs...)
***REMOVED***

// Empty asserts that the specified object is empty.  I.e. nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  a.Empty(obj)
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Empty(object interface***REMOVED******REMOVED***, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	Empty(a.t, object, msgAndArgs...)
***REMOVED***

// Equal asserts that two objects are equal.
//
//    a.Equal(123, 123, "123 and 123 should be equal")
//
// Returns whether the assertion was successful (true) or not (false).
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
func (a *Assertions) Equal(expected interface***REMOVED******REMOVED***, actual interface***REMOVED******REMOVED***, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	Equal(a.t, expected, actual, msgAndArgs...)
***REMOVED***

// EqualError asserts that a function returned an error (i.e. not `nil`)
// and that it is equal to the provided error.
//
//   actualObj, err := SomeFunction()
//   a.EqualError(err,  expectedErrorString, "An error was expected")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) EqualError(theError error, errString string, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	EqualError(a.t, theError, errString, msgAndArgs...)
***REMOVED***

// EqualValues asserts that two objects are equal or convertable to the same types
// and equal.
//
//    a.EqualValues(uint32(123), int32(123), "123 and 123 should be equal")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) EqualValues(expected interface***REMOVED******REMOVED***, actual interface***REMOVED******REMOVED***, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	EqualValues(a.t, expected, actual, msgAndArgs...)
***REMOVED***

// Error asserts that a function returned an error (i.e. not `nil`).
//
//   actualObj, err := SomeFunction()
//   if a.Error(err, "An error was expected") ***REMOVED***
// 	   assert.Equal(t, err, expectedError)
//   ***REMOVED***
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Error(err error, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	Error(a.t, err, msgAndArgs...)
***REMOVED***

// Exactly asserts that two objects are equal is value and type.
//
//    a.Exactly(int32(123), int64(123), "123 and 123 should NOT be equal")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Exactly(expected interface***REMOVED******REMOVED***, actual interface***REMOVED******REMOVED***, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	Exactly(a.t, expected, actual, msgAndArgs...)
***REMOVED***

// Fail reports a failure through
func (a *Assertions) Fail(failureMessage string, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	Fail(a.t, failureMessage, msgAndArgs...)
***REMOVED***

// FailNow fails test
func (a *Assertions) FailNow(failureMessage string, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	FailNow(a.t, failureMessage, msgAndArgs...)
***REMOVED***

// False asserts that the specified value is false.
//
//    a.False(myBool, "myBool should be false")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) False(value bool, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	False(a.t, value, msgAndArgs...)
***REMOVED***

// HTTPBodyContains asserts that a specified handler returns a
// body that contains a string.
//
//  a.HTTPBodyContains(myHandler, "www.google.com", nil, "I'm Feeling Lucky")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) HTTPBodyContains(handler http.HandlerFunc, method string, url string, values url.Values, str interface***REMOVED******REMOVED***) ***REMOVED***
	HTTPBodyContains(a.t, handler, method, url, values, str)
***REMOVED***

// HTTPBodyNotContains asserts that a specified handler returns a
// body that does not contain a string.
//
//  a.HTTPBodyNotContains(myHandler, "www.google.com", nil, "I'm Feeling Lucky")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) HTTPBodyNotContains(handler http.HandlerFunc, method string, url string, values url.Values, str interface***REMOVED******REMOVED***) ***REMOVED***
	HTTPBodyNotContains(a.t, handler, method, url, values, str)
***REMOVED***

// HTTPError asserts that a specified handler returns an error status code.
//
//  a.HTTPError(myHandler, "POST", "/a/b/c", url.Values***REMOVED***"a": []string***REMOVED***"b", "c"***REMOVED******REMOVED***
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) HTTPError(handler http.HandlerFunc, method string, url string, values url.Values) ***REMOVED***
	HTTPError(a.t, handler, method, url, values)
***REMOVED***

// HTTPRedirect asserts that a specified handler returns a redirect status code.
//
//  a.HTTPRedirect(myHandler, "GET", "/a/b/c", url.Values***REMOVED***"a": []string***REMOVED***"b", "c"***REMOVED******REMOVED***
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) HTTPRedirect(handler http.HandlerFunc, method string, url string, values url.Values) ***REMOVED***
	HTTPRedirect(a.t, handler, method, url, values)
***REMOVED***

// HTTPSuccess asserts that a specified handler returns a success status code.
//
//  a.HTTPSuccess(myHandler, "POST", "http://www.google.com", nil)
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) HTTPSuccess(handler http.HandlerFunc, method string, url string, values url.Values) ***REMOVED***
	HTTPSuccess(a.t, handler, method, url, values)
***REMOVED***

// Implements asserts that an object is implemented by the specified interface.
//
//    a.Implements((*MyInterface)(nil), new(MyObject), "MyObject")
func (a *Assertions) Implements(interfaceObject interface***REMOVED******REMOVED***, object interface***REMOVED******REMOVED***, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	Implements(a.t, interfaceObject, object, msgAndArgs...)
***REMOVED***

// InDelta asserts that the two numerals are within delta of each other.
//
// 	 a.InDelta(math.Pi, (22 / 7.0), 0.01)
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) InDelta(expected interface***REMOVED******REMOVED***, actual interface***REMOVED******REMOVED***, delta float64, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	InDelta(a.t, expected, actual, delta, msgAndArgs...)
***REMOVED***

// InDeltaSlice is the same as InDelta, except it compares two slices.
func (a *Assertions) InDeltaSlice(expected interface***REMOVED******REMOVED***, actual interface***REMOVED******REMOVED***, delta float64, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	InDeltaSlice(a.t, expected, actual, delta, msgAndArgs...)
***REMOVED***

// InEpsilon asserts that expected and actual have a relative error less than epsilon
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) InEpsilon(expected interface***REMOVED******REMOVED***, actual interface***REMOVED******REMOVED***, epsilon float64, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	InEpsilon(a.t, expected, actual, epsilon, msgAndArgs...)
***REMOVED***

// InEpsilonSlice is the same as InEpsilon, except it compares each value from two slices.
func (a *Assertions) InEpsilonSlice(expected interface***REMOVED******REMOVED***, actual interface***REMOVED******REMOVED***, epsilon float64, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	InEpsilonSlice(a.t, expected, actual, epsilon, msgAndArgs...)
***REMOVED***

// IsType asserts that the specified objects are of the same type.
func (a *Assertions) IsType(expectedType interface***REMOVED******REMOVED***, object interface***REMOVED******REMOVED***, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	IsType(a.t, expectedType, object, msgAndArgs...)
***REMOVED***

// JSONEq asserts that two JSON strings are equivalent.
//
//  a.JSONEq(`***REMOVED***"hello": "world", "foo": "bar"***REMOVED***`, `***REMOVED***"foo": "bar", "hello": "world"***REMOVED***`)
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) JSONEq(expected string, actual string, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	JSONEq(a.t, expected, actual, msgAndArgs...)
***REMOVED***

// Len asserts that the specified object has specific length.
// Len also fails if the object has a type that len() not accept.
//
//    a.Len(mySlice, 3, "The size of slice is not 3")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Len(object interface***REMOVED******REMOVED***, length int, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	Len(a.t, object, length, msgAndArgs...)
***REMOVED***

// Nil asserts that the specified object is nil.
//
//    a.Nil(err, "err should be nothing")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Nil(object interface***REMOVED******REMOVED***, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	Nil(a.t, object, msgAndArgs...)
***REMOVED***

// NoError asserts that a function returned no error (i.e. `nil`).
//
//   actualObj, err := SomeFunction()
//   if a.NoError(err) ***REMOVED***
// 	   assert.Equal(t, actualObj, expectedObj)
//   ***REMOVED***
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NoError(err error, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	NoError(a.t, err, msgAndArgs...)
***REMOVED***

// NotContains asserts that the specified string, list(array, slice...) or map does NOT contain the
// specified substring or element.
//
//    a.NotContains("Hello World", "Earth", "But 'Hello World' does NOT contain 'Earth'")
//    a.NotContains(["Hello", "World"], "Earth", "But ['Hello', 'World'] does NOT contain 'Earth'")
//    a.NotContains(***REMOVED***"Hello": "World"***REMOVED***, "Earth", "But ***REMOVED***'Hello': 'World'***REMOVED*** does NOT contain 'Earth'")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotContains(s interface***REMOVED******REMOVED***, contains interface***REMOVED******REMOVED***, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	NotContains(a.t, s, contains, msgAndArgs...)
***REMOVED***

// NotEmpty asserts that the specified object is NOT empty.  I.e. not nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  if a.NotEmpty(obj) ***REMOVED***
//    assert.Equal(t, "two", obj[1])
//  ***REMOVED***
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotEmpty(object interface***REMOVED******REMOVED***, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	NotEmpty(a.t, object, msgAndArgs...)
***REMOVED***

// NotEqual asserts that the specified values are NOT equal.
//
//    a.NotEqual(obj1, obj2, "two objects shouldn't be equal")
//
// Returns whether the assertion was successful (true) or not (false).
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
func (a *Assertions) NotEqual(expected interface***REMOVED******REMOVED***, actual interface***REMOVED******REMOVED***, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	NotEqual(a.t, expected, actual, msgAndArgs...)
***REMOVED***

// NotNil asserts that the specified object is not nil.
//
//    a.NotNil(err, "err should be something")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotNil(object interface***REMOVED******REMOVED***, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	NotNil(a.t, object, msgAndArgs...)
***REMOVED***

// NotPanics asserts that the code inside the specified PanicTestFunc does NOT panic.
//
//   a.NotPanics(func()***REMOVED***
//     RemainCalm()
//   ***REMOVED***, "Calling RemainCalm() should NOT panic")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotPanics(f assert.PanicTestFunc, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	NotPanics(a.t, f, msgAndArgs...)
***REMOVED***

// NotRegexp asserts that a specified regexp does not match a string.
//
//  a.NotRegexp(regexp.MustCompile("starts"), "it's starting")
//  a.NotRegexp("^start", "it's not starting")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotRegexp(rx interface***REMOVED******REMOVED***, str interface***REMOVED******REMOVED***, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	NotRegexp(a.t, rx, str, msgAndArgs...)
***REMOVED***

// NotZero asserts that i is not the zero value for its type and returns the truth.
func (a *Assertions) NotZero(i interface***REMOVED******REMOVED***, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	NotZero(a.t, i, msgAndArgs...)
***REMOVED***

// Panics asserts that the code inside the specified PanicTestFunc panics.
//
//   a.Panics(func()***REMOVED***
//     GoCrazy()
//   ***REMOVED***, "Calling GoCrazy() should panic")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Panics(f assert.PanicTestFunc, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	Panics(a.t, f, msgAndArgs...)
***REMOVED***

// Regexp asserts that a specified regexp matches a string.
//
//  a.Regexp(regexp.MustCompile("start"), "it's starting")
//  a.Regexp("start...$", "it's not starting")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Regexp(rx interface***REMOVED******REMOVED***, str interface***REMOVED******REMOVED***, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	Regexp(a.t, rx, str, msgAndArgs...)
***REMOVED***

// True asserts that the specified value is true.
//
//    a.True(myBool, "myBool should be true")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) True(value bool, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	True(a.t, value, msgAndArgs...)
***REMOVED***

// WithinDuration asserts that the two times are within duration delta of each other.
//
//   a.WithinDuration(time.Now(), time.Now(), 10*time.Second, "The difference should not be more than 10s")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) WithinDuration(expected time.Time, actual time.Time, delta time.Duration, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	WithinDuration(a.t, expected, actual, delta, msgAndArgs...)
***REMOVED***

// Zero asserts that i is the zero value for its type and returns the truth.
func (a *Assertions) Zero(i interface***REMOVED******REMOVED***, msgAndArgs ...interface***REMOVED******REMOVED***) ***REMOVED***
	Zero(a.t, i, msgAndArgs...)
***REMOVED***
