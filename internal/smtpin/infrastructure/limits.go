package infrastructure

import (
	"io"
	"net"
)

type LimitedConn struct {
	net.Conn
	Reader io.Reader
}

func (c LimitedConn) Read(p []byte) (int, error) { return c.Reader.Read(p) }
func LimitReader(r io.Reader, n int64) io.Reader { return io.LimitReader(r, n) }
