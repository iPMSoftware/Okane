package reports

import (
	"fmt"
	"okane/database/ware"
	"okane/mock"
)

func PrintWaresFull() {
	fmt.Println(mock.PrintHeader())
	fmt.Println(ware.GetBase().ToString())
	fmt.Println(mock.PrintFooter())
	mock.WaitForKey()
}
