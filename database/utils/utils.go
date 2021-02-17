package utils

import (
	"fmt"
)

type Price uint32

func (price Price) String() string {
	return fmt.Sprintf("%d.%02d", price/100, price%100)
}
