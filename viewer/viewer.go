package viewer

import (
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

func New(text string, tui tcell.Screen) *Viewer {
	max := strings.Count(text, "\n")
	maxStr := strconv.Itoa(max)
	rowMax := len(maxStr) + 4 // line number default space

	viewer := &Viewer{
		tui:    tui,
		drawer: drawer.New(text, 0, max, rowMax),
		event:  event.New(),
	}

	return viewer
}

func (v *Viewer) Run() (err error) {
	v.screenStyle = tcell.StyleDefault.
		Foreground(tcell.ColorBlueViolet).
		Background(tcell.ColorBlack)
	v.lineNumStyle = tcell.StyleDefault.Foreground(tcell.Color59)
	v.contentStyle = tcell.StyleDefault.
		Foreground(tcell.ColorGray).
		Background(tcell.ColorBlack)

	v.tui.SetStyle(v.screenStyle)

	v.drawer.Write(v.tui, v.contentStyle, v.lineNumStyle)

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
				v.setLimit()
				v.rewrite()
			}
		}
	}()
	<-v.event.DoneCh

	v.tui.Fini()
	return
}

func (v *Viewer) setLimit() {
	_, height := v.tui.Size()
	v.drawer.SetLimit(height)
}

func (v *Viewer) rewrite() {
	v.tui.Clear()
	v.drawer.Write(v.tui, v.contentStyle, v.lineNumStyle)
	v.tui.Show()
}
