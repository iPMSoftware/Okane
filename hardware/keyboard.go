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
	KeyC
)

func initKeyboard() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {

	}()

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
	default:
		return KeyNone, false
	}
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
		return translateRune(char)
	}
	return KeyNone, false
}
