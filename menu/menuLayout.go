package menu

import "fmt"

type Layout struct {
	PreviousMenuIndex Id
	positions         []Position
}

func (l *Layout) getPositions() []string {
	pos := []string{}
	for _, v := range l.positions {
		pos = append(pos, v.Text)
	}
	return pos
}

func (l Layout) Show() {
	positions := l.getPositions()
	buf := "=====================\n"
	i := 1
	for _, elem := range positions {
		buf += fmt.Sprintf("%d. %s\n", i, elem)
		i++
	}
	buf += "=====================\n"
	fmt.Println(buf)
}

var MenuTree = map[Id]Layout{}
