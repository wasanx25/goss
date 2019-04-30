package drawer_test

import (
	"strings"
	"testing"

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
	d.Reset()
	for _, tt := range tests {
		d.AddPosition(tt.args)

		col, row := d.Position()

		if col != tt.expectCol {
			t.Errorf(string(tt.args)+" is expected=%d, got=%d", tt.expectCol, col)
		}

		if row != tt.expectRow {
			t.Errorf(string(tt.args)+" is expected=%d, got=%d", tt.expectRow, row)
		}

		d.Reset() // reset
	}
}

func TestInitPosition(t *testing.T) {
	d := drawer.New("", 0, 0, 1) // dummy args
	d.Reset()

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
