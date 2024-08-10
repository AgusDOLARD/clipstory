package socket

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"sync"

	"golang.design/x/clipboard"
)

type Socket struct {
	addr     string
	listener net.Listener

	clipsLock sync.Mutex
	clips     []string
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

	go func() {
		newClip := clipboard.Watch(context.Background(), clipboard.FmtText)
		for clip := range newClip {
			s.clipsLock.Lock()
			s.clips = append(s.clips, string(clip))
			s.clipsLock.Unlock()
		}
	}()

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
	pkt := MarshalPacket(s.clips)
	conn.Write(pkt)
}

func GetClips(addr string) ([]string, error) {
	conn, err := net.Dial("unix", addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, conn)
	if err != nil {
		return nil, err
	}

	return UnmarshalPacket(buf.Bytes())
}
