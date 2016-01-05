// +build freebsd netbsd openbsd linux darwin

package cfmt

import (
	"bytes"
	"os"
	"strconv"
	"syscall"
	"unsafe"
)

const (
	// Start ANSI color sequence
	Start = "\033["
	// Bold formatring
	Bold = "1;"
	// Reset is the ANSI reset escape sequence
	Reset = "\033[0m"
	// DefaultBG is the default background
	DefaultBG = "\033[49m"
	// DefaultFG is the default foreground
	DefaultFG = "\033[39m"
)

// AnsiFormat performs ANSI formatting of provided data
func AnsiFormat(f Format) string {
	str := TextFormat(f)

	if !f.HasColors() {
		return str
	}

	buf := bytes.NewBufferString(Start)
	if f.Bold {
		buf.WriteString("1;")
	} else if f.Underline {
		buf.WriteString("4;")
	} else {
		buf.WriteString("0;")
	}
	if f.Fg > 0 {
		if f.Fg > 15 {
			buf.WriteString("38;5;")
			buf.WriteString(strconv.Itoa(f.Fg))
		} else {
			clr := f.Fg + 29
			if f.Intense {
				clr += 60
			}
			buf.WriteString(strconv.Itoa(clr))
		}
	}
	if f.Bg > 0 {
		if f.Fg > 0 {
			buf.WriteString(";")
		}
		if f.Bg > 15 {
			buf.WriteString("48;5;")
			buf.WriteString(strconv.Itoa(f.Bg))
		} else {
			clr := f.Bg + 39
			buf.WriteString(strconv.Itoa(clr))
			buf.WriteString(";")
		}
	}
	buf.WriteString("m")

	buf.WriteString(str)
	buf.WriteString(Reset)

	return buf.String()
}

// isTerminal returns true if current context is terminal
func isTerminal(fd int) bool {
	var termios syscall.Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), ioctlReadTermios, uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
	return err == 0
}

func init() {
	if isTerminal(int(os.Stdout.Fd())) {
		formatter = AnsiFormat
	}
}
