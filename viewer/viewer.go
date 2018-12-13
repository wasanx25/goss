package viewer

import (
	"fmt"

	"github.com/gdamore/tcell"

	"github.com/wasanx25/goss/drawer"
	"github.com/wasanx25/goss/event"
)

type Viewer struct {
	tui    tcell.Screen
	drawer *drawer.Drawer
	color  tcell.Style
	event  *event.Event
}

func New(text string) *Viewer {
	manager := &Viewer{
		drawer: drawer.New(text, 0),
		event:  event.New(),
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

	v.tui = tui
	v.color = tcell.StyleDefault.
		Foreground(tcell.ColorBlueViolet).
		Background(tcell.ColorBlack)
	v.tui.SetStyle(v.color)
	return nil
}

func (v *Viewer) Run() (err error) {
	v.write()

	v.tui.Show()
	_, height := v.tui.Size()
	v.drawer.SetLimit(height)

	go func() {
		for {
			v.event.Action(v.tui)
		}
	}()

	go func() {
		for {
			select {
			case t := <-v.event.DrawCh:
				v.drawer.AddOffset(t)
				v.rewrite()
			case <-v.event.ResizeCh:
				_, height = v.tui.Size()
				v.drawer.SetLimit(height)
				v.rewrite()
			}
		}
	}()
	<-v.event.DoneCh

	v.tui.Fini()
	return
}

func (v *Viewer) rewrite() {
	v.tui.Clear()
	v.write()
	v.tui.Show()
}

func (v *Viewer) write() {
	v.drawer.InitPosition()
	str, _ := v.drawer.GetContent()
	width, height := v.tui.Size()
	for _, s := range str {
		if v.drawer.Position.Col >= width {
			v.drawer.Break()
		}
		v.tui.SetContent(v.drawer.Position.Col, v.drawer.Position.Row, s, nil, tcell.StyleDefault)
		v.drawer.AddPosition(s)
		if height < v.drawer.Position.Row {
			break
		}
	}
}
