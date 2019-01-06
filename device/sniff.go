package device

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

type SniffProcess struct {
	PID           int
	Start_Time    time.Time
	End_Time      time.Time
	FilePath      string
	Duration      int
	ContentType   []string
	IsPromiscuous bool
	Save          bool
	SaveLocation  string
}

func (device Dev) Sniff() {
	sniff_command := exec.Command("sudo", "tshark", "-l", "-p", "-V", "-a", "duration:100", "-S", "\"===== THIS IS THE END ===\"")
	// Deal with output stream
	sniff_command.Stdout = os.Stdout
	sniff_command.Stderr = os.Stderr
	go sniff_command.Start()
	sniff_command.Wait()
	// fmt.Scanln()
}

func Sniff(interf string) {
	sniff_command := exec.Command("sudo", "tshark", "-i", interf, "-p", "-V", "-a", "duration:100", "-S", "\"===== THIS IS THE END ===\"")

	// Deal with output stream
	sniff_command.Stdout = os.Stdout
	sniff_command.Stderr = os.Stderr
	sniff_command.Start()
	fmt.Scanln()
}

func handle_stream(std io.ReadCloser) {
	b := make([]byte, 8)
	for {
		n, err := std.Read(b)
		fmt.Printf("%s", b[:n])

		if err == io.EOF {
			break
		}
	}
}
