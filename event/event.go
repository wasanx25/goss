package event

import (
	"github.com/gdamore/tcell"
)

type Type int

type Event struct {
	DrawCh   chan Type
	QuitCh   chan struct{}
	ResizeCh chan struct{}
}

const (
	PageUp Type = iota + 1
	PageDown
	PageUpScreen
	PageDownScreen
	PageUpHalf
	PageDownHalf
	PageTop
	PageEnd
)

func New() *Event {
	return &Event{
		DrawCh:   make(chan Type),
		QuitCh:   make(chan struct{}),
		ResizeCh: make(chan struct{}),
	}
}

func (e *Event) Action(tui tcell.Screen) {
	switch ev := tui.PollEvent().(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyEscape:
			e.QuitCh <- struct{}{}
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
		case 'k':
			e.DrawCh <- PageUp
		case 'j':
			e.DrawCh <- PageDown
		case 'q':
			e.QuitCh <- struct{}{}
		case 'g':
			e.DrawCh <- PageTop
		case 'G':
			e.DrawCh <- PageEnd
		}
	case *tcell.EventResize:
		e.ResizeCh <- struct{}{}
	}
}
