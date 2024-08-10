package socket

import (
	"fmt"
	"net"
	"os"
	"strings"
)

type Socket struct {
	addr     string
	listener net.Listener
	clips    []string
}

func NewSocket(addr string) *Socket {
	return &Socket{
		addr: addr,
	}
}

func (s *Socket) Start() error {
	listener, err := net.Listen("unix", s.addr)
	if err != nil {
		return err
	}
	s.listener = listener
	s.acceptLoop()
	return nil
}

func (s *Socket) Stop() error {
	if s.listener != nil {
		err := os.Remove(s.addr)
		if err != nil {
			return fmt.Errorf("failed to remove socket file: %w", err)
		}
		return s.listener.Close()
	}
	return nil
}

func (s *Socket) acceptLoop() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Errorf("failed to accept connection: %w", err)
			return
		}
		go s.handleConn(conn)
	}
}

func (s *Socket) handleConn(conn net.Conn) {
	defer conn.Close()
	clips := strings.Join(s.clips, "\n")
	conn.Write([]byte(clips))
}
