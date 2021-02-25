package menu

import (
	"os"
	"os/exec"
)

type Id int
type Mfunc func()

const (
	None Id = iota
	Main
	Sale
	ReadReports
	EraseReports
	Programming
	Configuration
	ServiceMode
	Log
	Communication
	DailyReport
	PeriodicReport
	WaresReport
)

var current Id = Main

func showMain() {
	show(Main)
}

func showReadReports() {
	show(ReadReports)
}

func showProgramming() {
	show(Programming)
}

func Init() {
	initPositions()
	MenuTree[Main] = Layout{PreviousMenuIndex: Main, positions: []Position{Positions[Sale], Positions[ReadReports], Positions[Programming]}}
	MenuTree[Sale] = Layout{PreviousMenuIndex: Main, positions: []Position{}}
	MenuTree[ReadReports] = Layout{PreviousMenuIndex: Main, positions: []Position{Positions[DailyReport], Positions[PeriodicReport], Positions[WaresReport]}}
	MenuTree[Programming] = Layout{PreviousMenuIndex: Main, positions: []Position{}}
}

func GetNextMenuToShow(next int) {
	if next > len(MenuTree[current].positions) {
		current = MenuTree[current].positions[0].id
	}

	current = MenuTree[current].positions[next-1].id
}

func show(id Id) {
	MenuTree[id].Show()
}

func GoBackToPrevious() {
	current = MenuTree[current].PreviousMenuIndex
}

func clearScren() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ExecuteCurrentPosition() {
	//clearScren()
	Positions[current].f()
	if Positions[current].layoutId != None {
		current = Positions[current].layoutId
		//clearScren()
		Positions[current].f()
	}
}
