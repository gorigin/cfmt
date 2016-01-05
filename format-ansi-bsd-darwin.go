// +build darwin dragonfly freebsd netbsd openbsd

package cfmt

import "syscall"

const ioctlReadTermios = syscall.TIOCGETA
