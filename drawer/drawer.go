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
	return &Drawer{
		Text:   text,
		Offset: offset,
	}
}

func (d *Drawer) Increment() {
	d.Offset++
}

func (d *Drawer) Decrement() {
	if d.Offset > 1 {
		d.Offset--
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
	d.Position.Row, d.Position.Col = 1, 1
}
