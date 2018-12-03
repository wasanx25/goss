package viewer

import (
	"fmt"

	"github.com/gdamore/tcell"

	"github.com/wasanx25/goss/drawer"
	"github.com/wasanx25/goss/event"
	"github.com/wasanx25/goss/window"
)

type Viewer struct {
	Window *window.Window
	Tui    tcell.Screen
	Drawer *drawer.Drawer
	Color  tcell.Style
	Event  *event.Event
}

func New(text string) *Viewer {
	manager := &Viewer{
		Window: window.New(),
		Drawer: drawer.New(text, 0),
		Event:  event.New(),
	}

	return manager
}

func (v *Viewer) Init() error {
	v.Window.SetSize()
	v.Drawer.Limit = int(v.Window.Row)

	tui, err := tcell.NewScreen()
	if err != nil {
		err = fmt.Errorf("tcell.NewScreen() error: %s", err)
		return err
	}

	err = tui.Init()
	if err != nil {
		err = fmt.Errorf("tcell.tui.Init() error: %s", err)
		return err
	}

	v.Tui = tui
	v.Color = tcell.StyleDefault.
		Foreground(tcell.ColorBlueViolet).
		Background(tcell.ColorBlack)
	v.Tui.SetStyle(v.Color)
	return nil
}

func (v *Viewer) Start() {
	v.write()

	v.Tui.Show()
	go func() {
		for {
			v.Event.Action(v.Tui)
		}
	}()

	go func() {
		for {
			select {
			case t := <-v.Event.DrawCh:
				switch t {
				case event.PageDown:
					v.Drawer.Increment()
					v.rewrite()
				case event.PageUp:
					v.Drawer.Decrement()
					v.rewrite()
				case event.PageDownHalf:
					v.Drawer.IncrementHalf()
					v.rewrite()
				case event.PageUpHalf:
					v.Drawer.DecrementHalf()
					v.rewrite()
				case event.PageDownScreen:
					v.Drawer.IncrementWindow()
					v.rewrite()
				case event.PageUpScreen:
					v.Drawer.DecrementWindow()
					v.rewrite()
				}
			case <-v.Event.ResizeCh:
				v.Window.SetSize()
				v.Drawer.Limit = int(v.Window.Row)
				v.rewrite()
			}
		}
	}()
	<-v.Event.DoneCh

	v.Tui.Fini()
}

func (v *Viewer) rewrite() {
	v.Tui.Clear()
	v.write()
	v.Tui.Show()
}

func (v *Viewer) write() {
	v.Drawer.InitPosition()
	str, _ := v.Drawer.Get()
	for _, s := range str {
		v.Tui.SetContent(v.Drawer.Position.Col, v.Drawer.Position.Row, s, nil, tcell.StyleDefault)
		v.Drawer.AddPosition(s)
		if int(v.Window.Row) < v.Drawer.Position.Row {
			break
		}
	}
}
