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
	tests := []struct{
		args uint
		body string
		offset uint
		expected string
	}{
		{1, "test1", 1, "test1"},
		{1, "test1\ntest2\ntest3", 1, "test1"},
		{2, "test1\ntest2\ntest3", 2, "test2\ntest3"},
		{3, "test1\ntest2\ntest3", 1, "test1\ntest2\ntest3"},
		{3, "test1\ntest2\ntest3", 2, "test2\ntest3"},
	}

	for _, tt := range tests {
		d := drawer.New(tt.body, tt.offset)
		result, _ := d.Get(tt.args)

		if result != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, result)
		}
	}

}