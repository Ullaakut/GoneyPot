package reporter

import (
	"net"
)

type Reporter interface {
	Info(string)
	Infof(string, ...interface{})
	Errorf(string, ...interface{})
	Event(source net.Addr, packet []byte, format string, a ...interface{})
}
