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
	Limit    int
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

func (d *Drawer) GetOffset() int {
	return d.offset
}

func (d *Drawer) AddOffset(e event.Type) {
	switch e {
	case event.PageDown:
		d.Increment()
	case event.PageUp:
		d.Decrement()
	case event.PageDownHalf:
		d.IncrementHalf()
	case event.PageUpHalf:
		d.DecrementHalf()
	case event.PageDownScreen:
		d.IncrementWindow()
	case event.PageUpScreen:
		d.DecrementWindow()
	}
}

func (d *Drawer) Increment() {
	if d.Max < d.Limit {
		return
	}

	if d.Limit+d.offset < d.Max {
		d.offset++
	}
}

func (d *Drawer) Decrement() {
	if d.offset > 1 {
		d.offset--
	}
}

func (d *Drawer) IncrementHalf() {
	if d.Max > d.offset+d.Limit/2*3 {
		d.offset = d.offset + d.Limit/2
	} else if d.offset+d.Limit < d.Max {
		d.offset = d.Max - d.Limit
	}
}

func (d *Drawer) DecrementHalf() {
	if d.offset > d.Limit/2 {
		d.offset = d.offset - d.Limit/2
	} else if d.offset > 1 && d.offset < d.Limit/2 {
		d.offset = 1
	}
}

func (d *Drawer) IncrementWindow() {
	if d.Max > d.offset+d.Limit*2 {
		d.offset = d.offset + d.Limit
	} else if d.offset+d.Limit < d.Max {
		d.offset = d.Max - d.Limit
	}
}

func (d *Drawer) DecrementWindow() {
	if d.offset > d.Limit {
		d.offset = d.offset - d.Limit
	} else if d.offset > 1 && d.offset < d.Limit {
		d.offset = 1
	}
}

func (d *Drawer) Get() (string, error) {
	scan := bufio.NewScanner(strings.NewReader(d.text))
	var (
		lines []string
		i     int
	)
	for scan.Scan() {
		i++
		if i < d.offset {
			continue
		} else if i >= d.Limit+d.offset {
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
