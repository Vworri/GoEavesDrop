package parsers

import (
	"errors"
	"strconv"
	"strings"
)

type tcp struct {
	Bytes_on_wire    int
	Source_ip        int
	SourcePort       int
	DestinationPort  int
	content_type     string
	delta_tcp_stream float64
	delta_tcp_packet float64
	payload          string
}

func bytes_on_wire(tokens []string) (int, error) {
	bytes_on_wire_err := errors.New("No value for bytes_on_wire found")
	for i, token := range tokens {
		if val, err := strconv.Atoi(token); err == nil {
			if tokens[i+1] == "bytes" && tokens[i+2] == "on" && tokens[i+3] == "wire" {
				return val, nil
			}

		}

	}
	return 0, bytes_on_wire_err

}

func source_port(tokens []string) (int, error) {
	source_port_err := errors.New("No value for source_port found")
	for i, token := range tokens {
		if val, err := strconv.Atoi(token); err == nil {
			if tokens[i-1] == "Port:" && tokens[i-2] == "Source" {
				return val, nil
			}

		}

	}
	return 0, source_port_err

}
func destination_port(tokens []string) (int, error) {
	destination_port_err := errors.New("No value for destination port found")
	for i, token := range tokens {
		if val, err := strconv.Atoi(token); err == nil {
			if tokens[i-1] == "Port:" && tokens[i-2] == "Destination" {
				return val, nil
			}

		}

	}
	return 0, destination_port_err

}

func content_type(tokens []string) (string, error) {
	content_type_err := errors.New("No value for content type found")
	begining_idx := -1
	for i, token := range tokens {
		if tokens[i-1] == "Type:" && tokens[i-2] == "Content" {
			begining_idx = i
		}
		if begining_idx != -1 && strings.Index(token, ":") != -1 {
			return strings.Join(tokens[begining_idx:i], " "), nil
		}

	}
	return "", content_type_err

}

func delta_tcp_stream(tokens []string) (float64, error) {
	delta_tcp_stream_err := errors.New("No value for time from TCP stream start found")
	for i, token := range tokens {
		if tokens[i-1] == "stream:" && tokens[i-2] == "TCP" && tokens[i-3] == "this" && tokens[i-4] == "in" && tokens[i-5] == "frame" && tokens[i-6] == "first" {
			return strconv.ParseFloat(token, 64)
		}

	}
	return -1.0, delta_tcp_stream_err
}

func delta_tcp_packet(tokens []string) (float64, error) {
	delta_tcp_packet_err := errors.New("No value for time from TCP stream last packet found")
	for i, token := range tokens {
		if tokens[i-1] == "stream:" && tokens[i-2] == "TCP" && tokens[i-3] == "this" && tokens[i-4] == "in" && tokens[i-5] == "previous" {
			return strconv.ParseFloat(token, 64)
		}

	}
	return -1.0, delta_tcp_packet_err
}
