package msgp

import (
	"math"
	"reflect"
	"time"
)

// ensure 'sz' extra bytes in 'b' btw len(b) and cap(b)
func ensure(b []byte, sz int) ([]byte, int) ***REMOVED***
	l := len(b)
	c := cap(b)
	if c-l < sz ***REMOVED***
		o := make([]byte, (2*c)+sz) // exponential growth
		n := copy(o, b)
		return o[:n+sz], n
	***REMOVED***
	return b[:l+sz], l
***REMOVED***

// AppendMapHeader appends a map header with the
// given size to the slice
func AppendMapHeader(b []byte, sz uint32) []byte ***REMOVED***
	switch ***REMOVED***
	case sz <= 15:
		return append(b, wfixmap(uint8(sz)))

	case sz <= math.MaxUint16:
		o, n := ensure(b, 3)
		prefixu16(o[n:], mmap16, uint16(sz))
		return o

	default:
		o, n := ensure(b, 5)
		prefixu32(o[n:], mmap32, sz)
		return o
	***REMOVED***
***REMOVED***

// AppendArrayHeader appends an array header with
// the given size to the slice
func AppendArrayHeader(b []byte, sz uint32) []byte ***REMOVED***
	switch ***REMOVED***
	case sz <= 15:
		return append(b, wfixarray(uint8(sz)))

	case sz <= math.MaxUint16:
		o, n := ensure(b, 3)
		prefixu16(o[n:], marray16, uint16(sz))
		return o

	default:
		o, n := ensure(b, 5)
		prefixu32(o[n:], marray32, sz)
		return o
	***REMOVED***
***REMOVED***

// AppendNil appends a 'nil' byte to the slice
func AppendNil(b []byte) []byte ***REMOVED*** return append(b, mnil) ***REMOVED***

// AppendFloat64 appends a float64 to the slice
func AppendFloat64(b []byte, f float64) []byte ***REMOVED***
	o, n := ensure(b, Float64Size)
	prefixu64(o[n:], mfloat64, math.Float64bits(f))
	return o
***REMOVED***

// AppendFloat32 appends a float32 to the slice
func AppendFloat32(b []byte, f float32) []byte ***REMOVED***
	o, n := ensure(b, Float32Size)
	prefixu32(o[n:], mfloat32, math.Float32bits(f))
	return o
***REMOVED***

// AppendInt64 appends an int64 to the slice
func AppendInt64(b []byte, i int64) []byte ***REMOVED***
	if i >= 0 ***REMOVED***
		switch ***REMOVED***
		case i <= math.MaxInt8:
			return append(b, wfixint(uint8(i)))
		case i <= math.MaxInt16:
			o, n := ensure(b, 3)
			putMint16(o[n:], int16(i))
			return o
		case i <= math.MaxInt32:
			o, n := ensure(b, 5)
			putMint32(o[n:], int32(i))
			return o
		default:
			o, n := ensure(b, 9)
			putMint64(o[n:], i)
			return o
		***REMOVED***
	***REMOVED***
	switch ***REMOVED***
	case i >= -32:
		return append(b, wnfixint(int8(i)))
	case i >= math.MinInt8:
		o, n := ensure(b, 2)
		putMint8(o[n:], int8(i))
		return o
	case i >= math.MinInt16:
		o, n := ensure(b, 3)
		putMint16(o[n:], int16(i))
		return o
	case i >= math.MinInt32:
		o, n := ensure(b, 5)
		putMint32(o[n:], int32(i))
		return o
	default:
		o, n := ensure(b, 9)
		putMint64(o[n:], i)
		return o
	***REMOVED***
***REMOVED***

// AppendInt appends an int to the slice
func AppendInt(b []byte, i int) []byte ***REMOVED*** return AppendInt64(b, int64(i)) ***REMOVED***

// AppendInt8 appends an int8 to the slice
func AppendInt8(b []byte, i int8) []byte ***REMOVED*** return AppendInt64(b, int64(i)) ***REMOVED***

// AppendInt16 appends an int16 to the slice
func AppendInt16(b []byte, i int16) []byte ***REMOVED*** return AppendInt64(b, int64(i)) ***REMOVED***

// AppendInt32 appends an int32 to the slice
func AppendInt32(b []byte, i int32) []byte ***REMOVED*** return AppendInt64(b, int64(i)) ***REMOVED***

// AppendUint64 appends a uint64 to the slice
func AppendUint64(b []byte, u uint64) []byte ***REMOVED***
	switch ***REMOVED***
	case u <= (1<<7)-1:
		return append(b, wfixint(uint8(u)))

	case u <= math.MaxUint8:
		o, n := ensure(b, 2)
		putMuint8(o[n:], uint8(u))
		return o

	case u <= math.MaxUint16:
		o, n := ensure(b, 3)
		putMuint16(o[n:], uint16(u))
		return o

	case u <= math.MaxUint32:
		o, n := ensure(b, 5)
		putMuint32(o[n:], uint32(u))
		return o

	default:
		o, n := ensure(b, 9)
		putMuint64(o[n:], u)
		return o

	***REMOVED***
***REMOVED***

// AppendUint appends a uint to the slice
func AppendUint(b []byte, u uint) []byte ***REMOVED*** return AppendUint64(b, uint64(u)) ***REMOVED***

// AppendUint8 appends a uint8 to the slice
func AppendUint8(b []byte, u uint8) []byte ***REMOVED*** return AppendUint64(b, uint64(u)) ***REMOVED***

// AppendByte is analogous to AppendUint8
func AppendByte(b []byte, u byte) []byte ***REMOVED*** return AppendUint8(b, uint8(u)) ***REMOVED***

// AppendUint16 appends a uint16 to the slice
func AppendUint16(b []byte, u uint16) []byte ***REMOVED*** return AppendUint64(b, uint64(u)) ***REMOVED***

// AppendUint32 appends a uint32 to the slice
func AppendUint32(b []byte, u uint32) []byte ***REMOVED*** return AppendUint64(b, uint64(u)) ***REMOVED***

// AppendBytes appends bytes to the slice as MessagePack 'bin' data
func AppendBytes(b []byte, bts []byte) []byte ***REMOVED***
	sz := len(bts)
	var o []byte
	var n int
	switch ***REMOVED***
	case sz <= math.MaxUint8:
		o, n = ensure(b, 2+sz)
		prefixu8(o[n:], mbin8, uint8(sz))
		n += 2
	case sz <= math.MaxUint16:
		o, n = ensure(b, 3+sz)
		prefixu16(o[n:], mbin16, uint16(sz))
		n += 3
	default:
		o, n = ensure(b, 5+sz)
		prefixu32(o[n:], mbin32, uint32(sz))
		n += 5
	***REMOVED***
	return o[:n+copy(o[n:], bts)]
***REMOVED***

// AppendBool appends a bool to the slice
func AppendBool(b []byte, t bool) []byte ***REMOVED***
	if t ***REMOVED***
		return append(b, mtrue)
	***REMOVED***
	return append(b, mfalse)
***REMOVED***

// AppendString appends a string as a MessagePack 'str' to the slice
func AppendString(b []byte, s string) []byte ***REMOVED***
	sz := len(s)
	var n int
	var o []byte
	switch ***REMOVED***
	case sz <= 31:
		o, n = ensure(b, 1+sz)
		o[n] = wfixstr(uint8(sz))
		n++
	case sz <= math.MaxUint8:
		o, n = ensure(b, 2+sz)
		prefixu8(o[n:], mstr8, uint8(sz))
		n += 2
	case sz <= math.MaxUint16:
		o, n = ensure(b, 3+sz)
		prefixu16(o[n:], mstr16, uint16(sz))
		n += 3
	default:
		o, n = ensure(b, 5+sz)
		prefixu32(o[n:], mstr32, uint32(sz))
		n += 5
	***REMOVED***
	return o[:n+copy(o[n:], s)]
***REMOVED***

// AppendStringFromBytes appends a []byte
// as a MessagePack 'str' to the slice 'b.'
func AppendStringFromBytes(b []byte, str []byte) []byte ***REMOVED***
	sz := len(str)
	var n int
	var o []byte
	switch ***REMOVED***
	case sz <= 31:
		o, n = ensure(b, 1+sz)
		o[n] = wfixstr(uint8(sz))
		n++
	case sz <= math.MaxUint8:
		o, n = ensure(b, 2+sz)
		prefixu8(o[n:], mstr8, uint8(sz))
		n += 2
	case sz <= math.MaxUint16:
		o, n = ensure(b, 3+sz)
		prefixu16(o[n:], mstr16, uint16(sz))
		n += 3
	default:
		o, n = ensure(b, 5+sz)
		prefixu32(o[n:], mstr32, uint32(sz))
		n += 5
	***REMOVED***
	return o[:n+copy(o[n:], str)]
***REMOVED***

// AppendComplex64 appends a complex64 to the slice as a MessagePack extension
func AppendComplex64(b []byte, c complex64) []byte ***REMOVED***
	o, n := ensure(b, Complex64Size)
	o[n] = mfixext8
	o[n+1] = Complex64Extension
	big.PutUint32(o[n+2:], math.Float32bits(real(c)))
	big.PutUint32(o[n+6:], math.Float32bits(imag(c)))
	return o
***REMOVED***

// AppendComplex128 appends a complex128 to the slice as a MessagePack extension
func AppendComplex128(b []byte, c complex128) []byte ***REMOVED***
	o, n := ensure(b, Complex128Size)
	o[n] = mfixext16
	o[n+1] = Complex128Extension
	big.PutUint64(o[n+2:], math.Float64bits(real(c)))
	big.PutUint64(o[n+10:], math.Float64bits(imag(c)))
	return o
***REMOVED***

// AppendTime appends a time.Time to the slice as a MessagePack extension
func AppendTime(b []byte, t time.Time) []byte ***REMOVED***
	o, n := ensure(b, TimeSize)
	t = t.UTC()
	o[n] = mext8
	o[n+1] = 12
	o[n+2] = TimeExtension
	putUnix(o[n+3:], t.Unix(), int32(t.Nanosecond()))
	return o
***REMOVED***

// AppendMapStrStr appends a map[string]string to the slice
// as a MessagePack map with 'str'-type keys and values
func AppendMapStrStr(b []byte, m map[string]string) []byte ***REMOVED***
	sz := uint32(len(m))
	b = AppendMapHeader(b, sz)
	for key, val := range m ***REMOVED***
		b = AppendString(b, key)
		b = AppendString(b, val)
	***REMOVED***
	return b
***REMOVED***

// AppendMapStrIntf appends a map[string]interface***REMOVED******REMOVED*** to the slice
// as a MessagePack map with 'str'-type keys.
func AppendMapStrIntf(b []byte, m map[string]interface***REMOVED******REMOVED***) ([]byte, error) ***REMOVED***
	sz := uint32(len(m))
	b = AppendMapHeader(b, sz)
	var err error
	for key, val := range m ***REMOVED***
		b = AppendString(b, key)
		b, err = AppendIntf(b, val)
		if err != nil ***REMOVED***
			return b, err
		***REMOVED***
	***REMOVED***
	return b, nil
***REMOVED***

// AppendIntf appends the concrete type of 'i' to the
// provided []byte. 'i' must be one of the following:
//  - 'nil'
//  - A bool, float, string, []byte, int, uint, or complex
//  - A map[string]interface***REMOVED******REMOVED*** or map[string]string
//  - A []T, where T is another supported type
//  - A *T, where T is another supported type
//  - A type that satisfieds the msgp.Marshaler interface
//  - A type that satisfies the msgp.Extension interface
func AppendIntf(b []byte, i interface***REMOVED******REMOVED***) ([]byte, error) ***REMOVED***
	if i == nil ***REMOVED***
		return AppendNil(b), nil
	***REMOVED***

	// all the concrete types
	// for which we have methods
	switch i := i.(type) ***REMOVED***
	case Marshaler:
		return i.MarshalMsg(b)
	case Extension:
		return AppendExtension(b, i)
	case bool:
		return AppendBool(b, i), nil
	case float32:
		return AppendFloat32(b, i), nil
	case float64:
		return AppendFloat64(b, i), nil
	case complex64:
		return AppendComplex64(b, i), nil
	case complex128:
		return AppendComplex128(b, i), nil
	case string:
		return AppendString(b, i), nil
	case []byte:
		return AppendBytes(b, i), nil
	case int8:
		return AppendInt8(b, i), nil
	case int16:
		return AppendInt16(b, i), nil
	case int32:
		return AppendInt32(b, i), nil
	case int64:
		return AppendInt64(b, i), nil
	case int:
		return AppendInt64(b, int64(i)), nil
	case uint:
		return AppendUint64(b, uint64(i)), nil
	case uint8:
		return AppendUint8(b, i), nil
	case uint16:
		return AppendUint16(b, i), nil
	case uint32:
		return AppendUint32(b, i), nil
	case uint64:
		return AppendUint64(b, i), nil
	case time.Time:
		return AppendTime(b, i), nil
	case map[string]interface***REMOVED******REMOVED***:
		return AppendMapStrIntf(b, i)
	case map[string]string:
		return AppendMapStrStr(b, i), nil
	case []interface***REMOVED******REMOVED***:
		b = AppendArrayHeader(b, uint32(len(i)))
		var err error
		for _, k := range i ***REMOVED***
			b, err = AppendIntf(b, k)
			if err != nil ***REMOVED***
				return b, err
			***REMOVED***
		***REMOVED***
		return b, nil
	***REMOVED***

	var err error
	v := reflect.ValueOf(i)
	switch v.Kind() ***REMOVED***
	case reflect.Array, reflect.Slice:
		l := v.Len()
		b = AppendArrayHeader(b, uint32(l))
		for i := 0; i < l; i++ ***REMOVED***
			b, err = AppendIntf(b, v.Index(i).Interface())
			if err != nil ***REMOVED***
				return b, err
			***REMOVED***
		***REMOVED***
		return b, nil
	case reflect.Ptr:
		if v.IsNil() ***REMOVED***
			return AppendNil(b), err
		***REMOVED***
		b, err = AppendIntf(b, v.Elem().Interface())
		return b, err
	default:
		return b, &ErrUnsupportedType***REMOVED***T: v.Type()***REMOVED***
	***REMOVED***
***REMOVED***
