package main

import (
	"fmt"

	aux "github.com/GoEavesDrop/Aux"
	ui "github.com/gizak/termui" // <- ui shortcut, optional
)

func main() {
	//find all devices
	var NetInfo = aux.GetNetworkDeviceInfo()
	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()
	var menu []string
	for i, netInt := range NetInfo {
		menu = append(menu, fmt.Sprintf("[%d] %s\n", i, netInt.Name))
	}
	ls := ui.NewList()
	ls.Items = menu
	ls.ItemFgColor = ui.ColorYellow
	ls.BorderLabel = "List"
	ls.Height = 7
	ls.Width = 25
	ls.Y = 0

	ui.Render(ls)
	ui.Handle("q", func(ui.Event) {
		// press q to quit
		ui.StopLoop()
	})
	ui.Loop()
}
