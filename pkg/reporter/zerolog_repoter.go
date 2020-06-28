package reporter

import (
	"net"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type ZeroLogReporter struct {
	zerolog.Logger
}

func NewZeroLog() ZeroLogReporter {
	return ZeroLogReporter{
		Logger: zerolog.New(os.Stdout),
	}
}

func (z ZeroLogReporter) Event(source net.Addr, packet []byte, format string, a ...interface{}) {
	eventID := uuid.New().String()

	log := z.Info().
		Str("id", eventID).
		Time("time", time.Now()).
		Int("packet_length", len(packet))

	// Always print essential info.
	if source != nil {
		log = log.Str("source", source.String())
	}

	if len(a) > 0 {
		log.Msgf(format, a)
	} else {
		log.Msg(format)
	}

	// Print packet contents only in debug.
	z.Debug().
		Str("id", eventID).
		Hex("packet", packet).Msg("")
}
