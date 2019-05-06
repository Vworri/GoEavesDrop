package ui

import (
	"github.com/rivo/tview"
)

func (app *App) buildSniffDetails() {
	// returns the device details when a device is selected from the
	//menu packaged and formatted
	app.generateNavBar()

	page := tview.NewGrid().
		SetRows(3, -1).
		SetColumns(0, 0).SetBorders(true)
	page.AddItem(app.currentMenu.details.navBar, 0, 0, 1, 3, 0, 0, true).AddItem(app.currentMenu.details.Info, 1, 0, 1, 3, 0, 0, true)
	app.currentMenu.currentPage = page
}
