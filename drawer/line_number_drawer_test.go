package drawer_test

import (
	"strings"
	"testing"

	"github.com/gdamore/tcell"
	"github.com/wasanx25/goss/drawer"
)

func TestLineNumberDrawerWrite(t *testing.T) {
	tui := tcell.NewSimulationScreen("")
	if err := tui.Init(); err != nil {
		t.Fatal(err)
	}

	maxLine := 12

	position := drawer.NewPositioner(1)
	l := drawer.NewLineNumberDrawer(maxLine, 0, position)

	tui.SetSize(5, maxLine)
	l.Write(tui, tcell.StyleDefault)
	tui.Show()

	// trim end space if use heredoc
	slice := []string{
		" 1   \n",
		" 2   \n",
		" 3   \n",
		" 4   \n",
		" 5   \n",
		" 6   \n",
		" 7   \n",
		" 8   \n",
		" 9   \n",
		" 10  \n",
		" 11  \n",
		" 12  \n",
	}
	expected := strings.Join(slice, "")

	actual := getString(tui)

	if actual != expected {
		t.Errorf("expected=%v, got=%v", expected, actual)
	}
}
