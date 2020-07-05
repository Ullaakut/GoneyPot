package listener

import (
	"time"

	"golang.org/x/net/icmp"
)

func (l *Listener) ListenICMP() error {
	conn, err := icmp.ListenPacket("udp4", "localhost")
	if err != nil {
		l.report.Errorf("unable to listen ICMP: %w", err)
		return err
	}

	go func() {
		for {
			time.Sleep(1 * time.Second)

			msg := make([]byte, 256)
			length, sourceIP, err := conn.ReadFrom(msg)
			if err != nil {
				l.report.Errorf("error when reading ICMP: %w", err)
				continue
			}

			l.report.Event(sourceIP, msg[:length], "Received ICMP request")
		}
	}()

	l.report.Info("Listening for ICMP requests on host")

	return nil
}
