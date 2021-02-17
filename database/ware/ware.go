package ware

import (
	"fmt"
	"okane/database/utils"
)

type index_t uint
type name_t string
type Vat_t byte

type Ware struct {
	id    index_t
	Name  name_t
	Vat   Vat_t
	Price utils.Price
}

func ToString(ware Ware) string {
	buf := ""
	buf += fmt.Sprintf("Numer: %35d\n", ware.id)
	buf += fmt.Sprintf("Nazwa: %35s\n", ware.Name)
	buf += fmt.Sprintf("Vat:   %35c\n", ware.Vat)
	buf += fmt.Sprintf("Cena:  %35s\n", ware.Price.String())
	return buf
}
