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
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
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

	fmt.Println("Configuration loaded", cfg.String())

	logger := reporter.NewZeroLog()
	logger.Level(zerolog.DebugLevel)
	ctx := context.Background()

	if cfg.ICMP {
		l := listener.New(ctx, 0, logger)

		if err := l.ListenICMP(); err != nil {
			os.Exit(1)
		}
	}

	if cfg.TCP != nil {
		portRanges, err := cfg.TCP.PortRanges()
		if err != nil {
			fmt.Println("invalid TCP port range:", err)
			os.Exit(1)
		}

		for _, portRange := range portRanges {
			// Single port, not a range.
			if portRange[1] == 0 {
				if err := listener.New(ctx, portRange[0], logger).ListenTCP(); err != nil {
					os.Exit(1)
				}

				continue
			}

			for port := portRange[0]; port <= portRange[1]; port++ {
				if err := listener.New(ctx, port, logger).ListenTCP(); err != nil {
					os.Exit(1)
				}
			}
		}
	}

	if cfg.UDP != nil {
		portRanges, err := cfg.UDP.PortRanges()
		if err != nil {
			fmt.Println("invalid UDP port range:", err)
			os.Exit(1)
		}

		for _, portRange := range portRanges {
			// Single port, not a range.
			if portRange[1] == 0 {
				if err := listener.New(ctx, portRange[0], logger).ListenUDP(); err != nil {
					os.Exit(1)
				}

				continue
			}

			for port := portRange[0]; port <= portRange[1]; port++ {
				if err := listener.New(ctx, port, logger).ListenUDP(); err != nil {
					os.Exit(1)
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
