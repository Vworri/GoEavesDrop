package aux

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

type Dev struct {
	Name             string
	Description      string
	Addresses        []Address
	TimeRegistered   time.Time
	TimeDeregistered time.Time
	PacketCount      int
	Handle           *pcap.Handle
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

func (device Dev) Sniff(promiscuous bool) {
	var err error
	if device.Handle, err = pcap.OpenLive(device.Name, 3200, promiscuous, pcap.BlockForever); err != nil {
		panic(err)
	} else {
		packetSource := gopacket.NewPacketSource(device.Handle, device.Handle.LinkType())
		for packet := range packetSource.Packets() {
			device.PacketCount += 1
			fmt.Println(packet) // Do something with a packet here.
		}
	}
}

func (device Dev) Kill() {
	device.Handle.Close()
}
