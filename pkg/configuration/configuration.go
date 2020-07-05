package configuration

import (
	"encoding/json"
)

type Configuration struct {
	ICMP bool       `yaml:"icmp"`
	TCP  *TCPConfig `yaml:"tcp"`
	UDP  *UDPConfig `yaml:"udp"`

	// TODO: Reporting options.
	Debug bool
}

func (c Configuration) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

type TCPConfig struct {
	Ports []string `mapstructure:"ports"`
	// TODO: Fake service configuration.
}

func (t *TCPConfig) PortRanges() (PortRanges, error) {
	return NewPortRanges(t.Ports)
}

type UDPConfig struct {
	Ports []string `mapstructure:"ports"`
	// TODO: Fake service configuration.
}

func (u *UDPConfig) PortRanges() (PortRanges, error) {
	return NewPortRanges(u.Ports)
}
