package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Ullaakut/goneypot/pkg/listener"
	"github.com/Ullaakut/goneypot/pkg/reporter"
	"github.com/rs/zerolog"
)

func main() {
	ctx := context.Background()

	logger := reporter.NewZeroLog()
	logger.Level(zerolog.DebugLevel)

	l := listener.New(ctx, 8554, logger)

	if err := l.ListenTCP(); err != nil {
		fmt.Println("unable to listen TCP:", err)
		return
	}

	if err := l.ListenUDP(); err != nil {
		fmt.Println("unable to listen UDP:", err)
		return
	}

	if err := l.ListenICMP(); err != nil {
		fmt.Println("unable to listen ICMP:", err)
		return
	}

	// Wait for SIGINT.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	select {
	case <-c:
		fmt.Println("bye")
	}
}
