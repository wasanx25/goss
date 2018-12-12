package viewer

import (
	"fmt"

	"github.com/gdamore/tcell"

	"github.com/wasanx25/goss/drawer"
	"github.com/wasanx25/goss/event"
)

type Viewer struct {
	Tui    tcell.Screen
	Drawer *drawer.Drawer
	Color  tcell.Style
	Event  *event.Event
}

func New(text string) *Viewer {
	manager := &Viewer{
		Drawer: drawer.New(text, 0),
		Event:  event.New(),
	}

	return manager
}

func (v *Viewer) Init() error {
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

func (v *Viewer) Start() (err error) {
	v.write()

	v.Tui.Show()
	_, height := v.Tui.Size()
	v.Drawer.SetLimit(height)

	go func() {
		for {
			v.Event.Action(v.Tui)
		}
	}()

	go func() {
		for {
			select {
			case t := <-v.Event.DrawCh:
				v.Drawer.AddOffset(t)
				v.rewrite()
			case <-v.Event.ResizeCh:
				_, height = v.Tui.Size()
				v.Drawer.SetLimit(height)
				v.rewrite()
			}
		}
	}()
	<-v.Event.DoneCh

	v.Tui.Fini()
	return
}

func (v *Viewer) rewrite() {
	v.Tui.Clear()
	v.write()
	v.Tui.Show()
}

func (v *Viewer) write() {
	v.Drawer.InitPosition()
	str, _ := v.Drawer.GetContent()
	width, height := v.Tui.Size()
	for _, s := range str {
		if v.Drawer.Position.Col >= width {
			v.Drawer.Break()
		}
		v.Tui.SetContent(v.Drawer.Position.Col, v.Drawer.Position.Row, s, nil, tcell.StyleDefault)
		v.Drawer.AddPosition(s)
		if height < v.Drawer.Position.Row {
			break
		}
	}
}
