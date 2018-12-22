package device

import (
	"fmt"
	"io"
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
	DeviceSniffs     []SniffProcess
}
type Address struct {
	IP     net.IP
	Subnet net.IPMask
}

type SniffProcess struct {
	PID         int
	Start_Time  time.Time
	End_Time    time.Time
	FilePath    string
	Duration    int
	ContentType []string
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

func (device Dev) Sniff() {
	sniff_command := exec.Command("sudo", "tshark", "-iwlo1", "-aduration:100")

	// Deal with output stream
	stdout, err := sniff_command.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	stderr, err := sniff_command.StderrPipe()
	if err != nil {
		log.Fatal()
	}

	sniff_command.Start()
	go handle_stream(stderr)
	go handle_stream(stdout)
	fmt.Scanln()
}

func handle_stream(std io.ReadCloser) {
	b := make([]byte, 1)
	for {
		n, err := std.Read(b)
		fmt.Printf("%s", b[:n])

		if err == io.EOF {
			break
		}
	}
}
