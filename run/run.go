package run

import (
	"fmt"

	"github.com/gdamore/tcell"
	runewidth "github.com/mattn/go-runewidth"
	"github.com/wasanx25/goss/window"
)

const (
	TAB      = '\t'
	NEW_LINE = '\n'
	SPACE    = ' '
)

func Exec(content string) (err error) {
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

	w, err := window.New()

	tui.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorBlueViolet).Background(tcell.ColorBlack))
	x := 1
	y := 1
	for _, s := range content {
		tui.SetContent(x, y, s, nil, tcell.StyleDefault)
		switch s {
		case TAB:
			x += 4
		case NEW_LINE:
			x = 1
			y++
		case SPACE:
			x++
		}
		if int(w.Row)-10 < y {
			tui.SetContent(x, y+1, 'y', nil, tcell.StyleDefault)
			break
		}
		x += runewidth.RuneWidth(s)
	}

	tui.Show()
loop:
	for {
		switch ev := tui.PollEvent().(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyEnter:
				break loop
			case tcell.KeyCtrlK:
				break loop
			}
		}
	}

	tui.Fini()

	return
}
