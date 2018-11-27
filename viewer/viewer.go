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
}

func New(body string) *Viewer {
	w := window.New()

	manager := &Viewer{
		Window: w,
		Drawer: drawer.New(body, 0),
	}

	return manager
}

func (v *Viewer) Init() error {
	v.Window.SetSize()
	v.Drawer.Limit = uint(v.Window.Row)

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
	v.Drawer.PositionInit()
	str, _ := v.Drawer.Get()
	for _, s := range str {
		v.Tui.SetContent(v.Drawer.Position.Col, v.Drawer.Position.Row, s, nil, tcell.StyleDefault)
		v.Drawer.AddPosition(s)
		if int(v.Window.Row) < v.Drawer.Position.Row {
			break
		}
	}
}
