package manager

import (
	"github.com/gdamore/tcell"
	runewidth "github.com/mattn/go-runewidth"

	"github.com/wasanx25/goss/drawer"
	"github.com/wasanx25/goss/window"
)

type Manager struct {
	Window *window.Window
	Tui    tcell.Screen
	Drawer *drawer.Drawer
}

func New(w *window.Window, tui tcell.Screen, d *drawer.Drawer) *Manager {
	manager := &Manager{
		Window: w,
		Tui:    tui,
		Drawer: d,
	}

	return manager
}

func (m *Manager) Write() {
	x, y := 1, 1
	for _, s := range m.Drawer.Body {
		m.Tui.SetContent(x, y, s, nil, tcell.StyleDefault)
		switch s {
		case drawer.TAB:
			x += 4
		case drawer.NEW_LINE:
			x = 1
			y++
		case drawer.SPACE:
			x++
		}
		if int(m.Window.Row)-10 < y {
			break
		}
		x += runewidth.RuneWidth(s)
	}
}
