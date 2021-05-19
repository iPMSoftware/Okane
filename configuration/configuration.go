package configuration

import (
	"fmt"
	"okane/hardware"
	"os"
	"os/exec"
)

type Language uint8

const (
	PL Language = iota
	ENG
	JP
	LAST_LANG
)

type SystemConfiguration struct {
	CurrentLanguage Language
}

var System SystemConfiguration

func clearScren() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func SetMenuText() {
	options := []string{
		"PL",
		"ENG",
		"JP",
	}

	new_val := Language(selectOption(&options, int(System.CurrentLanguage)))
	fmt.Println(options[new_val])
	System.CurrentLanguage = new_val
}

func selectOption(o *[]string, currentPos int) int {

	optionsDoubled := append(*o, *o...)
	optionsTripled := append(optionsDoubled, *o...)

	for true {
		clearScren()
		visibleOptions := optionsTripled[len(*o)+currentPos-1 : len(*o)+currentPos+2]
		for i, v := range visibleOptions {
			if i == 1 {
				fmt.Printf("-> %s\n", v)
			} else {
				fmt.Printf("   %s\n", v)
			}
		}
		key, _ := hardware.GetKey()
		if key == hardware.KeyArrowDown {
			currentPos = (currentPos + 1) % len(*o)
		}
		if key == hardware.KeyArrowUp {
			currentPos--
			if currentPos < 0 {
				currentPos = len(*o) - 1
			}
		}
		if key == hardware.KeyOk {
			break
		}
	}
	return currentPos
}
