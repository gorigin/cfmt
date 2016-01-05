package cfmt

// Static styles holder
var styles StyleSheet

// Style keys
const (
	S_ERROR = iota
	S_STRING
	S_INT
	S_DURATION
	S_KEY
	S_TYPE
	S_NIL
	S_BOOL
	S_TIME_LOG
)

// StyleSheet describes stylesheet holder format
type StyleSheet map[int]func(interface{}) Format

// Get returns formatting for requested type
func (s *StyleSheet) Get(t int, x interface{}) Format {
	if f, ok := styles[t]; ok {
		return f(x)
	}

	return noformat(x)
}

// noformat is utility function, that returns zero formatting function
func noformat(x interface{}) Format {
	return Format{Value: x}
}

// Init
func init() {
	styles = StyleSheet{}

	styles[S_NIL] = func(x interface{}) Format {
		return Format{
			Value: x,
			Fg:    161,
		}
	}
	styles[S_BOOL] = func(x interface{}) Format {
		return Format{
			Value: x,
			Fg:    161,
		}
	}
	styles[S_ERROR] = func(x interface{}) Format {
		return Format{
			Value: x,
			Bold:  true,
			Fg:    Red,
		}
	}
	styles[S_STRING] = func(x interface{}) Format {
		return Format{
			Value: x,
			Fg:    28,
		}
	}
	styles[S_DURATION] = func(x interface{}) Format {
		return Format{
			Value: x,
			Fg:    26,
		}
	}
	styles[S_INT] = func(x interface{}) Format {
		return Format{
			Value: x,
			Fg:    36,
		}
	}
	styles[S_KEY] = func(x interface{}) Format {
		return Format{
			Value: x,
			Fg:    136,
		}
	}
	styles[S_TYPE] = func(x interface{}) Format {
		return Format{
			Value: x,
			Bold:  true,
			Fg:    Black,
		}
	}
	styles[S_TIME_LOG] = func(x interface{}) Format {
		return Format{
			Value: x,
			Fg:    239,
		}
	}
}
