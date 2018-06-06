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

	data := `
# 見出し1
## 見出し2
### 見出し3
#### 見出し4


##### 見出し5
`

	err = tui.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "tui.Init() error: %s", err)
	}

	tui.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorBlueViolet).Background(tcell.ColorBlack))
	x := 1
	y := 1
	for _, s := range data {
		tui.SetContent(x, y, s, nil, tcell.StyleDefault)
		if string(s) == "\n" {
			x = 1
			y++
		}
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
