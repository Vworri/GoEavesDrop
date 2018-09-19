package aux

import (
	"log"
	"net"
	"time"

	"github.com/google/gopacket/pcap"
)

type Dev struct {
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

func GetNetworkDeviceInfo() []Dev {
	var devInfo []Dev
	devs, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	for _, device := range devs {
		var dev Dev
		dev.Name = device.Name
		dev.Description = device.Description
		dev.TimeDeregistered = time.Now()
		for _, address := range device.Addresses {
			var addr Address
			addr.IP = address.IP
			addr.Subnet = address.Netmask
			dev.Addresses = append(dev.Addresses, addr)
		}
		devInfo = append(devInfo, dev)
	}

	return devInfo
}
