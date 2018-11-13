package run

import (
	"github.com/gdamore/tcell"
	"github.com/wasanx25/goss/manager"
)

const (
	TAB      = '\t'
	NEW_LINE = '\n'
	SPACE    = ' '
)

func Exec(content string) (err error) {
	m, err := manager.New(content)
	if err != nil {
		return
	}

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
		}
	}

	m.Tui.Fini()

	return
}
