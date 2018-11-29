package event_test

import (
	"testing"

	"github.com/gdamore/tcell"
	"github.com/wasanx25/goss/event"
)

func TestAction(t *testing.T) {
	tui := tcell.NewSimulationScreen("")

	drawCh := make(chan event.Type)
	doneCh := make(chan struct{})

	e := event.New(drawCh, doneCh)

	if err := tui.Init(); err != nil {
		t.Fatal(err)
	}

	tui.SetSize(90, 20)
	go func() {
		for {
			e.Action(tui)
		}
	}()

	tui.InjectKey(tcell.KeyRune, 'j', tcell.ModNone)
	tui.InjectKey(tcell.KeyRune, 'k', tcell.ModNone)
	tui.InjectKey(tcell.KeyRune, 'q', tcell.ModNone)
	draw := <-e.DrawCh
	if draw != event.PageUp {
		t.Errorf("expected=%v, got=%v", event.PageUp, draw)
	}

	draw = <-e.DrawCh
	if draw != event.PageDown {
		t.Errorf("expected=%v, got=%v", event.PageDown, draw)
	}

	done := <-e.DoneCh
	if done != struct{}{} {
		t.Errorf("expected=%v, got=%v", struct{}{}, done)
	}
}
