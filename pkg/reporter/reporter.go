package reporter

import (
	"net"
)

type Reporter interface {
	Event(source net.Addr, packet []byte, format string, a ...interface{})
}
