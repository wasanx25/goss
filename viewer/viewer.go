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

// Viewer manages draw and event
type Viewer struct {
	tui           tcell.Screen
	contentDrawer drawer.Drawer
	numberDrawer  drawer.Drawer
	offsetter     offsetter.Offsetter
	event         *event.Event
	styles        *Styles
}

// Styles has all styles setting
type Styles struct {
	screenStyle  tcell.Style
	lineNumStyle tcell.Style
	contentStyle tcell.Style
}

// SetScreenStyle is goss main view styles
func (s *Styles) SetScreenStyle(style tcell.Style) { s.screenStyle = style }

// SetLineNumStyle is line number styles, it views left side and only integer
func (s *Styles) SetLineNumStyle(style tcell.Style) { s.lineNumStyle = style }

// SetContentStyle is main contents styles, there is setting 'text' in viewer.New
func (s *Styles) SetContentStyle(style tcell.Style) { s.contentStyle = style }

// New returns intializing Viewer pointer, there are args that want to view content for drawing
// and Screen interface for injection and styles options
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

// Open is opening goss
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
		return
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
					close(v.event.QuitCh)
				}
			case <-v.event.ResizeCh:
				v.setLimit()
				if err = v.rewrite(); err != nil {
					close(v.event.QuitCh)
				}
			}
		}
	}()
	<-v.event.QuitCh

	v.tui.Fini()
	return
}

func (v *Viewer) setLimit() {
	_, limitHeight := v.tui.Size()
	v.contentDrawer.SetLimitHeight(limitHeight)
}

func (v *Viewer) rewrite() error {
	v.tui.Clear()

	if err := v.contentDrawer.Write(v.tui, v.styles.contentStyle); err != nil {
		return err
	}

	if err := v.numberDrawer.Write(v.tui, v.styles.lineNumStyle); err != nil {
		return err
	}

	v.tui.Show()
	return nil
}
