package event

import (
	"github.com/gdamore/tcell"
)

type Type int

const (
	OneIncrement Type = iota
	OneDecrement
)

func Action(ev tcell.Event, drawCh chan Type, doneCh chan struct{}) {
	switch e := ev.(type) {
	case *tcell.EventKey:
		switch e.Key() {
		case tcell.KeyEscape:
			doneCh <- struct{}{}
		}

		switch e.Rune() {
		case 'j':
			drawCh <- OneIncrement
		case 'k':
			drawCh <- OneDecrement
		}
	}
}
