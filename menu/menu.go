package menu

import (
	"fmt"
)

//const

func MainMenu() string {
	buf := ""
	for i, s := range GetMainMenuTexts() {
		buf += fmt.Sprintf("%d. %s\n", i+1, s)
	}
	return buf
}

func SaleMenu() string {
	buf := ""
	for i, s := range GetSaleMenuTexts() {
		buf += fmt.Sprintf("%d. %s\n", i+1, s)
	}
	return buf
}

func ShowMenu(buf string) {
	fmt.Println(buf)
}
