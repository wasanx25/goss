package viewer_test

import (
	"github.com/gdamore/tcell"
	"github.com/wasanx25/goss/viewer"
	"testing"
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

	expected := ` 1   test1                    
 2   test2  test3    test4    
 3   test5                    
 4                            
 5                            
`

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