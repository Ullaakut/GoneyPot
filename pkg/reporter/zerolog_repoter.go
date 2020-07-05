package reporter

import (
	"net"
	"os"
	"time"

	"fmt"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type ZeroLogReporter struct {
	log zerolog.Logger
}

func NewZeroLog() ZeroLogReporter {
	return ZeroLogReporter{
		log: zerolog.New(os.Stdout),
	}
}

func (z ZeroLogReporter) Level(level zerolog.Level) {
	z.log.Level(level)
}

func (z ZeroLogReporter) Event(source net.Addr, packet []byte, format string, a ...interface{}) {
	eventID := uuid.New().String()

	log := z.log.Info().
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
	z.log.Debug().
		Str("id", eventID).
		Hex("packet", packet).Msg("")
}

func (z ZeroLogReporter) Infof(format string, a ...interface{}) {
	z.log.Info().Msgf(format, a)
}

func (z ZeroLogReporter) Errorf(format string, a ...interface{}) {
	z.log.Error().Err(fmt.Errorf(format, a)).Msg("")
}
