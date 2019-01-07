package device

import (
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
	var device_format = regexp.MustCompile(`(\d+.\s)`)
	var interface_patt = regexp.MustCompile(`(?:\w+)`)
	out, err := exec.Command("tshark", "-D").Output()

	sniffable_devices := device_format.Split(string(out), -1)
	if err != nil {
		log.Fatal(err)
	}
	for id, device := range sniffable_devices {
		if device == "" {
			continue
		}
		var dev Dev
		dev.DeviceID = id
		dev.CommonName = device
		dev.Name = string(interface_patt.Find([]byte(device)))
		devInfo = append(devInfo, dev)

	}
	return devInfo
}
