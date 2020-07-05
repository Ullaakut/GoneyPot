package configuration

type Configuration struct {
	ICMP bool       `yaml:"icmp"`
	TCP  *TCPConfig `yaml:"tcp"`
	UDP  *UDPConfig `yaml:"udp"`
	// TODO: Reporting options.
}

type TCPConfig struct {
	Ports PortRanges `yaml:"ports"`
	// TODO: Fake service configuration.
}

// UnmarshalYAML implements the Unmarshaler interface of the yaml pkg.
func (t *TCPConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var yamlPortRangeSequence []string
	if err := unmarshal(&yamlPortRangeSequence); err != nil {
		return err
	}

	var err error
	t.Ports, err = NewPortRanges(yamlPortRangeSequence)
	if err != nil {
		return err
	}

	return nil
}

type UDPConfig struct {
	Ports PortRanges `yaml:"ports"`
	// TODO: Fake service configuration.
}

// UnmarshalYAML implements the Unmarshaler interface of the yaml pkg.
func (u *UDPConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var yamlPortRangeSequence []string
	if err := unmarshal(&yamlPortRangeSequence); err != nil {
		return err
	}

	var err error
	u.Ports, err = NewPortRanges(yamlPortRangeSequence)
	if err != nil {
		return err
	}

	return nil
}
