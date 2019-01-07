package device

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/Vworri/GoEavesDrop/packet"
)

type SniffProcess struct {
	name          string
	Process       *os.Process
	Start_Time    time.Time
	End_Time      time.Time
	FilePath      string
	Duration      int
	ContentType   []string
	Command       *exec.Cmd
	IsPromiscuous bool
	Save          bool
	SaveLocation  string
	OutputStream  io.ReadCloser
	ErrorStream   io.ReadCloser
	Queue         []*packet.Packet
}

func (device Dev) Sniff() {
	var sniff SniffProcess
	sniff.Command = exec.Command("sudo", "tshark", "-l", "-p", "-V", "-a", "duration:20", "-S", "\"===== THIS IS THE END ===\"")

	// Deal with output stream
	var err error
	sniff.OutputStream, err = sniff.Command.StdoutPipe()
	if err != nil {
		log.Fatal("Cannot pipe standard output of sniff")
	}
	sniff.ErrorStream, err = sniff.Command.StderrPipe()
	if err != nil {
		log.Fatal("Cannot pipe standard error of sniff")
	}
	device.DeviceSniffs = append(device.DeviceSniffs, sniff)
	return
}

func (sniff SniffProcess) StopSniff() {
	sniff.Process.Kill()
}

func (sniff SniffProcess) handle_stream() {
	b := make([]byte, 48)
	var current_packet string
	for {
		n, err := sniff.OutputStream.Read(b)
		line := fmt.Sprintf("%s", b[:n])

		if strings.Contains(line, "\"===== THIS IS THE END ===\"") == true {
			pac := new(packet.Packet)
			bfore_after := strings.Split(line, "\"===== THIS IS THE END ===\"")
			current_packet += bfore_after[0]
			pac.Pac = current_packet
			pac.Processed = false
			pac.Complete = true
			sniff.Queue = append(sniff.Queue, pac)
			fmt.Println(len(sniff.Queue))
			current_packet = bfore_after[1]
		} else {
			current_packet += line
		}

		if err == io.EOF {
			break
		}
	}
}
