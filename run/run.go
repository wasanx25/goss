package run

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/wasanx25/goss/drawer"
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
loop:
	for {
		switch ev := m.Tui.PollEvent().(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyEnter:
				m.Tui.Clear()
				m.Tui.Show()
				break loop
			case tcell.KeyCtrlK:
				break loop
			}
			switch ev.Rune() {
			case 'j':
				m.Drawer.Increment()
				m.Tui.Clear()
				break loop
			}
		}
	}

	m.Tui.Fini()

	return
}
