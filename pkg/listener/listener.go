package listener

import (
	"context"

	"github.com/Ullaakut/goneypot/pkg/reporter"
)

type Listener struct {
	ctx    context.Context
	report reporter.Reporter

	port uint16
}

func New(ctx context.Context, port uint16, reporter reporter.Reporter) *Listener {
	return &Listener{
		ctx:    ctx,
		report: reporter,
		port:   port,
	}
}
