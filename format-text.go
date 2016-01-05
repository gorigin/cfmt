package cfmt

import (
	"strings"
)

// TextFormat performs alignment formatting
func TextFormat(f Format) string {
	str := f.String()

	if f.Width > 0 && len(str) < f.Width {
		// Padding
		switch f.Align {
		case RIGHT:
			str = strings.Repeat(" ", f.Width-len(str)) + str
		case CENTER:
			left := (f.Width - len(str)) / 2
			right := f.Width - len(str) - left
			str = strings.Repeat(" ", left) + str + strings.Repeat(" ", right)
		default:
			str = str + strings.Repeat(" ", f.Width-len(str))
		}
	}

	return str
}
