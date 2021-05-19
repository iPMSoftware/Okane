package mock

import (
	"fmt"
	"okane/hardware"
)

func Sale() {
	fmt.Println("Sale Function")
	WaitForKey()
}

func DailyReport() {
	fmt.Println("Daily report")
	WaitForKey()
}

func PeriodicReport() {
	fmt.Println("Periodic report")
	WaitForKey()
}

func PrintHeader() string {
	var buf string = "                     \n" +
		"      iPM Okane      \n" +
		"       GoVer 1       \n"

	return buf
}

func PrintFooter() string {
	var buf string = "=====================\n" +
		"#1   01-01-1970 00:01\n"
	return buf
}

func WaitForKey() {
	fmt.Println("Press C to exit function")
	for true {
		key, _ := hardware.GetKey()
		if key == hardware.KeyC {
			break
		}
	}
}
