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
	//process information for sniff commands
	name          string
	Process       *os.Process
	StartTime     time.Time
	EndTime       time.Time
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

func (device *Dev) Sniff() error {
	///dynamically creates a tshark sniff process
	// this is the core of Eavesdrop
	var sniff SniffProcess
	sniff.Command = exec.Command("sudo", "tshark", "-l", "-p", "-V", "-a", "duration:20", "-S", "\"===== THIS IS THE END ===\"")

	// Deal with output stream
	var err error
	sniff.OutputStream, err = sniff.Command.StdoutPipe()
	if err != nil {
		return err
	}
	sniff.ErrorStream, err = sniff.Command.StderrPipe()
	if err != nil {
		log.Fatal("Cannot pipe standard error of sniff")
	}
	device.DeviceSniffs = append(device.DeviceSniffs, sniff)
	return nil
}

func (sniff *SniffProcess) StopSniff() {
	// Kills the sniff command
	sniff.Process.Kill()
}

func (sniff *SniffProcess) handleStream() {
	b := make([]byte, 48)
	fmt.Println("Hello world!")
	var currentPacket string
	for {
		n, err := sniff.OutputStream.Read(b)
		line := fmt.Sprintf("%s", b[:n])

		if strings.Contains(line, "\"===== THIS IS THE END ===\"") == true {
			pac := new(packet.Packet)
			bforeAfter := strings.Split(line, "\"===== THIS IS THE END ===\"")
			currentPacket += bforeAfter[0]
			pac.Pac = currentPacket
			pac.Processed = false
			pac.Complete = true
			sniff.Queue = append(sniff.Queue, pac)
			fmt.Println(len(sniff.Queue))
			currentPacket = bforeAfter[1]
		} else {
			currentPacket += line
		}

		if err == io.EOF {
			break
		}
	}
}

func (sniff *SniffProcess) Start() {
	// Start go routines for the sniff
	go sniff.Command.Start()
	sniff.Command.Wait()
	go sniff.handleStream()
	fmt.Scanln()

}
