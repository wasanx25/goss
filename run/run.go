package run

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/wasanx25/goss/drawer"
	"github.com/wasanx25/goss/event"
	"github.com/wasanx25/goss/manager"
	"github.com/wasanx25/goss/window"
)

const (
	TAB      = '\t'
	NEW_LINE = '\n'
	SPACE    = ' '
)

func Exec(body string) (err error) {
	w, err := window.New()
	if err != nil {
		err = fmt.Errorf("window.New() error: %s", err)
		return
	}

	tui, err := tcell.NewScreen()
	if err != nil {
		err = fmt.Errorf("tcell.NewScreen() error: %s", err)
		return
	}

	err = tui.Init()
	if err != nil {
		err = fmt.Errorf("tcell.tui.Init() error: %s", err)
		return
	}

	m := manager.New(w, tui, drawer.New(body, 0))

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
	return
}
