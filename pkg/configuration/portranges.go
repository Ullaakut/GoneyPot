package configuration

import (
	"fmt"
	"strconv"
	"strings"
)

type PortRanges []PortRange

type PortRange [2]uint16

func NewPortRanges(s []string) (PortRanges, error) {
	var result PortRanges

	for _, portRange := range s {
		fmt.Println("Port range found:", s)

		ports := strings.Split(portRange, "-")

		// Case where only one port is specified.
		if len(ports) != 2 {
			port, err := strconv.Atoi(ports[0])
			if err != nil {
				return nil, fmt.Errorf("invalid port %q: %w", ports[0], err)
			}
			result = append(result, PortRange{uint16(port)})
			continue
		}

		// Case where a range is specified.
		start, err := strconv.Atoi(ports[0])
		if err != nil {
			return nil, fmt.Errorf("invalid port %q: %w", ports[0], err)
		}

		end, err := strconv.Atoi(ports[1])
		if err != nil {
			return nil, fmt.Errorf("invalid port %q: %w", ports[1], err)
		}

		// Ensure that the ports are sorted.
		if start > end {
			start, end = end, start
		}

		result = append(result, PortRange{
			uint16(start),
			uint16(end),
		})
	}

	fmt.Printf("Result: %+v\n", result)

	return result, nil
}
