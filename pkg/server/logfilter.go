package server

import (
	"bytes"
	"io"
)

var (
	TLS_STRING = []byte("http: TLS handshake error from")
)

type HandshakeFilterWriter struct {
	io.Writer
}

func NewHandshakeFilterWriter(w io.Writer) io.Writer {
	return &HandshakeFilterWriter{
		Writer: w,
	}
}

func (w *HandshakeFilterWriter) Write(b []byte) (n int, err error) {
	if bytes.Contains(b, TLS_STRING) {
		return 0, nil
	}
	return w.Writer.Write(b)
}
