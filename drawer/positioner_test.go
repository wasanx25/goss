package drawer_test

import (
	"testing"

	"github.com/wasanx25/goss/drawer"
)

func TestPositioner(t *testing.T) {
	tests := []struct {
		args      rune
		expectedX int
		expectedY int
	}{
		{'\n', 1, 1},
		{'\t', 5, 0},
		{'a', 2, 0},
		{'か', 3, 0},
	}

	rowNumMax := 1

	p := drawer.NewPositioner(rowNumMax) // dummy args

	t.Run("Default Value", func(t *testing.T) {
		if p.X() != 0 {
			t.Errorf("Default X is expected=0, got=%d", p.X())
		}

		if p.Y() != 0 {
			t.Errorf("Default Y is expected=0, got=%d", p.Y())
		}
	})

	t.Run("Break()", func(t *testing.T) {
		oldY := p.Y()
		p.Break()
		x, y := p.XAndY()

		if x != 1 {
			t.Errorf("After call Break() value X is expected=1, got=%d", x)
		}

		expectedY := oldY + 1
		if y != expectedY {
			t.Errorf("After call Break() value Y is expected=%d, got=%d", expectedY, y)
		}
	})

	// Even if it is first parameter,
	// execute tests under the same condition of other parameters.
	p.Reset()

	for _, tt := range tests {
		p.Add(tt.args)

		x, y := p.XAndY()

		if x != tt.expectedX {
			t.Errorf(string(tt.args)+" is expected=%d, got=%d", tt.expectedX, x)
		}

		if y != tt.expectedY {
			t.Errorf(string(tt.args)+" is expected=%d, got=%d", tt.expectedY, y)
		}

		p.Reset()
	}
}
