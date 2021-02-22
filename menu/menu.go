package menu

import (
	"fmt"
	"okane/mock"
)

type Id int
type Mfunc func()

const (
	NoneMenu Id = iota
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
)

type Position struct {
	Text string
	id   Id
	f    Mfunc
}

type Layout struct {
	PreviousMenuIndex Id
	positions         []Position
}

var current Id = Main

func (l *Layout) getPositions() []string {
	pos := []string{}
	for _, v := range l.positions {
		pos = append(pos, v.Text)
	}
	return pos
}

func (l Layout) Show() {
	positions := l.getPositions()
	buf := "=====================\n"
	i := 1
	for _, elem := range positions {
		buf += fmt.Sprintf("%d. %s\n", i, elem)
		i++
	}
	buf += "=====================\n"
	fmt.Println(buf)
}

var MenuTree = map[Id]Layout{}
var Positions = map[Id]Position{}

func showMain() {
	show(Main)
}

func showReadReports() {
	show(ReadReports)
}

func showProgramming() {
	show(Programming)
}

func initPositions() {
	Positions[Main] = Position{Text: "", id: Main, f: showMain}
	Positions[Sale] = Position{Text: "Sprzedaż", id: Sale, f: mock.Sale}
	Positions[ReadReports] = Position{Text: "Raporty czytające", id: ReadReports, f: showReadReports}
	Positions[Programming] = Position{Text: "Programowanie", id: Programming, f: showProgramming}

	Positions[DailyReport] = Position{Text: "Raport dobowy", id: DailyReport, f: mock.DailyReport}
	Positions[PeriodicReport] = Position{Text: "Raport okresowy", id: PeriodicReport, f: mock.PeriodicReport}
}

func Init() {
	initPositions()
	MenuTree[Main] = Layout{PreviousMenuIndex: Main, positions: []Position{Positions[Sale], Positions[ReadReports], Positions[Programming]}}
	MenuTree[Sale] = Layout{PreviousMenuIndex: Main, positions: []Position{}}
	MenuTree[ReadReports] = Layout{PreviousMenuIndex: Main, positions: []Position{Positions[DailyReport], Positions[PeriodicReport]}}
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

func ExecuteCurrentPosition() {
	Positions[current].f()
}
