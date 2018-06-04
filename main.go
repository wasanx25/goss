package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
)

func main() {
	var err error

	tui, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "tcell.NewScreen() error: %s", err)
	}

	err = tui.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "tui.Init() error: %s", err)
	}

	tui.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack))

	tui.Show()

loop:
	for {
		switch ev := tui.PollEvent().(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyEnter:
				fmt.Println("Exit!")
				break loop
			}
		}
	}

	tui.Fini()
}
