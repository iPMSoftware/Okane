package menu

import (
	"fmt"
)

type MenuId int

const (
	NoneMenu MenuId = iota
	MainMenu
	SaleMenu
	ReadReportsMenu
	EraseReportsMenu
	ProgrammingMenu
	ConfigurationMenu
	ServiceModeMenu
	Log
	Communication
)

type MenuPosition struct {
	PreviousMenuIndex MenuId
	Show              func() string
	Positions         []MenuId
}

var MenuTree = map[MenuId]MenuPosition{}

func InitMenu() {
	MenuTree[MainMenu] = MenuPosition{PreviousMenuIndex: MainMenu, Show: ShowMainMenu, Positions: []MenuId{SaleMenu, ReadReportsMenu, ProgrammingMenu}}
	MenuTree[SaleMenu] = MenuPosition{PreviousMenuIndex: MainMenu, Show: ShowSaleMenu, Positions: []MenuId{}}
}

func GetNextMenuToShow(next int, current MenuId) MenuId {
	if next > len(MenuTree[current].Positions) {
		return current
	}

	return MenuTree[current].Positions[next-1]
}

func ShowMainMenu() string {
	buf := ""
	for i, s := range GetMainMenuTexts() {
		buf += fmt.Sprintf("%d. %s\n", i+1, s)
	}
	return buf
}

func ShowSaleMenu() string {
	buf := ""
	for i, s := range GetSaleMenuTexts() {
		buf += fmt.Sprintf("%d. %s\n", i+1, s)
	}
	return buf
}

func ShowMenu(buf string) {
	fmt.Println(buf)
}
