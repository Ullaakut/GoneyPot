package listener

import (
	"fmt"
	"time"

	"golang.org/x/net/icmp"
)

func (l *Listener) ListenICMP() error {
	conn, err := icmp.ListenPacket("udp4", "localhost")
	if err != nil {
		return err
	}

	go func() {
		for {
			time.Sleep(1 * time.Second)

			msg := make([]byte, 256)
			length, sourceIP, err := conn.ReadFrom(msg)
			if err != nil {
				fmt.Println("error when reading ICMP")
				continue
			}

			l.report.Event(sourceIP, msg[:length], "Received ICMP request")
		}
	}()

	return nil
}
