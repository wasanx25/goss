package drawer

import (
	"strconv"

	"github.com/gdamore/tcell"
)

type lineNumberDrawer struct {
	maxLine  int
	offset   int
	position Positioner
}

func NewLineNumberDrawer(maxLine, offset int, positioner Positioner) Drawer {
	return &lineNumberDrawer{
		maxLine:  maxLine,
		offset:   offset,
		position: positioner,
	}
}

func (l *lineNumberDrawer) SetOffset(offset int) {
	l.offset = offset
}

func (l *lineNumberDrawer) SetLimitHeight(limitHeight int) {}

func (l *lineNumberDrawer) Write(tui tcell.Screen, style tcell.Style) error {
	offsetInt := l.offset
	max := l.maxLine

	_, height := tui.Size()

	for i := 1; i <= height+1; i++ {
		if offsetInt > max+1 {
			break
		}

		offsetStr := strconv.Itoa(offsetInt)
		for _, r := range offsetStr {
			x, y := l.position.XAndY()
			tui.SetContent(x, y-1, r, nil, style)
			l.position.Add(r)
		}
		l.position.Break()
		offsetInt++
	}
	l.position.Reset()

	return nil
}
