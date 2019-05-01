package drawer

import "github.com/gdamore/tcell"

type Drawer interface {
	SetOffset(offset int)
	SetLimitHeight(limitHeight int)
	Write(tui tcell.Screen, style tcell.Style) error
}
