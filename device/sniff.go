package device

import (
	"fmt"
	"io"
	"log"
	"os/exec"

	"time"
)

type SniffProcess struct {
	PID           int
	Name          string
	StartTime     time.Time
	EndTime       time.Time
	FilePath      string
	Duration      int
	ContentType   []string
	IsPromiscuous bool
	Save          bool
	SaveLocation  string
}

//Sniff is an extention of any choses device
func (device Dev) Sniff() {
	sniffCommand := exec.Command("sudo", "tshark", "-I", "-i", device.Name, "-a", "duration:100")

	// Deal with output stream
	stdout, err := sniffCommand.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	stderr, err := sniffCommand.StderrPipe()
	if err != nil {
		log.Fatal()
	}

	sniffCommand.Start()
	go handleStream(stderr)
	go handleStream(stdout)
	fmt.Scanln()
}

//Sniff creates a new sniff taking in a string representation of the interface
func Sniff(interf string) {

	sniffCommand := exec.Command("sudo", "tshark", "-I", "-i", interf, "-a", "duration:100")

	// Deal with output stream
	stdout, err := sniffCommand.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	stderr, err := sniffCommand.StderrPipe()
	if err != nil {
		log.Fatal()
	}

	sniffCommand.Start()
	go handleStream(stderr)
	go handleStream(stdout)
	fmt.Scanln()
}

func handleStream(std io.ReadCloser) {
	b := make([]byte, 1)
	for {
		n, err := std.Read(b)
		fmt.Printf("%s", b[:n])

		if err == io.EOF {
			break
		}
	}
}
