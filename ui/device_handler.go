package ui

import (
	"log"

	"github.com/Vworri/GoEavesDrop/device"
	tui "github.com/marcusolsson/tui-go"
)

func DevicePage(device device.Dev) {
	box := tui.NewVBox(
		tui.NewLabel("device list"), tui.NewPadder(10, 1, tui.NewLabel(device.CommonName)),
	)
	device_page, err := tui.New(box)
	if err != nil {
		log.Fatal(err)
	}
	if err := device_page.Run(); err != nil {
		log.Fatal(err)
	}

}
