package goss

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/wasanx25/goss/drawer"
	"github.com/wasanx25/goss/viewer"
	"github.com/wasanx25/goss/window"
)

func Run(body string) (err error) {
	w := window.New()

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

	v := viewer.New(w, tui, drawer.New(body, 0))
	v.Start()

	return
}
