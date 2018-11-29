package drawer_test

import (
	"testing"

	"github.com/wasanx25/goss/drawer"
)

func TestIncrement(t *testing.T) {
	d := drawer.New("test", 0)
	d.Increment()

	if d.Offset != 1 {
		t.Errorf("expected=1, got=%d", d.Offset)
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
		offset   uint
		limit    uint
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

func TestPositionInit(t *testing.T) {
	d := drawer.New("", 0) // dummy args
	d.PositionInit()

	if d.Position.Col != 1 {
		t.Errorf("expected=%d, got=%d", 1, d.Position.Col)
	}

	if d.Position.Row != 1 {
		t.Errorf("expected=%d, got=%d", 1, d.Position.Row)
	}
}
