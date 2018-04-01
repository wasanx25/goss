package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/famz/SetLocale"
	gc "github.com/rthornton128/goncurses"
)

func main() {
	flag.Parse()
	fileName := flag.Arg(0)
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

	_, x := stdscr.MaxYX()
	dataLN := strings.Count(string(data), "\n")
	stdscr.Keypad(true)
	stdscr.Resize(dataLN+100, x)
	stdscr.Print(string(data))
	stdscr.Refresh()
	stdscr.ScrollOk(true)

loop:
	for {
		k := stdscr.GetChar()
		switch byte(k) {
		case 'j':
			stdscr.Scroll(1)
		case 'k':
			stdscr.Scroll(-1)
		case 'q':
			stdscr.Print("oh my god.")
			break loop
		}
	}
}
