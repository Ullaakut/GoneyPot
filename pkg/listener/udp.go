package listener

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

func (l *Listener) ListenUDP() error {
	s, err := net.ResolveUDPAddr("udp4", ":"+strconv.Itoa(int(l.port)))
	if err != nil {
		return fmt.Errorf("unable to resolve UDP address on port %d: %w", l.port, err)
	}

	go func() {
		c, err := net.ListenUDP("udp4", s)
		if err != nil {
			return
		}

		for {
			d, err := bufio.NewReader(c).ReadBytes('\n')
			if err != nil {
				l.report.Event(c.RemoteAddr(), nil, "unable to read from UDP connection on port %d", l.port)
			}

			l.report.Event(c.RemoteAddr(), d, "UDP request received")
		}
	}()

	fmt.Println("Listening for UDP connections on port", l.port)

	return nil
}
