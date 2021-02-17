package main

import (
	"fmt"
	//"okane2/src/database/ware"
	//"okane2/src/menu"
	"okane/database/utils"
	"okane/database/ware"
	"okane/hardware"
	"okane/menu"
)

func main() {

	defer hardware.DeInitHardware()

	//var Jablko ware.Ware
	//Jablko.Id = 1
	//Jablko.Name = "Jabłko"
	//Jablko.Vat = ware.Vat('A')
	//Jablko.Price = 12343

	//fmt.Println(ware.ToString(Jablko))

	//menu.ShowMenu(menu.MainMenu())
	hardware.InitHardware()
	fmt.Println("OKANE GO!!")

	var p utils.Price

	fmt.Println(p.String())

	var w ware.Ware

	fmt.Println(ware.ToString(w))

	menu.ShowMenu(menu.MainMenu())

	fmt.Println("Press ESC to quit")
	for true {
		key, b := hardware.GetKey()
		if b == false {
			break
		}
		if key == hardware.Key1 {
			menu.ShowMenu(menu.SaleMenu())
		}
		if key == hardware.Key2 {
			fmt.Println("2")
		}

	}

}
