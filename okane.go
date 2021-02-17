package main

import (
	"fmt"
	"okane/hardware"
	"okane/menu"
)

func main() {

	defer hardware.DeInitHardware()

	hardware.InitHardware()
	menu.InitMenu()
	fmt.Println("OKANE GO!!")

	currentMenu := menu.MainMenu

	for true {
		menu.ShowMenu(menu.MenuTree[currentMenu].Show())
		key, b := hardware.GetKey()
		if b == false {
			break
		}
		fmt.Println(key)
		switch key {
		case hardware.Key1, hardware.Key2, hardware.Key3, hardware.Key4, hardware.Key5, hardware.Key6, hardware.Key7, hardware.Key8, hardware.Key9:
			currentMenu = menu.GetNextMenuToShow(hardware.KeyToValue(key), currentMenu)
		}
		if key == hardware.KeyC {
			currentMenu = menu.MenuTree[currentMenu].PreviousMenuIndex
		}

	}

}
