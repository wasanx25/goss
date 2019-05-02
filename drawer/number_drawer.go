package drawer

import (
	"strconv"

	"github.com/gdamore/tcell"
)

type numberDrawer struct {
	maxLine  int
	offset   int
	position Positioner
}

func NewNumberDrawer(maxLine, offset int, positioner Positioner) Drawer {
	return &numberDrawer{
		maxLine:  maxLine,
		offset:   offset,
		position: positioner,
	}
}

func (n *numberDrawer) SetOffset(offset int) {
	n.offset = offset
}

func (n *numberDrawer) SetLimitHeight(limitHeight int) {}

func (n *numberDrawer) Write(tui tcell.Screen, style tcell.Style) error {
	offsetInt := n.offset
	max := n.maxLine

	_, height := tui.Size()

	for i := 1; i <= height+1; i++ {
		if offsetInt > max+1 {
			break
		}

		offsetStr := strconv.Itoa(offsetInt)
		for _, r := range offsetStr {
			x, y := n.position.XAndY()
			tui.SetContent(x, y-1, r, nil, style)
			n.position.Add(r)
		}
		n.position.Break()
		offsetInt++
	}
	n.position.Reset()

	return nil
}
