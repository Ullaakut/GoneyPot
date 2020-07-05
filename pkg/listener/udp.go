package listener

import (
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
			d := make([]byte, 16)
			_, source, err := c.ReadFromUDP(d)
			if err != nil {
				l.report.Event(source, nil, "unable to read from UDP connection on port %d", l.port)
			}

			l.report.Event(source, d, "UDP request received")
		}
	}()

	l.report.Infof("Listening for UDP connections on port %d", l.port)

	return nil
}
