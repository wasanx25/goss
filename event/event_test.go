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

	if err := tui.Init(); err != nil {
		t.Fatal(err)
	}

	tui.SetSize(90, 20)
	go func() {
		for {
			event.Action(tui, drawCh, doneCh)
		}
	}()

	tui.InjectKey(tcell.KeyRune, 'j', tcell.ModNone)
	tui.InjectKey(tcell.KeyRune, 'k', tcell.ModNone)
	d := <-drawCh
	if d != event.OneIncrement {
		t.Errorf("expected=%v, got=%v", event.OneIncrement, d)
	}

	d = <-drawCh
	if d != event.OneDecrement {
		t.Errorf("expected=%v, got=%v", event.OneDecrement, d)
	}
}
