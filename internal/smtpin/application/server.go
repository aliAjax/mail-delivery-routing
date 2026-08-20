package application

import (
	"bufio"
	"context"
	"example.com/maildelivery/internal/message/application"
	"example.com/maildelivery/internal/message/domain"
	"fmt"
	"io"
	"net"
	"strings"
)

type Server struct {
	Addr     string
	Messages *application.Service
}

func (s *Server) Run(ctx context.Context) error {
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	defer ln.Close()
	go func() { <-ctx.Done(); _ = ln.Close() }()
	for {
		c, err := ln.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				return nil
			default:
				return err
			}
		}
		go s.session(ctx, c)
	}
}
func (s *Server) session(ctx context.Context, c net.Conn) {
	defer c.Close()
	rw := bufio.NewReadWriter(bufio.NewReader(io.LimitReader(c, 1<<20)), bufio.NewWriter(c))
	fmt.Fprintln(rw, "220 maildelivery ESMTP")
	_ = rw.Flush()
	var from, to string
	for {
		line, err := rw.ReadString('\n')
		if err != nil {
			return
		}
		cmd := strings.TrimSpace(line)
		upper := strings.ToUpper(cmd)
		switch {
		case strings.HasPrefix(upper, "EHLO"), strings.HasPrefix(upper, "HELO"):
			fmt.Fprintln(rw, "250-maildelivery\n250-SIZE 1048576\n250 OK")
		case strings.HasPrefix(upper, "MAIL FROM:"):
			from = strings.TrimSpace(cmd[10:])
			fmt.Fprintln(rw, "250 OK")
		case strings.HasPrefix(upper, "RCPT TO:"):
			to = strings.TrimSpace(cmd[8:])
			fmt.Fprintln(rw, "250 OK")
		case upper == "DATA":
			fmt.Fprintln(rw, "354 End data with <CR><LF>.<CR><LF>")
			_ = rw.Flush()
			var b strings.Builder
			for {
				l, e := rw.ReadString('\n')
				if e != nil {
					return
				}
				if strings.TrimSpace(l) == "." {
					break
				}
				b.WriteString(l)
			}
			_, e := s.Messages.Submit(ctx, domain.Message{TenantID: "smtp", From: strings.Trim(from, "<>"), To: strings.Trim(to, "<>"), Subject: "inbound", Text: b.String()})
			if e != nil {
				fmt.Fprintln(rw, "550", e)
			} else {
				fmt.Fprintln(rw, "250 queued")
			}
		case upper == "QUIT":
			fmt.Fprintln(rw, "221 bye")
			_ = rw.Flush()
			return
		default:
			fmt.Fprintln(rw, "502 command not implemented")
		}
		_ = rw.Flush()
	}
}
