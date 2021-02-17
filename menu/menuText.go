package menu

const (
	PL = iota
	EN
	JP
	LAST_LANG
)

const mainMenuNumberOfPositions = 6
const saleMenuNumberOfPositions = 3
const reportsMenuNumberOfPositions = 4

type MainMenuType [mainMenuNumberOfPositions]string
type SaleMenuType [saleMenuNumberOfPositions]string

var CurrentLang = PL

var mainMenuTextPL = MainMenuType{
	"Sprzedaż",
	"Raporty",
	"Programowanie",
	"Kasowanie",
	"Logowanie",
	"Komunikacja"}

var mainMenuTextEN = MainMenuType{
	"Sale",
	"Reports",
	"Programming",
	"Erasing",
	"Logging",
	"Communication"}

var mainMenuTextJP = MainMenuType{
	"セール",
	"レポート",
	"プログラミング",
	"消去",
	"ログイン",
	"通信"}

var saleMenuTextPL = SaleMenuType{
	"Sprzedaż",
	"Wpłata kasjera",
	"Wypłata kasjera"}

var mainMenuTexts = []MainMenuType{
	mainMenuTextPL,
	mainMenuTextEN,
	mainMenuTextJP}

var saleMenuTexts = []SaleMenuType{
	saleMenuTextPL}

func GetMainMenuTexts() MainMenuType {
	return mainMenuTexts[CurrentLang]
}

func GetSaleMenuTexts() SaleMenuType {
	return saleMenuTexts[0]
}
