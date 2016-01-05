package cfmt

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// FAuto detects type of provided value and returns most suitable format
// for it
func FAuto(in interface{}) Format {
	if in == nil {
		return FNil()
	}
	if x, ok := in.(string); ok {
		return FString(x)
	}
	if x, ok := in.(int); ok {
		return FInt(x)
	}
	if x, ok := in.(bool); ok {
		return FBool(x)
	}
	if x, ok := in.(error); ok {
		return FError(x)
	}
	if x, ok := in.(Format); ok {
		return x
	}

	return Format{
		Value: in,
	}
}

// FHeader returns formatting for header
func FHeader(h string) FormatList {
	return []interface{}{
		" ",
		Format{Value: h, Fg: 70},
		"\n",
		" ",
		Format{Value: strings.Repeat("─", len(h)), Fg: 64},
		"\n",
	}
}

// FNil returns format for nil value
func FNil() Format {
	return styles.Get(S_NIL, "<nil>")
}

// FError returns format for errors
func FError(value interface{}) Format {
	if x, ok := value.(error); ok {
		return styles.Get(S_ERROR, x.Error())
	}

	return styles.Get(S_ERROR, value)
}

// FBool returns true/false formatting for bools
func FBool(value bool) Format {
	if value {
		return styles.Get(S_BOOL, "true")
	}

	return styles.Get(S_BOOL, "false")
}

// FYesNo returns yes/no formatting for bools
func FYesNo(value bool) Format {
	if value {
		return styles.Get(S_BOOL, "yes")
	}

	return styles.Get(S_BOOL, "no")
}

// FString returns format for strings
func FString(s string) Format {
	return styles.Get(S_STRING, s)
}

// FStringl returns format for string slice
func FStringl(s []string) Format {
	if s == nil {
		return FNil()
	} else if len(s) == 0 {
		return styles.Get(S_STRING, "")
	}
	buf := bytes.NewBufferString("")
	for i, v := range s {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(v)
	}
	return styles.Get(S_STRING, buf.String())
}

// FInt returns format for integers
func FInt(x int) Format {
	return styles.Get(S_INT, x)
}

// FFloat returns format for floats
func FFloat(value float64, precision int) Format {
	return styles.Get(S_INT, fmt.Sprintf("%."+strconv.Itoa(precision)+"f", value))
}

// FType returns format for types (structs, etc.)
func FType(x interface{}) Format {
	return styles.Get(S_TYPE, x)
}

// FKey returns format for map keys
func FKey(x string) Format {
	return styles.Get(S_KEY, x)
}

// FDuration returns format for durations
func FDuration(x time.Duration) Format {
	return styles.Get(S_DURATION, x)
}

// FDuration2 returns format for durations with arbitrary precision
func FDuration2(duration time.Duration, dimension time.Duration) Format {
	switch dimension {
	case time.Second:
		return styles.Get(S_DURATION, fmt.Sprintf("%.3f s", duration.Seconds()))
	case time.Millisecond:
		return styles.Get(S_DURATION, fmt.Sprintf("%d ms", duration.Nanoseconds()/1000000))
	default:
		return FDuration(duration)
	}
}

// FTimeShortMs returns formatting for time with millisecond precision
func FTimeShortMs(time time.Time) Format {
	return styles.Get(
		S_TIME_LOG,
		fmt.Sprintf(
			"%02d:%02d:%02d.%03d",
			time.Hour(),
			time.Minute(),
			time.Second(),
			time.Nanosecond()/1000000,
		),
	)
}
