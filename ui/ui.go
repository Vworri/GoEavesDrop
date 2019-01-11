package ui

import (
	"github.com/Vworri/GoEavesDrop/device"
	"github.com/rivo/tview"
)

type App struct {
	*tview.Application
	AvailableDevices []device.Dev
}

func CreateApplicaion() App {
	app := tview.NewApplication()

	return App{Application: app}

}
func newTextPrimitive(text string) *tview.TextView {
	return tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText(text)
}
