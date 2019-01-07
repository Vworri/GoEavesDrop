package ui

import "github.com/rivo/tview"

func CreateApplicaion() {
	app := tview.NewApplication()
	splash := SplashPage()
	if err := app.SetRoot(splash, true).Run(); err != nil {
		panic(err)
	}
}
