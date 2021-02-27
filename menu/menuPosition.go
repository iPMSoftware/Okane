package menu

import (
	"okane/configuration"
	"okane/mock"
	"okane/reports"
)

type Position struct {
	Text     string
	id       Id
	layoutId Id
	f        Mfunc
}

var Positions = map[Id]Position{}

func initPositions() {
	Positions[Main] = Position{Text: "", id: Main, layoutId: None, f: showMain}
	Positions[Sale] = Position{Text: "Sprzedaż", id: Sale, layoutId: Main, f: mock.Sale}
	Positions[ReadReports] = Position{Text: "Raporty czytające", id: ReadReports, layoutId: None, f: showReadReports}
	Positions[Programming] = Position{Text: "Programowanie", id: Programming, layoutId: None, f: showProgramming}
	Positions[Configuration] = Position{Text: "Konfiguracja", id: Configuration, layoutId: None, f: showConfiguration}

	Positions[DailyReport] = Position{Text: "Raport dobowy", id: DailyReport, layoutId: ReadReports, f: mock.DailyReport}
	Positions[PeriodicReport] = Position{Text: "Raport okresowy", id: PeriodicReport, layoutId: ReadReports, f: mock.PeriodicReport}
	Positions[WaresReport] = Position{Text: "Raport towarów", id: WaresReport, layoutId: ReadReports, f: reports.PrintWaresFull}

	Positions[ConfigLanguage] = Position{Text: "Język", id: ConfigLanguage, layoutId: Configuration, f: configuration.SetMenuText}
}
