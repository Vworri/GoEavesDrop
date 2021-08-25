package ui

import (
	"github.com/vworri/GoEavesDrop/device"
	"github.com/rivo/tview"
)

type App struct {
	*tview.Application
	AvailableDevices  []device.Dev
	currentDevicePage *tview.Flex
	currentMenu       menu
}

func CreateApplicaion() App {
	app := App{
		Application:      tview.NewApplication(),
		AvailableDevices: device.GetNetworkDeviceInfo()}

	return app

}
func newTextPrimitive(text string) *tview.TextView {
	return tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText(text).SetWrap(true)
}
