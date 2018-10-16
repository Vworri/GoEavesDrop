package main

import (
	"fmt"

	aux "github.com/GoEavesDrop/device"
)

func main() {
	//find all devices
	var NetInfo = aux.GetNetworkDeviceInfo()
	for _, interf := range NetInfo {
		fmt.Println(interf.Name)
		if interf.Name == "wlo1" {
			interf.Sniff(true)
		}
	}
}
