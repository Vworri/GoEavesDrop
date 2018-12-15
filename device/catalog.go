package device

import (
	"log"
	"net"
	"os/exec"
	"regexp"
	"time"
)

type Dev struct {
	Common_Name      string
	Name             string
	Description      string
	Addresses        []Address
	TimeRegistered   time.Time
	TimeDeregistered time.Time
	PacketCount      int
}
type Address struct {
	IP     net.IP
	Subnet net.IPMask
}

type SniffProcess struct {
	PID        int
	Start_Time time.Time
	End_Time   time.Time
	FilePath   string
	Duration   int
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
	for _, device := range sniffable_devices {
		var dev Dev
		dev.Common_Name = device
		dev.Name = string(interface_patt.Find([]byte(device)))
		devInfo = append(devInfo, dev)
	}

	return devInfo
}

func (device Dev) Sniff(promiscuous bool) {

}
