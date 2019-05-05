package viewer

import (
	"strconv"
	"strings"

	"github.com/gdamore/tcell"
	"golang.org/x/sync/errgroup"

	"github.com/wasanx25/goss/drawer"
	"github.com/wasanx25/goss/event"
	"github.com/wasanx25/goss/offsetter"
)

type Viewer struct {
	tui           tcell.Screen
	contentDrawer drawer.Drawer
	numberDrawer  drawer.Drawer
	offsetter     offsetter.Offsetter
	event         *event.Event
	styles        *Styles
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
	maxLine := strings.Count(text, "\n")
	maxStr := strconv.Itoa(maxLine)
	rowNumMax := len(maxStr) + 4 // line number default space
	defaultOffset := 0
	_, limitHeight := tui.Size()

	p1 := drawer.NewPositioner(rowNumMax)
	p2 := drawer.NewPositioner(rowNumMax)
	c := drawer.NewContentDrawer(text, defaultOffset, limitHeight, p1)
	n := drawer.NewNumberDrawer(maxLine, defaultOffset, p2)
	o := offsetter.NewOffsetter(defaultOffset, maxLine, limitHeight)

	viewer := &Viewer{
		tui:           tui,
		event:         event.New(),
		contentDrawer: c,
		numberDrawer:  n,
		offsetter:     o,
		styles:        styles,
	}

	return viewer
}

func (v *Viewer) Open() (err error) {
	v.tui.SetStyle(v.styles.screenStyle)

	eg := errgroup.Group{}

	eg.Go(func() error {
		return v.contentDrawer.Write(v.tui, v.styles.contentStyle)
	})

	eg.Go(func() error {
		return v.numberDrawer.Write(v.tui, v.styles.lineNumStyle)
	})

	if err = eg.Wait(); err != nil {
		return err
	}

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
				offset := v.offsetter.UpdateAndGet(t)
				v.numberDrawer.SetOffset(offset)
				v.contentDrawer.SetOffset(offset)
				if err = v.rewrite(); err != nil {
					close(v.event.DoneCh)
				}
			case <-v.event.ResizeCh:
				v.setLimit()
				if err = v.rewrite(); err != nil {
					close(v.event.DoneCh)
				}
			}
		}
	}()
	<-v.event.DoneCh

	v.tui.Fini()
	return
}

func (v *Viewer) setLimit() {
	_, limitHeight := v.tui.Size()
	v.contentDrawer.SetLimitHeight(limitHeight)
}

func (v *Viewer) rewrite() (err error) {
	v.tui.Clear()

	if err = v.contentDrawer.Write(v.tui, v.styles.contentStyle); err != nil {
		return
	}

	if err = v.numberDrawer.Write(v.tui, v.styles.lineNumStyle); err != nil {
		return
	}

	v.tui.Show()
	return
}
