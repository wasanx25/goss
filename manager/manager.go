package manager

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/wasanx25/goss/window"
)

type Manager struct {
	Window *window.Window
	Tui    tcell.Screen
}

func New() (*Manager, error) {
	tui, err := tcell.NewScreen()
	if err != nil {
		err = fmt.Errorf("tcell.NewScreen() error: %s", err)
		return nil, err
	}

	err = tui.Init()
	if err != nil {
		err = fmt.Errorf("tcell.tui.Init() error: %s", err)
		return nil, err
	}

	window, err := window.New()
	if err != nil {
		err = fmt.Errorf("window.New() error: %s", err)
		return nil, err
	}

	manager := &Manager{
		Window: window,
		Tui:    tui,
	}

	return manager, nil
}
