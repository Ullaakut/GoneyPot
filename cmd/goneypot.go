package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Ullaakut/goneypot/pkg/configuration"
	"github.com/Ullaakut/goneypot/pkg/listener"
	"github.com/Ullaakut/goneypot/pkg/reporter"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config.yaml")
	viper.AddConfigPath("/etc/goneypot/")
	viper.AddConfigPath("$HOME/.config/goneypot/")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("unable to read configuration file:", err)
		os.Exit(1)
	}

	var cfg configuration.Configuration
	err := viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Println("invalid configuration:", err)
		os.Exit(1)
	}

	logger := reporter.NewZeroLog()
	logger.Level(zerolog.DebugLevel)

	ctx := context.Background()

	if cfg.ICMP {
		l := listener.New(ctx, 0, logger)

		if err := l.ListenICMP(); err != nil {
			fmt.Println("unable to listen ICMP:", err)
			return
		}
	}

	if cfg.TCP != nil {
		for _, portRange := range cfg.TCP.Ports {
			// Single port, not a range.
			if portRange[1] == 0 {
				if err := listener.New(ctx, portRange[0], logger).ListenTCP(); err != nil {
					fmt.Println("unable to listen TCP:", err)
					return
				}

				continue
			}

			for port := portRange[0]; port <= portRange[1]; port++ {
				if err := listener.New(ctx, port, logger).ListenTCP(); err != nil {
					fmt.Println("unable to listen TCP:", err)
					return
				}
			}
		}
	}

	if cfg.UDP != nil {
		for _, portRange := range cfg.UDP.Ports {
			// Single port, not a range.
			if portRange[1] == 0 {
				if err := listener.New(ctx, portRange[0], logger).ListenUDP(); err != nil {
					fmt.Println("unable to listen UDP:", err)
					return
				}

				continue
			}

			for port := portRange[0]; port <= portRange[1]; port++ {
				if err := listener.New(ctx, port, logger).ListenUDP(); err != nil {
					fmt.Println("unable to listen UDP:", err)
					return
				}
			}
		}
	}

	// Wait for SIGINT.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	select {
	case <-c:
		os.Exit(0)
	}
}
