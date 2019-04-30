package goss

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/wasanx25/goss/viewer"
)

type StyleOptions func(*viewer.Styles)

func ScreenStyle(style tcell.Style) StyleOptions {
	return func(s *viewer.Styles) {
		s.SetScreenStyle(style)
	}
}

func LineNumStyle(style tcell.Style) StyleOptions {
	return func(s *viewer.Styles) {
		s.SetLineNumStyle(style)
	}
}

func ContentStyle(style tcell.Style) StyleOptions {
	return func(s *viewer.Styles) {
		s.SetContentStyle(style)
	}
}

func Run(text string, styleOptions ...StyleOptions) error {
	tui, err := tcell.NewScreen()

	if err != nil {
		return fmt.Errorf("tcell.NewScreen() error: %s", err)
	}

	if err = tui.Init(); err != nil {
		return fmt.Errorf("tcell.tui.Init() error: %s", err)
	}

	styles := &viewer.Styles{}
	// Default style
	styles.SetScreenStyle(tcell.StyleDefault.Foreground(tcell.ColorBlueViolet))
	styles.SetLineNumStyle(tcell.StyleDefault.Foreground(tcell.Color59))
	styles.SetContentStyle(tcell.StyleDefault.Foreground(tcell.ColorGray).Background(tcell.ColorBlack))

	for _, option := range styleOptions {
		option(styles)
	}

	v := viewer.New(text, tui, styles)

	if err := v.Open(); err != nil {
		return err
	}

	return nil
}
