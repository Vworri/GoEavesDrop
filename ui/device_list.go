package ui

import (
	"github.com/rivo/tview"
)

func deviceMenu(devices []string) *tview.DropDown {

	dropdown := tview.NewDropDown().
		SetLabel("Select an interface (hit Enter): ").
		SetOptions(devices, nil)
	return dropdown
}
