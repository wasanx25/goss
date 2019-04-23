package drawer_test

import (
	"strings"
	"testing"

	"github.com/wasanx25/goss/drawer"
)

// pageDown is private method. export_test.go is necessary to run below test
func TestPageDown(t *testing.T) {
	tests := []struct {
		max      int
		offset   int
		limit    int
		expected int
	}{
		{10, 1, 1, 1 + 1},
		{10, 8, 1, 8 + 1},
		{10, 12, 1, 12},
		{10, 1, 11, 1},
	}

	for _, tt := range tests {
		d := drawer.New("test", tt.offset, tt.max, 1)
		d.SetLimit(tt.limit)

		drawer.PageDown(d)
		if d.Offset() != tt.expected {
			t.Errorf("expected=%d, got=%d", tt.expected, d.Offset())
		}
	}
}

// pageUp is private method. export_test.go is necessary to run below test
func TestPageUp(t *testing.T) {
	d := drawer.New("test", 10, 1, 1)
	drawer.PageUp(d)

	if d.Offset() != 9 {
		t.Errorf("expected=0, got=%d", d.Offset())
	}
}

// pageDownHalf is private method. export_test.go is necessary to run below test
func TestPageDownHalf(t *testing.T) {
	tests := []struct {
		max      int
		offset   int
		limit    int
		expected int
	}{
		{10, 1, 6, 1 + 3},
		{10, 7, 2, 7 + 1},
		{10, 12, 1, 12},
		{10, 1, 11, 1},
	}

	for _, tt := range tests {
		d := drawer.New("test", tt.offset, tt.max, 1)
		d.SetLimit(tt.limit)

		drawer.PageDownHalf(d)
		if d.Offset() != tt.expected {
			t.Errorf("expected=%d, got=%d", tt.expected, d.Offset())
		}
	}
}

// pageUpHalf is private method. export_test.go is necessary to run below test
func TestPageUpHalf(t *testing.T) {
	tests := []struct {
		max      int
		offset   int
		limit    int
		expected int
	}{
		{10, 1, 6, 1},
		{10, 7, 2, 7 - 1},
		{10, 3, 8, 1},
	}

	for _, tt := range tests {
		d := drawer.New("test", tt.offset, tt.max, 1)
		d.SetLimit(tt.limit)

		drawer.PageUpHalf(d)
		if d.Offset() != tt.expected {
			t.Errorf("expected=%d, got=%d", tt.expected, d.Offset())
		}
	}
}

// pageDownWindow is private method. export_test.go is necessary to run below test
func TestPageDownWindow(t *testing.T) {
	tests := []struct {
		max      int
		offset   int
		limit    int
		expected int
	}{
		{10, 1, 3, 4},
		{10, 3, 6, 4},
		{10, 5, 5, 5},
	}

	for _, tt := range tests {
		d := drawer.New("test", tt.offset, tt.max, 1)
		d.SetLimit(tt.limit)

		drawer.PageDownWindow(d)
		if d.Offset() != tt.expected {
			t.Errorf("expected=%d, got=%d", tt.expected, d.Offset())
		}
	}
}

// pageUpWindow is private method. export_test.go is necessary to run below test
func TestPageUpWindow(t *testing.T) {
	tests := []struct {
		max      int
		offset   int
		limit    int
		expected int
	}{
		{10, 1, 6, 1},
		{10, 7, 2, 7 - 2},
		{10, 3, 8, 1},
	}

	for _, tt := range tests {
		d := drawer.New("test", tt.offset, tt.max, 1)
		d.SetLimit(tt.limit)

		drawer.PageUpWindow(d)
		if d.Offset() != tt.expected {
			t.Errorf("expected=%d, got=%d", tt.expected, d.Offset())
		}
	}
}

// pageEnd is private method. export_test.go is necessary to run below test
func TestPageEnd(t *testing.T) {
	d := drawer.New("test\nlong\ntext\ntest", 10, 3, 1)
	d.SetLimit(2)
	drawer.PageEnd(d)

	if d.Offset() != (3 - 1) {
		t.Errorf("expected=1, got=%d", d.Offset())
	}
}

// pageTop is private method. export_test.go is necessary to run below test
func TestPageTop(t *testing.T) {
	d := drawer.New("test", 10, 0, 1)
	drawer.PageTop(d)

	if d.Offset() != 0 {
		t.Errorf("expected=0, got=%d", d.Offset())
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
