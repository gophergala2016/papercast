package main

import (
	"io"
	"net/http"
	"strconv"
)

// ResponseBuffer provides an in-memory content buffer for HTTP requests while also implementing the
// ReaderAt and WriterAt interfaces in addition to the http.ResponseWriter interface.
// ReaderSeeker is implemented to allow byte serving.
type ResponseBuffer struct {
	buf     []byte
	readOfs int
}

func (rb *ResponseBuffer) Read(data []byte) (n int, err error) {
	n = copy(data, rb.buf[rb.readOfs:])
	rb.readOfs += n
	if rb.readOfs >= len(rb.buf) {
		err = io.EOF
	}
	return
}

// grow grows the buffer to guarantee space for n more bytes, increasing the length to accomdate them
func (rb *ResponseBuffer) grow(n int) int {
	m := len(rb.buf)
	if m+n > cap(rb.buf) {
		var buf []byte
		buf = make([]byte, 2*cap(rb.buf)+n)
		copy(buf, rb.buf)
		rb.buf = buf
	}
	rb.buf = rb.buf[0 : m+n]
	return m
}

// WriteTo writes the buffered contents and all http header information to another http.ResponseWriter.
func (rb *ResponseBuffer) WriteTo(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "audio/mp4")
	w.Header().Set("Content-Length", strconv.Itoa(len(rb.buf)))
	if len(rb.buf) > 0 {
		if _, err := w.Write(rb.buf); err != nil {
			return err
		}
	}
	return nil
}

func (rb *ResponseBuffer) Write(p []byte) (int, error) {
	m := rb.grow(len(p))
	return copy(rb.buf[m:], p), nil
}

func (rb *ResponseBuffer) WriteAt(p []byte, off int64) (n int, err error) {
	need := len(p) + int(off) - len(rb.buf)
	if need > 0 {
		rb.grow(need)
	}
	return copy(rb.buf[off:], p), nil
}

func (rb *ResponseBuffer) ReadAt(p []byte, off int64) (n int, err error) {
	if int(off) >= len(rb.buf) {
		if len(p) == 0 {
			return
		}
		return 0, io.EOF
	}
	n = copy(p, rb.buf[off:])
	return
}
