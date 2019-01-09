package main

import (
	"log"

	"github.com/Vworri/GoEavesDrop/device"
	tui "github.com/marcusolsson/tui-go"
)

func main() {

	available_devices := device.GetNetworkDeviceInfo()
	l := tui.NewList()
	l.SetFocused(true)
	for _, interf := range available_devices {
		l.AddItems(interf.CommonName)
	}
	l.SetSelected(0)
	box := tui.NewVBox(
		tui.NewLabel("tui-go"), l,
	)
	ui, err := tui.New(box)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("p", func() {
		ui.Quit()
		selcted := l.Selected()
		available_devices[selcted].Sniff()
		available_devices[selcted].DeviceSniffs[0].Start() //assuming only one sniff for now

	})

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
