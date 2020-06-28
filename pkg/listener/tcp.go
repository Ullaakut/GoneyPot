package listener

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func (l *Listener) ListenTCP() error {
	tcp, err := net.Listen("tcp", ":"+strconv.Itoa(int(l.port)))
	if err != nil {
		return fmt.Errorf("unable to TCP listen on port %d: %w", l.port, err)
	}

	go func() {
		for {
			c, err := tcp.Accept()
			if err != nil {
				fmt.Println("unable to accept TCP")
				continue
			}

			go l.handleTCPConnection(c)
		}
	}()

	return nil
}

func (l *Listener) handleTCPConnection(c net.Conn) {
	defer func() {
		_ = c.Close()
	}()

	d, err := bufio.NewReader(c).ReadBytes('\n')
	if err != nil {
		l.report.Event(nil, nil, "unable to read from TCP connection on port %d", l.port)
		return
	}

	if strings.TrimSpace(string(d)) == "STOP" {
		l.report.Event(c.RemoteAddr(), d, "TCP STOP request received")
		return
	}

	l.report.Event(c.RemoteAddr(), d, "TCP request received")
}
