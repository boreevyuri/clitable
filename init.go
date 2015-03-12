package clitable

import (
	"runtime"
	"syscall"
	"unsafe"
//	"os"
//	"fmt"
)

type WindowSize struct {
	Row uint16
	Col uint16
}

var (
	EOL         = []byte{'\n'}
	WS          = " "

	WinSize     *WindowSize
	_TIOCGWINSZ int64
)

func init() {
	WinSize = new(WindowSize)

	switch runtime.GOOS {
	case "linux": _TIOCGWINSZ = 0x5413
	case "darwin": _TIOCGWINSZ = 1074295912
	case "windows": EOL = []byte{'\r', '\n'}
	}

	r1, _, _ := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(_TIOCGWINSZ),
		uintptr(unsafe.Pointer(WinSize)),
	)

	if int(r1) == -1 {
//		fmt.Println("Error:", os.NewSyscallError("window size", errno))
	}
}
