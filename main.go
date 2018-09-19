package main

import (
	"fmt"

	aux "github.com/GoEavesDrop/Aux"
)

func main() {
	//find all devices
	var NetInfo = aux.GetNetworkDeviceInfo()
	fmt.Print(NetInfo)
}
