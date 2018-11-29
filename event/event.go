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
	PageUpScreen
	PageDownScreen
	PageUpHalf
	PageDownHalf
	PageTop
	PageEnd
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
		case tcell.KeyCtrlB:
			e.DrawCh <- PageUpScreen
		case tcell.KeyCtrlF:
			e.DrawCh <- PageDownScreen
		case tcell.KeyCtrlU:
			e.DrawCh <- PageUpHalf
		case tcell.KeyCtrlD:
			e.DrawCh <- PageDownHalf
		}

		switch ev.Rune() {
		case 'j':
			e.DrawCh <- PageUp
		case 'k':
			e.DrawCh <- PageDown
		case 'q':
			e.DoneCh <- struct{}{}
		case 'G':
			e.DrawCh <- PageEnd
		}
	}
}
