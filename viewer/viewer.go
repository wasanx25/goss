package viewer

import (
	"strconv"
	"strings"

	"github.com/gdamore/tcell"

	"github.com/wasanx25/goss/drawer"
	"github.com/wasanx25/goss/event"
)

type Viewer struct {
	tui    tcell.Screen
	drawer *drawer.Drawer
	event  *event.Event
	styles *Styles
}

type Styles struct {
	screenStyle  tcell.Style
	lineNumStyle tcell.Style
	contentStyle tcell.Style
}

func (s *Styles) SetScreenStyle(style tcell.Style)  { s.screenStyle = style }
func (s *Styles) SetLineNumStyle(style tcell.Style) { s.lineNumStyle = style }
func (s *Styles) SetContentStyle(style tcell.Style) { s.contentStyle = style }

func New(text string, tui tcell.Screen, styles *Styles) *Viewer {
	max := strings.Count(text, "\n")
	maxStr := strconv.Itoa(max)
	rowMax := len(maxStr) + 4 // line number default space

	viewer := &Viewer{
		tui:    tui,
		drawer: drawer.New(text, 0, max, rowMax),
		event:  event.New(),
		styles: styles,
	}

	return viewer
}

func (v *Viewer) Open() (err error) {
	v.tui.SetStyle(v.styles.screenStyle)

	v.drawer.Write(v.tui, v.styles.contentStyle, v.styles.lineNumStyle)

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
	v.drawer.Write(v.tui, v.styles.contentStyle, v.styles.lineNumStyle)
	v.tui.Show()
}
