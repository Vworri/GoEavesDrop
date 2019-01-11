package ui

import (
	"fmt"

	"github.com/Vworri/GoEavesDrop/device"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type menu struct {
	SelectedDevice *device.Dev
	SelectedRow    int
	devicePages    *tview.Pages
}

func (app *App) loadDevicesPage() {
	deviceMenu := tview.NewTable().
		SetBorders(true)
	for idx, dev := range app.AvailableDevices {
		deviceMenu.SetCell(idx, 0, tview.NewTableCell(dev.CommonName).
			SetAlign(tview.AlignCenter))
		deviceMenu.devicePages.AddPage(fmt.Sprintf("page-%d", dev.CommonName))
	}

	deviceMenu.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
		if key == tcell.KeyEnter {
			deviceMenu.SetSelectable(true, true)
		}
	}).SetSelectedFunc(func(row int, column int) {
		app.ClearMenu(deviceMenu, row)
	})
	flex := tview.NewFlex()
	flex.AddItem(deviceMenu, 0, 1, true)
	// AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
	// 	AddItem(tview.NewBox().SetBorder(true).SetTitle("Top"), 0, 1, false).
	// 	AddItem(tview.NewForm().SetBorder(true).SetTitle("Middle (3 x height of Top)"), 0, 3, false).
	// 	AddItem(tview.NewBox().SetBorder(true).SetTitle("Bottom (5 rows)"), 5, 1, false), 0, 2, false).
	// 	AddItem(tview.NewBox().SetBorder(true).SetTitle("Right (20 cols)"), 20, 1, false)
	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}


func (app *App) retireveDevice(selectedDevice string) *device.Dev {
	for _, device := range app.AvailableDevices {
		if device.CommonName == selectedDevice {
			return &device
		}
	}
	return nil
}

func (app App) ClearMenu(men *tview.Table, selectedDev int) {
	men.Clear()
	for idx, dev := range app.AvailableDevices {

		men.SetCell(idx, 0, tview.NewTableCell(dev.CommonName).
			SetAlign(tview.AlignCenter))
	}
	men.GetCell(selectedDev, 0).SetTextColor(tcell.ColorRed)
}
