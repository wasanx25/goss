package drawer

import (
	"bufio"
	"strings"

	"github.com/gdamore/tcell"
)

// TODO already use Drawer... (struct)
type DDrawer interface {
	SetLimitHeight(limitHeight int)
	Write(tui tcell.Screen, style tcell.Style) error
}

type contentDrawer struct {
	text        string
	maxLine     int
	offset      int
	limitHeight int
	position    Positioner
}

func NewContentDrawer(text string, maxLine, offset, limitHeight int, positioner Positioner) DDrawer {
	return &contentDrawer{
		text:        text,
		maxLine:     maxLine,
		offset:      offset,
		limitHeight: limitHeight,
		position:    positioner,
	}
}

func (c *contentDrawer) SetLimitHeight(limitHeight int) {
	c.limitHeight = limitHeight
}

func (c *contentDrawer) Write(tui tcell.Screen, style tcell.Style) error {
	c.position.Reset()
	viewText, err := c.getContent()
	if err != nil {
		return err
	}

	windowWidth, windowHeight := tui.Size()

	for _, r := range viewText {
		x, y := c.position.XAndY()
		if x >= windowWidth {
			c.position.Break()
		}

		tui.SetContent(x, y, r, nil, style)
		c.position.Add(r)

		if windowHeight < y {
			break
		}
	}

	return nil
}

func (c *contentDrawer) getContent() (string, error) {
	scan := bufio.NewScanner(strings.NewReader(c.text))
	var (
		lines []string
		i     int
	)
	for scan.Scan() {
		i++
		if i <= c.offset {
			continue
		} else if i >= c.limitHeight+c.offset+1 {
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
