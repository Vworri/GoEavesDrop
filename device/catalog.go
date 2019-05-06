package device

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"regexp"
	"time"
)

type Dev struct {
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
	IP     net.IP
	Subnet net.IPMask
}

func GetNetworkDeviceInfo() []Dev {
	var devInfo []Dev
	var deviceFormat = regexp.MustCompile(`(\d+.\s)`)
	var interfacePatt = regexp.MustCompile(`(?:\w+)`)
	out, err := exec.Command("tshark", "-D").Output()

	sniffableDevices := deviceFormat.Split(string(out), -1)
	if err != nil {
		log.Fatal(err)
	}
	for id, device := range sniffableDevices {
		if device == "" {
			continue
		}
		var dev Dev
		dev.DeviceID = id
		dev.CommonName = device
		dev.Name = string(interfacePatt.Find([]byte(device)))
		dev.DeviceSniffs = append(dev.DeviceSniffs,
			SniffProcess{Name: fmt.Sprintf("Sniff # %s", device)})
		devInfo = append(devInfo, dev)

	}
	return devInfo
}
