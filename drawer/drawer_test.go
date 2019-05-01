package drawer_test

import (
	"github.com/gdamore/tcell"
)

func getString(tui tcell.SimulationScreen) string {
	width, _ := tui.Size()
	cells, _, _ := tui.GetContents()

	var runes []rune
	for i, c := range cells {
		runes = append(runes, c.Runes...)
		if (i+1)%width == 0 {
			runes = append(runes, '\n')
		}
	}

	return string(runes)
}
