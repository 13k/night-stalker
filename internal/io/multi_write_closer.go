package io

import (
	"errors"
	"io"
)

type multiWriteCloser struct {
	io.Writer
	closers []io.Closer
}

var _ io.StringWriter = (*multiWriteCloser)(nil)
var errNotStringWriter = errors.New("multiWriteCloser initialized with a non-StringWriter writer")

func (t *multiWriteCloser) Close() (err error) {
	for _, w := range t.closers {
		err = w.Close()

		if err != nil {
			return
		}
	}

	return nil
}

func (t *multiWriteCloser) WriteString(s string) (n int, err error) {
	if sw, ok := t.Writer.(io.StringWriter); ok {
		return sw.WriteString(s)
	}

	panic(errNotStringWriter)
}

func MultiWriteCloser(streams ...io.WriteCloser) io.WriteCloser {
	writers := make([]io.Writer, 0, len(streams))
	closers := make([]io.Closer, 0, len(streams))

	for _, w := range streams {
		if mwc, ok := w.(*multiWriteCloser); ok {
			closers = append(closers, mwc.closers...)
			writers = append(writers, mwc.Writer)
		} else {
			closers = append(closers, w)
			writers = append(writers, w)
		}
	}

	return &multiWriteCloser{
		Writer:  io.MultiWriter(writers...),
		closers: closers,
	}
}
