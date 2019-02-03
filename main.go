package main

import (
	"log"

	"github.com/Vworri/GoEavesDrop/device"
	tui "github.com/marcusolsson/tui-go"
)

func main() {

	availableDevices, _ := device.GetNetworkDeviceInfo()
	l := tui.NewList()
	l.SetFocused(true)
	for _, interf := range availableDevices {
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
		availableDevices[selcted].Sniff()
		availableDevices[selcted].DeviceSniffs[0].Start() //assuming only one sniff for now

	})

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
