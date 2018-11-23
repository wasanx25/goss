package manager

import (
	"github.com/gdamore/tcell"
	runewidth "github.com/mattn/go-runewidth"

	"github.com/wasanx25/goss/drawer"
	"github.com/wasanx25/goss/event"
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

func (m *Manager) Start() {
	m.Tui.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorBlueViolet).Background(tcell.ColorBlack))
	m.Write()

	m.Tui.Show()
	drawCh := make(chan event.Type, 0)
	doneCh := make(chan struct{}, 0)
	go func() {
		for {
			event.Action(m.Tui.PollEvent(), drawCh, doneCh)
		}
	}()

	go func() {
		for {
			select {
			case <-doneCh:
			case t := <-drawCh:
				switch t {
				case event.OneDecrement:
					m.Drawer.Decrement()
					m.Rewrite()
				case event.OneIncrement:
					m.Drawer.Increment()
					m.Rewrite()
				}
			}
		}
	}()
	<-doneCh

	m.Tui.Fini()
}

func (m *Manager) Rewrite() {
	m.Tui.Clear()
	m.Write()
	m.Tui.Show()
}

func (m *Manager) Write() {
	x, y := 1, 0
	str, _ := m.Drawer.Get(uint(m.Window.Row))
	for _, s := range str {
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
		if int(m.Window.Row) < y {
			break
		}
		x += runewidth.RuneWidth(s)
	}
}
