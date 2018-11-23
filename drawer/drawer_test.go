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
