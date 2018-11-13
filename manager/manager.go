package manager

import (
	"fmt"

	"github.com/gdamore/tcell"

	"github.com/wasanx25/goss/drawer"
	"github.com/wasanx25/goss/window"
)

type Manager struct {
	Window *window.Window
	Tui    tcell.Screen
	Drawer *drawer.Drawer
}

func New(body string) (*Manager, error) {
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

	drawer := drawer.New(body, 0)

	manager := &Manager{
		Window: window,
		Tui:    tui,
		Drawer: drawer,
	}

	return manager, nil
}
