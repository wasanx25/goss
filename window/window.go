package window

import (
	"syscall"
	"unsafe"
)

type Window struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func New() *Window {
	return &Window{}
}

func (w *Window) SetSize() error {
	_, _, err := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(w)),
	)

	if err != 0 {
		return err
	}

	return nil
}
