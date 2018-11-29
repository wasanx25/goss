package event

import (
	"github.com/gdamore/tcell"
)

type Type int

const (
	PageUp Type = iota
	PageDown
)

func Action(tui tcell.Screen, drawCh chan Type, doneCh chan struct{}) {
	switch e := tui.PollEvent().(type) {
	case *tcell.EventKey:
		switch e.Key() {
		case tcell.KeyEscape:
			doneCh <- struct{}{}
		}

		switch e.Rune() {
		case 'j':
			drawCh <- PageUp
		case 'k':
			drawCh <- PageDown
		case 'q':
			doneCh <- struct{}{}
		}
	}
}
