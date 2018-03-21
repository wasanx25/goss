package main

import (
	"fmt"
	"os"

	gc "github.com/rthornton128/goncurses"
)

func main() {
	stdscr, err := gc.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	defer gc.End()
	gc.StartColor()
	gc.Raw(true)
	gc.Echo(false)
	stdscr.Keypad(true)

loop:
	for {
		var y int
		k := stdscr.GetChar()
		switch byte(k) {
		case 'j':
			y, _ = stdscr.CursorYX()
			stdscr.Move(y+1, 0)
		case 'k':
			y, _ = stdscr.CursorYX()
			stdscr.Move(y-1, 0)
		case 'q':
			stdscr.Print("oh my god.")
			break loop
		}
	}
}
