package drawer_test

import (
	"strings"
	"testing"

	"github.com/gdamore/tcell"
	"github.com/wasanx25/goss/drawer"
)

func TestContentDrawerWrite(t *testing.T) {
	tui := tcell.NewSimulationScreen("")
	if err := tui.Init(); err != nil {
		t.Fatal(err)
	}

	position := drawer.NewPositioner(1)
	c := drawer.NewContentDrawer("test1\ntest2 test3\ttest4\ntest5", 0, 5, position)

	tui.SetSize(30, 5)
	c.Write(tui, tcell.StyleDefault)
	tui.Show()

	// trim end space if use heredoc
	slice := []string{
		" test1                        \n",
		" test2  test3    test4        \n",
		" test5                        \n",
		"                              \n",
		"                              \n",
	}
	expected := strings.Join(slice, "")

	actual := getString(tui)

	if actual != expected {
		t.Errorf("expected=%v, got=%v", expected, actual)
	}
}
