package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
	runewidth "github.com/mattn/go-runewidth"
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

	tui.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorBlueViolet).Background(tcell.ColorBlack))
	x := 1
	for _, s := range "ジャ-ジャ-麺" {
		tui.SetContent(x, 1, s, nil, tcell.StyleDefault)
		x += runewidth.RuneWidth(s)
	}

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
