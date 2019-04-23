package viewer

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gdamore/tcell"

	"github.com/wasanx25/goss/drawer"
	"github.com/wasanx25/goss/event"
)

type Viewer struct {
	tui          tcell.Screen
	drawer       *drawer.Drawer
	screenStyle  tcell.Style
	lineNumStyle tcell.Style
	contentStyle tcell.Style
	event        *event.Event
}

func New(text string) *Viewer {
	max := strings.Count(text, "\n")
	maxStr := strconv.Itoa(max)
	rowMax := len(maxStr) + 4 // line number default space

	viewer := &Viewer{
		drawer: drawer.New(text, 0, max, rowMax),
		event:  event.New(),
	}

	return viewer
}

func (v *Viewer) Init() error {
	tui, err := tcell.NewScreen()
	if err != nil {
		err = fmt.Errorf("tcell.NewScreen() error: %s", err)
		return err
	}

	err = tui.Init()
	if err != nil {
		err = fmt.Errorf("tcell.tui.Init() error: %s", err)
		return err
	}

	v.setTui(tui)
	v.screenStyle = tcell.StyleDefault.
		Foreground(tcell.ColorBlueViolet).
		Background(tcell.ColorBlack)
	v.lineNumStyle = tcell.StyleDefault.Foreground(tcell.Color59)
	v.contentStyle = tcell.StyleDefault.
		Foreground(tcell.ColorGray).
		Background(tcell.ColorBlack)

	v.tui.SetStyle(v.screenStyle)
	return nil
}

func (v *Viewer) Run() (err error) {
	v.write()

	v.tui.Show()
	v.setLimit()

	go func() {
		for {
			v.event.Action(v.tui)
		}
	}()

	go func() {
		for {
			select {
			case t := <-v.event.DrawCh:
				v.drawer.AddOffset(t)
				v.rewrite()
			case <-v.event.ResizeCh:
				_, height := v.tui.Size()
				v.drawer.SetLimit(height)
				v.rewrite()
			}
		}
	}()
	<-v.event.DoneCh

	v.tui.Fini()
	return
}

func (v *Viewer) setTui(tui tcell.Screen) {
	v.tui = tui
}

func (v *Viewer) setLimit() {
	_, height := v.tui.Size()
	v.drawer.SetLimit(height)
}

func (v *Viewer) rewrite() {
	v.tui.Clear()
	v.write()
	v.tui.Show()
}

func (v *Viewer) write() {
	v.drawer.Reset()
	str, _ := v.drawer.GetContent()
	width, height := v.tui.Size()

	v.writeLineNumber(height)

	v.drawer.Reset()
	for _, s := range str {
		col, row := v.drawer.Position()
		if col >= width {
			v.drawer.Break()
		}
		v.tui.SetContent(col, row, s, nil, v.contentStyle)
		v.drawer.AddPosition(s)
		if height < row {
			break
		}
	}
}

func (v *Viewer) writeLineNumber(height int) {
	offsetInt := v.drawer.Offset()

	for i := 1; i <= height+1; i++ {
		offsetStr := strconv.Itoa(offsetInt)
		for _, r := range offsetStr {
			col, row := v.drawer.Position()
			v.tui.SetContent(col, row-1, r, nil, v.lineNumStyle)
			v.drawer.AddPosition(r)
		}
		v.drawer.Break()
		offsetInt++
	}
}
