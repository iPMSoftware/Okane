package ware

import (
	"fmt"
	"okane/database/utils"
)

const Max = 2000

type index_t uint
type name_t string
type Vat_t byte

type Wares []Ware

var base Wares

type Ware struct {
	id    index_t
	Name  name_t
	Vat   Vat_t
	Price utils.Price
}

func (w *Ware) ToString() string {
	buf := ""
	buf += fmt.Sprintf("Numer: %35d\n", w.id)
	buf += fmt.Sprintf("Nazwa: %35s\n", w.Name)
	buf += fmt.Sprintf("Vat:   %35c\n", w.Vat)
	buf += fmt.Sprintf("Cena:  %35s\n", w.Price.String())
	return buf
}

func (w *Wares) empty() bool {
	return len(*w) == 0
}

func (w *Wares) ToString() string {
	var empty = w.empty()
	buf := "=====================\n"
	buf += "Raport Towarów\n"
	buf += "=====================\n"
	if empty {
		buf += "\n   Brak towarów\n"
	} else {
		for _, elem := range *w {
			if elem.id != 0 {
				buf += elem.ToString()
				buf += "=====================\n"
			}
		}
	}
	return buf
}

func Init() {
	base = Wares{}
}
func GetBase() *Wares {
	return &base
}
