package manager

import (
	"fmt"

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

func New(body string) (*Manager, error) {
	tui, err := tcell.NewScreen()
	if err != nil {
		err = fmt.Errorf("tcell.NewScreen() error: %s", err)
		return nil, err
	}

	err = tui.Init()
	if err != nil {
		err = fmt.Errorf("tcell.tui.Init() error: %s", err)
		return nil, err
	}

	window, err := window.New()
	if err != nil {
		err = fmt.Errorf("window.New() error: %s", err)
		return nil, err
	}

	drawer := drawer.New(body, 0)

	manager := &Manager{
		Window: window,
		Tui:    tui,
		Drawer: drawer,
	}

	return manager, nil
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
