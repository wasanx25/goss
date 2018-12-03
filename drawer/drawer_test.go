package drawer_test

import (
	"testing"

	"github.com/wasanx25/goss/drawer"
)

func TestIncrement(t *testing.T) {
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
		d := drawer.New("test", tt.offset)
		d.Limit = tt.limit
		d.Max = tt.max

		d.Increment()
		if d.Offset != tt.expected {
			t.Errorf("expected=%d, got=%d", tt.expected, d.Offset)
		}
	}
}

func TestIncrementHalf(t *testing.T) {
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
		d := drawer.New("test", tt.offset)
		d.Limit = tt.limit
		d.Max = tt.max

		d.IncrementHalf()
		if d.Offset != tt.expected {
			t.Errorf("expected=%d, got=%d", tt.expected, d.Offset)
		}
	}
}

func TestDecrementHalf(t *testing.T) {
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
		d := drawer.New("test", tt.offset)
		d.Limit = tt.limit
		d.Max = tt.max

		d.DecrementHalf()
		if d.Offset != tt.expected {
			t.Errorf("expected=%d, got=%d", tt.expected, d.Offset)
		}
	}
}

func TestIncrementWindow(t *testing.T) {
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
		d := drawer.New("test", tt.offset)
		d.Limit = tt.limit
		d.Max = tt.max

		d.IncrementWindow()
		if d.Offset != tt.expected {
			t.Errorf("expected=%d, got=%d", tt.expected, d.Offset)
		}
	}
}

func TestDecrementWindow(t *testing.T) {
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
		d := drawer.New("test", tt.offset)
		d.Limit = tt.limit
		d.Max = tt.max

		d.DecrementWindow()
		if d.Offset != tt.expected {
			t.Errorf("expected=%d, got=%d", tt.expected, d.Offset)
		}
	}
}

func TestDecrement(t *testing.T) {
	d := drawer.New("test", 10)
	d.Decrement()

	if d.Offset != 9 {
		t.Errorf("expected=0, got=%d", d.Offset)
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		text     string
		offset   int
		limit    int
		expected string
	}{
		{"test1", 1, 1, "test1"},
		{"test1\ntest2\ntest3", 1, 1, "test1"},
		{"test1\ntest2\ntest3", 2, 2, "test2\ntest3"},
		{"test1\ntest2\ntest3", 1, 3, "test1\ntest2\ntest3"},
		{"test1\ntest2\ntest3", 2, 3, "test2\ntest3"},
	}

	for _, tt := range tests {
		d := drawer.New(tt.text, tt.offset)
		d.Limit = tt.limit
		result, err := d.Get()

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
		{'\t', 4, 0},
		{'a', 1, 0},
		{'か', 2, 0},
	}

	d := drawer.New("", 0) // dummy args
	for _, tt := range tests {
		d.AddPosition(tt.args)

		if d.Position.Col != tt.expectCol {
			t.Errorf(string(tt.args)+" is expected=%d, got=%d", tt.expectCol, d.Position.Col)
		}

		if d.Position.Row != tt.expectRow {
			t.Errorf(string(tt.args)+" is expected=%d, got=%d", tt.expectRow, d.Position.Row)
		}

		d.Position.Row, d.Position.Col = 0, 0 // reset
	}
}

func TestInitPosition(t *testing.T) {
	d := drawer.New("", 0) // dummy args
	d.InitPosition()

	if d.Position.Col != 1 {
		t.Errorf("expected=%d, got=%d", 1, d.Position.Col)
	}

	if d.Position.Row != 0 {
		t.Errorf("expected=%d, got=%d", 0, d.Position.Row)
	}
}
