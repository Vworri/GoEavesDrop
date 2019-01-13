package ui

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type deviceDetails struct {
	navBar *tview.Table
	Info   *tview.Form
}

func (men *menu) buildDeviceDetails() {
	// returns the device details when a device is selected from the
	//menu packaged and formatted
	men.generateNavBar()

	page := tview.NewGrid().
		SetRows(3, -1).
		SetColumns(0, 0).SetBorders(true)
	page.AddItem(men.details.navBar, 0, 0, 1, 3, 0, 0, true).AddItem(men.details.Info, 1, 0, 1, 3, 0, 0, true)
	men.currentPage = page
}

func (men *menu) generateNavBar() {
	//Builds the navigation bar and the device information page
	devDash := deviceDetails{navBar: tview.NewTable(),
		Info: men.generateDeviceInfoPage()}

	devDash.navBar.SetBorders(true)

	devDash.navBar.SetCellSimple(0, 0, fmt.Sprintf("%s Info", men.SelectedDevice.Name))

	for idx, sniff := range men.SelectedDevice.DeviceSniffs {
		devDash.navBar.SetCellSimple(0, idx+1, sniff.Name)

	}

	men.details = devDash
}

func (app *App) handleNavbar() {
	// handles navbar navigation
	app.currentMenu.details.navBar.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.switchToMenu()
		}
		if key == tcell.KeyEnter {
			app.currentMenu.details.navBar.SetSelectable(true, true)

		}
		if key == tcell.KeyTab {
			app.switchToDetails()
		}
	})
}

func (men menu) generateDeviceInfoPage() *tview.Form {

	form := tview.NewForm().
		AddDropDown("Title", []string{"Mr.", "Ms.", "Mrs.", "Dr.", "Prof."}, 0, nil).
		AddInputField("First name", "", 20, nil, nil).
		AddInputField("Last name", "", 20, nil, nil).
		AddCheckbox("Age 18+", false, nil).
		AddPasswordField("Password", "", 10, '*', nil).
		AddButton("Save", nil).
		AddButton("Quit", nil)
	form.SetBorder(true).SetTitle("Enter some data").SetTitleAlign(tview.AlignLeft)

	return form

}

func (app *App) handleNavBarSelect() {
	// handles what to do when a nav bar cell is selected
	app.currentMenu.details.navBar.SetSelectedFunc(func(row int, column int) {
		app.QueueUpdateDraw(func() {
			app.currentMenu.details.navBar.GetCell(row, column).SetTextColor(tcell.ColorRed)
		})

	})
}
