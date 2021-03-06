package hardware

import (
	"github.com/eiannone/keyboard"
)

type Key int

const (
	KeyNone Key = iota
	Key1
	Key2
	Key3
	Key4
	Key5
	Key6
	Key7
	Key8
	Key9
	KeyC
	KeyArrowDown
	KeyArrowUp
	KeyOk
)

func initKeyboard() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}

}

func deinitKeyboard() {
	_ = keyboard.Close()
}

func translateRune(r rune) (Key, bool) {
	switch r {
	case '1':
		return Key1, true
	case '2':
		return Key2, true
	case '3':
		return Key3, true
	case '4':
		return Key4, true
	case '5':
		return Key5, true
	case '6':
		return Key6, true
	case '7':
		return Key7, true
	case '8':
		return Key8, true
	case '9':
		return Key9, true
	case 'c':
		return KeyC, true
	default:
		return KeyNone, false
	}
}

func KeyToValue(key Key) int {
	switch key {
	case Key1:
		return 1
	case Key2:
		return 2
	case Key3:
		return 3
	case Key4:
		return 4
	case Key5:
		return 5
	case Key6:
		return 6
	case Key7:
		return 7
	case Key8:
		return 8
	case Key9:
		return 9
	}
	return 0
}

func GetKey() (Key, bool) {
	for true {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyEsc {
			return KeyC, false
		}
		if key == keyboard.KeyArrowDown {
			return KeyArrowDown, true
		}
		if key == keyboard.KeyArrowUp {
			return KeyArrowUp, true
		}
		if key == keyboard.KeyEnter {
			return KeyOk, true
		}
		return translateRune(char)
	}
	return KeyNone, false
}
