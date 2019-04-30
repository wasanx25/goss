package goss

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/wasanx25/goss/viewer"
)

func Run(text string) error {
	tui, err := tcell.NewScreen()

	if err != nil {
		return fmt.Errorf("tcell.NewScreen() error: %s", err)
	}

	if err = tui.Init(); err != nil {
		return fmt.Errorf("tcell.tui.Init() error: %s", err)
	}

	v := viewer.New(text, tui)

	if err := v.Open(); err != nil {
		return err
	}

	return nil
}
