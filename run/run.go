package run

import (
	"github.com/gdamore/tcell"
	runewidth "github.com/mattn/go-runewidth"
	"github.com/wasanx25/goss/manager"
)

const (
	TAB      = '\t'
	NEW_LINE = '\n'
	SPACE    = ' '
)

func Exec(content string) (err error) {
	m, err := manager.New(content)
	if err != nil {
		return
	}

	m.Tui.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorBlueViolet).Background(tcell.ColorBlack))
	x, y := 1, 1
	for _, s := range content {
		m.Tui.SetContent(x, y, s, nil, tcell.StyleDefault)
		switch s {
		case TAB:
			x += 4
		case NEW_LINE:
			x = 1
			y++
		case SPACE:
			x++
		}
		if int(m.Window.Row)-10 < y {
			m.Tui.SetContent(x, y+1, 'y', nil, tcell.StyleDefault)
			break
		}
		x += runewidth.RuneWidth(s)
	}

	m.Tui.Show()
loop:
	for {
		switch ev := m.Tui.PollEvent().(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyEnter:
				m.Tui.Clear()
				m.Tui.Show()
				break loop
			case tcell.KeyCtrlK:
				break loop
			}
		}
	}

	m.Tui.Fini()

	return
}

func write() {
	// r := bufio.NewReader(strings.NewReader(content))
	// for {
	// 	line, err := r.ReadString('\n')
	// 	if err == io.EOF {
	// 		break
	// 	} else if err != nil {
	// 		fmt.Println(err)
	// 		break
	// 	}
	// 	fmt.Print(line)
	// }
}
