package drawer

import (
	"bufio"
	"strings"

	runewidth "github.com/mattn/go-runewidth"
)

type Drawer struct {
	Text     string
	Offset   int
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
		Text:   text,
		Offset: offset,
		Max:    max,
	}
}

func (d *Drawer) Increment() {
	if d.Max < d.Limit {
		return
	}

	if d.Limit+d.Offset < d.Max {
		d.Offset++
	}
}

func (d *Drawer) Decrement() {
	if d.Offset > 1 {
		d.Offset--
	}
}

func (d *Drawer) IncrementHalf() {
	if d.Max > d.Offset+d.Limit/2*3 {
		d.Offset = d.Offset + d.Limit/2
	} else if d.Offset+d.Limit < d.Max {
		d.Offset = d.Max - d.Limit
	}
}

func (d *Drawer) DecrementHalf() {
	if d.Offset > d.Limit/2 {
		d.Offset = d.Offset - d.Limit/2
	} else if d.Offset > 1 && d.Offset < d.Limit/2 {
		d.Offset = 1
	}
}

func (d *Drawer) Get() (string, error) {
	scan := bufio.NewScanner(strings.NewReader(d.Text))
	var (
		lines []string
		i     int
	)
	for scan.Scan() {
		i++
		if i < d.Offset {
			continue
		} else if i >= d.Limit+d.Offset {
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

func (d *Drawer) PositionInit() {
	d.Position.Row, d.Position.Col = 0, 1
}
