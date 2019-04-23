package viewer_test

import (
	"testing"

	"github.com/gdamore/tcell"
	"github.com/wasanx25/goss/viewer"
)

func TestWrite(t *testing.T) {
	var err error

	tui := tcell.NewSimulationScreen("")
	v := viewer.New("test1\ntest2 test3\ttest4\ntest5")

	if err = tui.Init(); err != nil {
		t.Fatal(err)
	}

	tui.SetSize(30, 5)

	viewer.SetTui(v, tui)
	viewer.SetLimit(v)
	viewer.Write(v)

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
