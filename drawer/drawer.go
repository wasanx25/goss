package drawer

import (
	"bufio"
	"strings"

	runewidth "github.com/mattn/go-runewidth"
	"github.com/wasanx25/goss/event"
)

type Drawer struct {
	text     string
	offset   int
	limit    int
	Max      int
	Position DrawPosition
}

type DrawPosition struct {
	Row int
	Col int
}

const (
	TAB      = '\t'
	NEW_LINE = '\n'
	SPACE    = ' '
)

func New(text string, offset int) *Drawer {
	max := strings.Count(text, "\n")
	return &Drawer{
		text:   text,
		offset: offset,
		Max:    max,
	}
}

func (d *Drawer) Offset() int {
	return d.offset
}

func (d *Drawer) Limit() int {
	return d.limit
}

func (d *Drawer) SetLimit(limit int) {
	d.limit = limit
}

func (d *Drawer) AddOffset(e event.Type) {
	switch e {
	case event.PageDown:
		d.pageDown()
	case event.PageUp:
		d.pageUp()
	case event.PageDownHalf:
		d.pageDownHalf()
	case event.PageUpHalf:
		d.pageUpHalf()
	case event.PageDownScreen:
		d.pageDownWindow()
	case event.PageUpScreen:
		d.pageUpWindow()
	case event.PageEnd:
		d.pageEnd()
	case event.PageTop:
		d.pageTop()
	}
}

func (d *Drawer) pageDown() {
	if d.Max < d.limit {
		return
	}

	if d.limit+d.offset < d.Max {
		d.offset++
	}
}

func (d *Drawer) pageUp() {
	if d.offset > 1 {
		d.offset--
	}
}

func (d *Drawer) pageDownHalf() {
	if d.Max > d.offset+d.limit/2*3 {
		d.offset = d.offset + d.limit/2
	} else if d.offset+d.limit < d.Max {
		d.offset = d.Max - d.limit
	}
}

func (d *Drawer) pageUpHalf() {
	if d.offset > d.limit/2 {
		d.offset = d.offset - d.limit/2
	} else if d.offset > 1 && d.offset < d.limit/2 {
		d.offset = 1
	}
}

func (d *Drawer) pageDownWindow() {
	if d.Max > d.offset+d.limit*2 {
		d.offset = d.offset + d.limit
	} else if d.offset+d.limit < d.Max {
		d.offset = d.Max - d.limit
	}
}

func (d *Drawer) pageUpWindow() {
	if d.offset > d.limit {
		d.offset = d.offset - d.limit
	} else if d.offset > 1 && d.offset < d.limit {
		d.offset = 1
	}
}

func (d *Drawer) pageEnd() {
	d.offset = d.Max - d.limit
}

func (d *Drawer) pageTop() {
	d.offset = 0
}

func (d *Drawer) GetContent() (string, error) {
	scan := bufio.NewScanner(strings.NewReader(d.text))
	var (
		lines []string
		i     int
	)
	for scan.Scan() {
		i++
		if i < d.offset {
			continue
		} else if i >= d.limit+d.offset {
			break
		}
		lines = append(lines, scan.Text())
	}
	err := scan.Err()
	if err != nil {
		return "", err
	}
	return strings.Join(lines, "\n"), nil
}

func (d *Drawer) AddPosition(r rune) {
	switch r {
	case TAB:
		d.Position.Col += 4
	case NEW_LINE:
		d.Position.Col = 1
		d.Position.Row++
	case SPACE:
		d.Position.Col++
	}
	d.Position.Col += runewidth.RuneWidth(r)
}

func (d *Drawer) InitPosition() {
	d.Position.Row, d.Position.Col = 0, 1
}

func (d *Drawer) Break() {
	d.Position.Row++
	d.Position.Col = 1
}
