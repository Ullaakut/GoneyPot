package configuration

type Configuration struct {
	ICMP bool       `yaml:"icmp"`
	TCP  *TCPConfig `yaml:"tcp"`
	UDP  *UDPConfig `yaml:"udp"`
	// TODO: Reporting options.
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
