package parsers

import (
	"errors"
	"strconv"
	"strings"
)

type tcp struct {
	bytesOnWire     int
	SourceIP      string
	SourcePort      int
	DestinationIP 	string
	DestinationPort int
	contentType     string
	deltaTCPStream  float64
	deltaTCPPacket  float64
	payload         string
}

func bytesOnWire(tokens []string) (int, error) {
	bytesOnWireErr := errors.New("No value for bytes_on_wire found")
	for i, token := range tokens {
		if val, err := strconv.Atoi(token); err == nil {
			if tokens[i+1] == "bytes" && tokens[i+2] == "on" && tokens[i+3] == "wire" {
				return val, nil
			}

		}

	}
	return 0, bytesOnWireErr

}

func sourcePort(tokens []string) (int, error) {
	sourcePortErr := errors.New("No value for source_port found")
	for i, token := range tokens {
		if val, err := strconv.Atoi(token); err == nil {
			if tokens[i-1] == "Port:" && tokens[i-2] == "Source" {
				return val, nil
			}

		}

	}
	return 0, sourcePortErr

}
func destinationPort(tokens []string) (int, error) {
	destinationPortErr := errors.New("No value for destination port found")
	for i, token := range tokens {
		if val, err := strconv.Atoi(token); err == nil {
			if tokens[i-1] == "Port:" && tokens[i-2] == "Destination" {
				return val, nil
			}

		}

	}
	return 0, destinationPortErr

}

func content_type(tokens []string) (string, error) {
	contentTypeErr := errors.New("No value for content type found")
	beginingIdx := -1
	for i, token := range tokens {
		if tokens[i-1] == "Type:" && tokens[i-2] == "Content" {
			beginingIdx = i
		}
		if beginingIdx != -1 && strings.Index(token, ":") != -1 {
			return strings.Join(tokens[beginingIdx:i], " "), nil
		}

	}
	return "", contentTypeErr

}

func deltaTCPStream(tokens []string) (float64, error) {
	deltaTCPStreamErr := errors.New("No value for time from TCP stream start found")
	for i, token := range tokens {
		if tokens[i-1] == "stream:" && tokens[i-2] == "TCP" && tokens[i-3] == "this" && tokens[i-4] == "in" && tokens[i-5] == "frame" && tokens[i-6] == "first" {
			return strconv.ParseFloat(token, 64)
		}

	}
	return -1.0, deltaTCPStreamErr
}

func deltaTCPPacket(tokens []string) (float64, error) {
	deltaTCPPacketErr := errors.New("No value for time from TCP stream last packet found")
	for i, token := range tokens {
		if tokens[i-1] == "stream:" && tokens[i-2] == "TCP" && tokens[i-3] == "this" && tokens[i-4] == "in" && tokens[i-5] == "previous" {
			return strconv.ParseFloat(token, 64)
		}

	}
	return -1.0, deltaTCPPacketErr
}
