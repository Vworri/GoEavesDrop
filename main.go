package main

import (
	device "github.com/Vworri/GoEavesDrop/device"
)

func main() {
	//find all devices
	var NetInfo = device.GetNetworkDeviceInfo()
	for _, interf := range NetInfo {
		if interf.Name == "wlo1" {
			interf.Sniff()
		}
	}

}
