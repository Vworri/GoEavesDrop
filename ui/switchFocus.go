package ui

func (app *App) switchToNavBar() {
	// switched focus to nav bar
	app.handleNavbar()
	app.handleNavBarSelect()
	app.SetFocus(app.currentMenu.currentPage)
}

func (app *App) switchToMenu() {
	//switches focus to device menu
	app.currentMenu.SetSelectable(true, true)
	app.SetFocus(app.currentMenu.Table)
}

func (app *App) switchToDetails() {
	app.SetFocus(app.currentMenu.details.Info)

}
