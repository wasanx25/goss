package drawer_test

import (
	"strings"
	"testing"

	"github.com/gdamore/tcell"
	"github.com/wasanx25/goss/drawer"
	"github.com/wasanx25/goss/event"
)

func TestAddOffset(t *testing.T) {
	tests := []struct {
		eventType event.Type
		offset    int
		max       int
		limit     int
		expected  int
	}{
		{event.PageDown, 1, 10, 1, 1 + 1},
		{event.PageDown, 8, 10, 1, 8 + 1},
		{event.PageDown, 12, 10, 1, 12},
		{event.PageDown, 1, 10, 11, 1},
		{event.PageUp, 10, 1, 1, 9},
		{event.PageDownHalf, 1, 10, 6, 1 + 3},
		{event.PageDownHalf, 7, 10, 2, 7 + 1},
		{event.PageDownHalf, 12, 10, 1, 12},
		{event.PageDownHalf, 1, 10, 11, 1},
		{event.PageUpHalf, 1, 10, 6, 1},
		{event.PageUpHalf, 7, 10, 2, 7 - 1},
		{event.PageUpHalf, 3, 10, 8, 1},
		{event.PageDownScreen, 1, 10, 3, 4},
		{event.PageDownScreen, 3, 10, 6, 4},
		{event.PageDownScreen, 5, 10, 5, 5},
		{event.PageUpScreen, 1, 10, 6, 1},
		{event.PageUpScreen, 7, 10, 2, 7 - 2},
		{event.PageUpScreen, 3, 10, 8, 1},
		{event.PageEnd, 10, 3, 2, 3 - 1},
		{event.PageTop, 10, 0, 0, 0},
	}

	for _, tt := range tests {
		d := drawer.New("test", tt.offset, tt.max, 1)
		d.SetLimit(tt.limit)

		d.AddOffset(tt.eventType)
		if d.Offset() != tt.expected {
			t.Errorf("eventType=%v, expected=%d, got=%d", tt.eventType, tt.expected, d.Offset())
		}
	}
}

func TestGetContent(t *testing.T) {
	tests := []struct {
		text     string
		offset   int
		limit    int
		expected string
	}{
		{"test1", 1, 1, ""},
		{"test1\ntest2\ntest3", 1, 1, "test2"},
		{"test1\ntest2\ntest3", 2, 2, "test3"},
		{"test1\ntest2\ntest3", 1, 3, "test2\ntest3"},
		{"test1\ntest2\ntest3", 2, 3, "test3"},
	}

	for _, tt := range tests {
		d := drawer.New(tt.text, tt.offset, strings.Count(tt.text, "\n"), 1)
		d.SetLimit(tt.limit)
		result, err := d.GetContent()

		if err != nil {
			t.Fatal(err)
		}

		if result != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, result)
		}
	}
}

func TestAddPosition(t *testing.T) {
	tests := []struct {
		args      rune
		expectCol int
		expectRow int
	}{
		{'\n', 1, 1},
		{'\t', 5, 0},
		{'a', 2, 0},
		{'か', 3, 0},
	}

	d := drawer.New("", 0, 0, 1) // dummy args
	d.PositionReset()
	for _, tt := range tests {
		d.AddPosition(tt.args)

		col, row := d.Position()

		if col != tt.expectCol {
			t.Errorf(string(tt.args)+" is expected=%d, got=%d", tt.expectCol, col)
		}

		if row != tt.expectRow {
			t.Errorf(string(tt.args)+" is expected=%d, got=%d", tt.expectRow, row)
		}

		d.PositionReset() // reset
	}
}

func TestInitPosition(t *testing.T) {
	d := drawer.New("", 0, 0, 1) // dummy args
	d.PositionReset()

	col, row := d.Position()

	if col != 1 {
		t.Errorf("expected=%d, got=%d", 1, col)
	}

	if row != 0 {
		t.Errorf("expected=%d, got=%d", 0, row)
	}
}

func TestBreak(t *testing.T) {
	d := drawer.New("", 0, 0, 1) // dummy args
	d.Break()

	col, row := d.Position()

	if col != 1 {
		t.Errorf("expected=%d, got=%d", 1, col)
	}

	if row != 1 {
		t.Errorf("expected=%d, got=%d", 1, row)
	}
}

func TestWrite(t *testing.T) {
	tui := tcell.NewSimulationScreen("")
	if err := tui.Init(); err != nil {
		t.Fatal(err)
	}

	d := drawer.New("test1\ntest2 test3\ttest4\ntest5", 0, 2, 5)
	d.SetLimit(5)

	tui.SetSize(30, 5)

	d.Write(tui, tcell.StyleDefault, tcell.StyleDefault)

	tui.Show()

	// trim end space if use heredoc
	slice := []string{
		" 1   test1                    \n",
		" 2   test2  test3    test4    \n",
		" 3   test5                    \n",
		"                              \n",
		"                              \n",
	}
	expected := slice[0] + slice[1] + slice[2] + slice[3] + slice[4]

	actual := getString(tui)

	if actual != expected {
		t.Errorf("expected=%v, got=%v", expected, actual)
	}
}

func getString(tui tcell.SimulationScreen) string {
	width, _ := tui.Size()
	cells, _, _ := tui.GetContents()

	var runes []rune
	for i, c := range cells {
		runes = append(runes, c.Runes...)
		if (i+1)%width == 0 {
			runes = append(runes, '\n')
		}
	}

	return string(runes)
}
