package main

import (
	"fmt"

	device "github.com/Vworri/GoEavesDrop/device"
)

func main() {
	//find all devices
	var NetInfo = device.GetNetworkDeviceInfo()
	for _, interf := range NetInfo {
		fmt.Println(interf.Name)
	}
}
