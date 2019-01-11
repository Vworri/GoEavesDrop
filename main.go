package main

import (
	"github.com/Vworri/GoEavesDrop/device"
	"github.com/Vworri/GoEavesDrop/ui"
)

var app ui.App

func main() {
	app.SplashPage()
}

func init() {
	app = ui.CreateApplicaion()
	app.AvailableDevices = device.GetNetworkDeviceInfo()

}
