package configuration

import "errors"

type Configuration struct {
	ICMP bool       `yaml:"icmp"`
	TCP  *TCPConfig `yaml:"tcp"`
	UDP  *UDPConfig `yaml:"udp"`
	// TODO: Reporting options.
}

type TCPConfig struct {
	Ports PortRanges `mapstructure:"ports"`
	// TODO: Fake service configuration.
}

// UnmarshalMap implements the Unmarshaler interface.
func (t *TCPConfig) UnmarshalMap(value interface{}) error {
	var (
		tcpCfg map[string]interface{}
		ok     bool
	)
	if tcpCfg, ok = value.(map[string]interface{}); !ok {
		return errors.New("invalid TCP port range")
	}

	portRanges, ok := tcpCfg["ports"].([]string)
	if !ok {
		return errors.New("invalid TCP port range")
	}

	var err error
	t.Ports, err = NewPortRanges(portRanges)
	if err != nil {
		return err
	}

	return nil
}

type UDPConfig struct {
	Ports PortRanges `mapstructure:"ports"`
	// TODO: Fake service configuration.
}

// UnmarshalMap implements the Unmarshaler interface.
func (u *UDPConfig) UnmarshalMap(value interface{}) error {
	var (
		udpCfg map[string]interface{}
		ok     bool
	)
	if udpCfg, ok = value.(map[string]interface{}); !ok {
		return errors.New("invalid UDP port range")
	}

	portRanges, ok := udpCfg["ports"].([]string)
	if !ok {
		return errors.New("invalid UDP port range")
	}

	var err error
	t.Ports, err = NewPortRanges(portRanges)
	if err != nil {
		return err
	}
}
