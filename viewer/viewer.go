package viewer

import (
	"github.com/gdamore/tcell"
	runewidth "github.com/mattn/go-runewidth"

	"github.com/wasanx25/goss/drawer"
	"github.com/wasanx25/goss/event"
	"github.com/wasanx25/goss/window"
)

type Viewer struct {
	Window *window.Window
	Tui    tcell.Screen
	Drawer *drawer.Drawer
}

func New(w *window.Window, tui tcell.Screen, d *drawer.Drawer) *Viewer {
	_ = w.GetSize()
	manager := &Viewer{
		Window: w,
		Tui:    tui,
		Drawer: d,
	}

	return manager
}

func (v *Viewer) Start() {
	v.Tui.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorBlueViolet).Background(tcell.ColorBlack))
	v.Write()

	v.Tui.Show()
	drawCh := make(chan event.Type, 0)
	doneCh := make(chan struct{}, 0)
	go func() {
		for {
			event.Action(v.Tui, drawCh, doneCh)
		}
	}()

	go func() {
		for {
			select {
			case <-doneCh:
			case t := <-drawCh:
				switch t {
				case event.OneDecrement:
					v.Drawer.Decrement()
					v.Rewrite()
				case event.OneIncrement:
					v.Drawer.Increment()
					v.Rewrite()
				}
			}
		}
	}()
	<-doneCh

	v.Tui.Fini()
}

func (v *Viewer) Rewrite() {
	v.Tui.Clear()
	v.Write()
	v.Tui.Show()
}

func (v *Viewer) Write() {
	x, y := 1, 0
	str, _ := v.Drawer.Get(uint(v.Window.Row))
	for _, s := range str {
		v.Tui.SetContent(x, y, s, nil, tcell.StyleDefault)
		switch s {
		case drawer.TAB:
			x += 4
		case drawer.NEW_LINE:
			x = 1
			y++
		case drawer.SPACE:
			x++
		}
		if int(v.Window.Row) < y {
			break
		}
		x += runewidth.RuneWidth(s)
	}
}
