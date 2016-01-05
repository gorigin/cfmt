package cfmt

import (
	"fmt"
	"os"
)

// Panic shows error message and stops application
// with status code 1
func Panic(err error) {
	Print(err, "\n")
	os.Exit(1)
}

// Panicf shows error and stops application
func Panicf(message string, args ...interface{}) {
	Panic(fmt.Errorf(message, args...))
}
