package texts

import (
	"okane/configuration"
	"okane/menu"
)

type positionNames []string

var currentNames *positionNames

var positionNamesPL = positionNames{
	menu.Sale:          "Sprzedaż",
	menu.ReadReports:   "Raporty czytające",
	menu.Programming:   "Programowanie",
	menu.Erasing:       "Kasowanie",
	menu.Configuration: "Konfiguracja",
	menu.ServiceMode:   "Serwis",
	menu.Log:           "Logowanie",
	menu.Communication: "Komunikacja"}

var positionNamesEN = positionNames{
	menu.Sale:          "Sale",
	menu.ReadReports:   "Reports",
	menu.Programming:   "Programming",
	menu.Erasing:       "Erasing",
	menu.Log:           "Logging",
	menu.Communication: "Communication"}

var positionNamesJP = positionNames{
	menu.Sale:          "セール",
	menu.ReadReports:   "レポート",
	menu.Programming:   "プログラミング",
	menu.Erasing:       "消去",
	menu.Log:           "ログイン",
	menu.Communication: "通信"}

// var saleMenuTextPL = SaleMenuType{
// 	"Sprzedaż",
// 	"Wpłata kasjera",
// 	"Wypłata kasjera"}

func Init() {
	currentNames = &positionNamesAll[configuration.System.CurrentLanguage]
}

var positionNamesAll = []positionNames{
	positionNamesPL,
	positionNamesEN,
	positionNamesJP}

// func getNames() positionNames {
// 	return positionNamesAll[CurrentLang]
// }

// func GetSaleMenuTexts() SaleMenuType {
// 	return saleMenuTexts[0]
// }
