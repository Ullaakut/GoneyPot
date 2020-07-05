package listener

import (
	"bufio"
	"net"
	"strconv"
	"strings"
)

func (l *Listener) ListenTCP() error {
	tcp, err := net.Listen("tcp", ":"+strconv.Itoa(int(l.port)))
	if err != nil {
		l.report.Errorf("unable to TCP listen on port %d: %w", l.port, err)
		return err
	}

	go func() {
		for {
			c, err := tcp.Accept()
			if err != nil {
				l.report.Errorf("unable to accept TCP connection on port %d: %w", l.port, err)
				continue
			}

			go l.handleTCPConnection(c)
		}
	}()

	l.report.Infof("Listening for TCP connections on port %d", l.port)

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
