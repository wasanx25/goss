package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/famz/SetLocale"
	gc "github.com/rthornton128/goncurses"
)

func main() {
	flag.Parse()
	fileName := flag.Arg(0)
	fmt.Println(fileName)
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Failed reading file. err: ", err)
		os.Exit(2)
	}

	SetLocale.SetLocale(SetLocale.LC_ALL, "")
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
	stdscr.Print(string(data))

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
