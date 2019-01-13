package device

import (
	"net"
	"os/exec"
	"regexp"
	"time"
)

type Dev struct {
	// houses all device information
	DeviceID         int
	CommonName       string
	Name             string
	Description      string
	Addresses        []Address
	TimeRegistered   time.Time
	TimeDeregistered time.Time
	PacketCount      int
	DeviceSniffs     []SniffProcess
}
type Address struct {
	// Used to house addresses of devices
	IP     net.IP
	Subnet net.IPMask
}

func GetNetworkDeviceInfo() ([]Dev, error) {
	// Uses t-hark to find sniff-able devices
	var devInfo []Dev
	var deviceFormat = regexp.MustCompile(`(\d+.\s)`)
	var interfacePatt = regexp.MustCompile(`(?:\w+)`)
	out, err := exec.Command("tshark", "-D").Output()
	if err != nil {
		return nil, err
	}
	sniffableDevices := deviceFormat.Split(string(out), -1)
	for id, device := range sniffableDevices {
		if device == "" {
			continue
		}
		var dev Dev
		dev.DeviceID = id
		dev.CommonName = device
		dev.Name = string(interfacePatt.Find([]byte(device)))
		devInfo = append(devInfo, dev)

	}
	return devInfo, err
}
