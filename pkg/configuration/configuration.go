package configuration

type Configuration struct {
	ICMP bool
	TCP  *TCPConfig
	UDP  *UDPConfig
}

type TCPConfig struct {
	Ports PortRanges
}

type UDPConfig struct {
	Ports PortRanges
}
