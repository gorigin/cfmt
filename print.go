package cfmt

import (
	"fmt"
	"io"
	"os"
)

// Replaces entities with known ones
func replace(x []interface{}) []interface{} {
	if x == nil {
		return nil
	} else if len(x) == 0 {
		return x
	}

	y := []interface{}{}

	for _, j := range x {
		// Real format
		if f, ok := j.(Format); ok {
			y = append(y, formatter(f))
		} else if f, ok := j.(FormatList); ok {
			y = append(y, replace(f)...)
		} else if f, ok := j.(Formatted); ok {
			y = append(y, replace(f.Cfmt())...)
		} else if f, ok := j.(error); ok {
			y = append(y, formatter(styles.Get(S_ERROR, f)))
		} else {
			y = append(y, j)
		}
	}

	return y
}

// Fprint prints data to provided Writer
func Fprint(w io.Writer, x ...interface{}) {
	fmt.Fprint(w, replace(x)...)
}

// Fprintf prints data to provided Writer with pattern
func Fprintf(w io.Writer, pattern string, x ...interface{}) {
	fmt.Fprintf(w, pattern, replace(x)...)
}

// Printf prints data with pattern to default Writer (os.Stdout)
func Printf(pattern string, x ...interface{}) {
	Fprintf(os.Stdout, pattern, replace(x)...)
}

// Println outputs dat a to default Writer (os.Stdout) with line end
func Println(x ...interface{}) {
	if len(x) == 0 {
		fmt.Fprintf(os.Stdout, "\n")
	} else {
		y := append(x, "\n")
		Print(y...)
	}
}

// Print outputs dat a to default Writer (os.Stdout)
func Print(x ...interface{}) {
	Fprint(os.Stdout, x...)
}
