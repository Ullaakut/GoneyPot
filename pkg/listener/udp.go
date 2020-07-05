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
		println("1")
		c, err := net.ListenUDP("udp4", s)
		if err != nil {
			return
		}
		println("2")

		for {
			println("3")
			d := make([]byte, 1024*1024)
			_, source, err := c.ReadFromUDP(d)
			if err != nil {
				println("4.1")
				l.report.Event(source, nil, "unable to read from UDP connection on port %d", l.port)
			}

			println("4.2")
			l.report.Event(source, d, "UDP request received")
		}
	}()

	fmt.Println("Listening for UDP connections on port", l.port)

	return nil
}
