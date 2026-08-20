package grpcapi

import (
	"context"
	"log/slog"
	"net"
)

type Server struct {
	Addr string
	Log  *slog.Logger
}

func (s *Server) Listen(ctx context.Context) error {
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	defer ln.Close()
	go func() { <-ctx.Done(); _ = ln.Close() }()
	for {
		conn, err := ln.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				return nil
			default:
				return err
			}
		}
		go func(c net.Conn) {
			defer c.Close()
			_, _ = c.Write([]byte("maildelivery grpc endpoint; use REST for JSON\n"))
		}(conn)
	}
}
