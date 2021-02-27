package main

import (
	"fmt"
	"okane/database/ware"
	"okane/hardware"
	"okane/menu"
	"okane/sim"
	"okane/texts"
)

func main() {

	defer hardware.DeInitHardware()

	hardware.InitHardware()
	texts.Init()
	menu.Init()
	sim.Init()
	ware.Init()
	fmt.Println("OKANE GO!!")

	for true {
		menu.ExecuteCurrentPosition()
		key, b := hardware.GetKey()
		if b == false {
			break
		}
		switch key {
		case hardware.Key1, hardware.Key2, hardware.Key3, hardware.Key4, hardware.Key5, hardware.Key6, hardware.Key7, hardware.Key8, hardware.Key9:
			menu.GetNextMenuToShow(hardware.KeyToValue(key))
		}
		if key == hardware.KeyC {
			menu.GoBackToPrevious()
		}

	}

}
