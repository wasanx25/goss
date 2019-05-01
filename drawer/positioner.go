package drawer

import "github.com/mattn/go-runewidth"

type Positioner interface {
	Add(r rune)
	X() int
	Y() int
	XAndY() (int, int)
	Reset()
	Break()
}

type positioner struct {
	x         int
	y         int
	rowNumMax int
}

const (
	TAB      = '\t'
	NEW_LINE = '\n'
	SPACE    = ' '
)

func NewPositioner(rowNumMax int) Positioner {
	return &positioner{
		x:         0,
		y:         0,
		rowNumMax: rowNumMax,
	}
}

func (p *positioner) Add(r rune) {
	switch r {
	case TAB:
		p.x += 4
	case NEW_LINE:
		p.x = p.rowNumMax
		p.y++
	case SPACE:
		p.x++
	}
	p.x += runewidth.RuneWidth(r)
}

func (p *positioner) X() int {
	return p.x
}

func (p *positioner) Y() int {
	return p.y
}

func (p *positioner) XAndY() (int, int) {
	return p.x, p.y
}

func (p *positioner) Reset() {
	p.x, p.y = p.rowNumMax, 0
}

func (p *positioner) Break() {
	p.x = 1
	p.y++
}
