package ui

import (
	"github.com/vworri/GoEavesDrop/device"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type menu struct {
	*tview.Table
	SelectedDevice *device.Dev
	SelectedName   string
	currentPage    *tview.Grid
	details        deviceDetails
}

func (app *App) createDevicePage() {
	// This handles loading the device page
	// This keeps track of the current state of the application
	// by updating the different parts of the page when an option is
	// selected from the menu
	app.createMenu()
	app.handleMenuSelect()

	app.currentDevicePage = tview.NewFlex()
	app.currentDevicePage.AddItem(app.currentMenu.Table, 40, 1, true)

}

func (app *App) handleMenuSelect() {
	// Input handler for the device menu
	// when a device is selected, the navbar and the sniff slection will update
	app.currentMenu.Table.SetSelectedFunc(func(row int, column int) {
		app.QueueUpdateDraw(func() {
			app.reInitializeDevicePage(row)
			app.switchToNavBar()

		})

	})
}

func (app *App) reInitializeDevicePage(row int) {
	app.currentDevicePage.RemoveItem(app.currentMenu.currentPage)
	app.currentMenu.SelectedName = app.selectMenu(&app.currentMenu, row)
	app.currentMenu.SelectedDevice = app.retireveDevice(&app.currentMenu.SelectedName)
	app.buildDeviceDetails()
	app.currentDevicePage.AddItem(app.currentMenu.currentPage, 0, 1, true)
}

func (app *App) createMenu() {

	// Creates the actual table that functions as the device menu
	// this only creates the table
	// input hanldlers are housed in the handler function
	app.currentMenu.Table = tview.NewTable()
	app.currentMenu.SelectedDevice = app.getInitialDevice()
	app.buildDeviceDetails()
	app.currentMenu.Table.SetBorders(true)
	for idx, dev := range app.AvailableDevices {
		app.currentMenu.SetCell(idx, 0, tview.NewTableCell(dev.CommonName).
			SetAlign(tview.AlignCenter))
	}
	app.currentMenu.Table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
		if key == tcell.KeyEnter {
			app.currentMenu.SetSelectable(true, true)
		}
		if key == tcell.KeyTab {
			app.switchToNavBar()
		}
	})
}

func (app *App) loadDevicePage() {
	// simply loads the device page
	app.createDevicePage()
	if err := app.SetRoot(app.currentDevicePage, true).Run(); err != nil {
		panic(err)
	}
}
func (app *App) retireveDevice(selectedDevice *string) *device.Dev {
	for _, device := range app.AvailableDevices {
		if device.CommonName == *selectedDevice {
			return &device
		}
	}
	return nil
}

func (app *App) getInitialDevice() *device.Dev {
	// returns the initial state of the menu
	return &app.AvailableDevices[0]
}
func (app *App) selectMenu(men *menu, selectedDev int) string {
	// clears the formattting and recreates the menu when
	// a new oftion is selected
	men.Clear()
	for idx, dev := range app.AvailableDevices {
		men.SetCell(idx, 0, tview.NewTableCell(dev.CommonName).
			SetAlign(tview.AlignCenter))
	}
	cell := men.GetCell(selectedDev, 0).SetTextColor(tcell.ColorRed)
	return cell.Text
}
