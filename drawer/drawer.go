package drawer

import (
	"bufio"
	"strings"
)

type Drawer struct {
	Body   string
	Offset int
}

const (
	TAB      = '\t'
	NEW_LINE = '\n'
	SPACE    = ' '
)

func New(body string, offset int) *Drawer {
	return &Drawer{
		Body:   body,
		Offset: offset,
	}
}

func (d *Drawer) Increment() {
	d.Offset++
}

func (d *Drawer) Get() (string, error) {
	scan := bufio.NewScanner(strings.NewReader(d.Body))
	var lines []string
	var i int
	for scan.Scan() {
		i++
		if i < d.Offset {
			continue
		}
		lines = append(lines, scan.Text())
	}
	err := scan.Err()
	if err != nil {
		return "", err
	}
	return strings.Join(lines, "\n"), nil
}
