package ioutils

import (
	"errors"
	"io"
	"sync"
)

// maxCap is the highest capacity to use in byte slices that buffer data.
const maxCap = 1e6

// minCap is the lowest capacity to use in byte slices that buffer data
const minCap = 64

// blockThreshold is the minimum number of bytes in the buffer which will cause
// a write to BytesPipe to block when allocating a new slice.
const blockThreshold = 1e6

var (
	// ErrClosed is returned when Write is called on a closed BytesPipe.
	ErrClosed = errors.New("write to closed BytesPipe")

	bufPools     = make(map[int]*sync.Pool)
	bufPoolsLock sync.Mutex
)

// BytesPipe is io.ReadWriteCloser which works similarly to pipe(queue).
// All written data may be read at most once. Also, BytesPipe allocates
// and releases new byte slices to adjust to current needs, so the buffer
// won't be overgrown after peak loads.
type BytesPipe struct ***REMOVED***
	mu       sync.Mutex
	wait     *sync.Cond
	buf      []*fixedBuffer
	bufLen   int
	closeErr error // error to return from next Read. set to nil if not closed.
***REMOVED***

// NewBytesPipe creates new BytesPipe, initialized by specified slice.
// If buf is nil, then it will be initialized with slice which cap is 64.
// buf will be adjusted in a way that len(buf) == 0, cap(buf) == cap(buf).
func NewBytesPipe() *BytesPipe ***REMOVED***
	bp := &BytesPipe***REMOVED******REMOVED***
	bp.buf = append(bp.buf, getBuffer(minCap))
	bp.wait = sync.NewCond(&bp.mu)
	return bp
***REMOVED***

// Write writes p to BytesPipe.
// It can allocate new []byte slices in a process of writing.
func (bp *BytesPipe) Write(p []byte) (int, error) ***REMOVED***
	bp.mu.Lock()

	written := 0
loop0:
	for ***REMOVED***
		if bp.closeErr != nil ***REMOVED***
			bp.mu.Unlock()
			return written, ErrClosed
		***REMOVED***

		if len(bp.buf) == 0 ***REMOVED***
			bp.buf = append(bp.buf, getBuffer(64))
		***REMOVED***
		// get the last buffer
		b := bp.buf[len(bp.buf)-1]

		n, err := b.Write(p)
		written += n
		bp.bufLen += n

		// errBufferFull is an error we expect to get if the buffer is full
		if err != nil && err != errBufferFull ***REMOVED***
			bp.wait.Broadcast()
			bp.mu.Unlock()
			return written, err
		***REMOVED***

		// if there was enough room to write all then break
		if len(p) == n ***REMOVED***
			break
		***REMOVED***

		// more data: write to the next slice
		p = p[n:]

		// make sure the buffer doesn't grow too big from this write
		for bp.bufLen >= blockThreshold ***REMOVED***
			bp.wait.Wait()
			if bp.closeErr != nil ***REMOVED***
				continue loop0
			***REMOVED***
		***REMOVED***

		// add new byte slice to the buffers slice and continue writing
		nextCap := b.Cap() * 2
		if nextCap > maxCap ***REMOVED***
			nextCap = maxCap
		***REMOVED***
		bp.buf = append(bp.buf, getBuffer(nextCap))
	***REMOVED***
	bp.wait.Broadcast()
	bp.mu.Unlock()
	return written, nil
***REMOVED***

// CloseWithError causes further reads from a BytesPipe to return immediately.
func (bp *BytesPipe) CloseWithError(err error) error ***REMOVED***
	bp.mu.Lock()
	if err != nil ***REMOVED***
		bp.closeErr = err
	***REMOVED*** else ***REMOVED***
		bp.closeErr = io.EOF
	***REMOVED***
	bp.wait.Broadcast()
	bp.mu.Unlock()
	return nil
***REMOVED***

// Close causes further reads from a BytesPipe to return immediately.
func (bp *BytesPipe) Close() error ***REMOVED***
	return bp.CloseWithError(nil)
***REMOVED***

// Read reads bytes from BytesPipe.
// Data could be read only once.
func (bp *BytesPipe) Read(p []byte) (n int, err error) ***REMOVED***
	bp.mu.Lock()
	if bp.bufLen == 0 ***REMOVED***
		if bp.closeErr != nil ***REMOVED***
			bp.mu.Unlock()
			return 0, bp.closeErr
		***REMOVED***
		bp.wait.Wait()
		if bp.bufLen == 0 && bp.closeErr != nil ***REMOVED***
			err := bp.closeErr
			bp.mu.Unlock()
			return 0, err
		***REMOVED***
	***REMOVED***

	for bp.bufLen > 0 ***REMOVED***
		b := bp.buf[0]
		read, _ := b.Read(p) // ignore error since fixedBuffer doesn't really return an error
		n += read
		bp.bufLen -= read

		if b.Len() == 0 ***REMOVED***
			// it's empty so return it to the pool and move to the next one
			returnBuffer(b)
			bp.buf[0] = nil
			bp.buf = bp.buf[1:]
		***REMOVED***

		if len(p) == read ***REMOVED***
			break
		***REMOVED***

		p = p[read:]
	***REMOVED***

	bp.wait.Broadcast()
	bp.mu.Unlock()
	return
***REMOVED***

func returnBuffer(b *fixedBuffer) ***REMOVED***
	b.Reset()
	bufPoolsLock.Lock()
	pool := bufPools[b.Cap()]
	bufPoolsLock.Unlock()
	if pool != nil ***REMOVED***
		pool.Put(b)
	***REMOVED***
***REMOVED***

func getBuffer(size int) *fixedBuffer ***REMOVED***
	bufPoolsLock.Lock()
	pool, ok := bufPools[size]
	if !ok ***REMOVED***
		pool = &sync.Pool***REMOVED***New: func() interface***REMOVED******REMOVED*** ***REMOVED*** return &fixedBuffer***REMOVED***buf: make([]byte, 0, size)***REMOVED*** ***REMOVED******REMOVED***
		bufPools[size] = pool
	***REMOVED***
	bufPoolsLock.Unlock()
	return pool.Get().(*fixedBuffer)
***REMOVED***
