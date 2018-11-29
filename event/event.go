package event

import (
	"github.com/gdamore/tcell"
)

type Type int

type Event struct {
	DrawCh chan Type
	DoneCh chan struct{}
}

const (
	PageUp Type = iota
	PageDown
)

func New(drawCh chan Type, doneCh chan struct{}) *Event {
	return &Event{
		DrawCh: drawCh,
		DoneCh: doneCh,
	}
}

func (e *Event) Action(tui tcell.Screen) {
	switch ev := tui.PollEvent().(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyEscape:
			e.DoneCh <- struct{}{}
		}

		switch ev.Rune() {
		case 'j':
			e.DrawCh <- PageUp
		case 'k':
			e.DrawCh <- PageDown
		case 'q':
			e.DoneCh <- struct{}{}
		}
	}
}
