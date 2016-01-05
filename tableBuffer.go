package cfmt

// TableBuffer2 is helper structure, used to build common tables
// with two columns
type TableBuffer2 struct {
	width1, width2 int
	values         []Format
}

// Add adds row to table buffer
func (t *TableBuffer2) Add(one string, two interface{}) *TableBuffer2 {
	c1 := FKey(one)
	c2 := FAuto(two)

	l1 := len(c1.String())
	l2 := len(c2.String())

	if l1 > t.width1 {
		t.width1 = l1
	}

	if l2 > t.width2 {
		t.width2 = l2
	}

	t.values = append(t.values, c1, c2)

	return t
}

// Cfmt returns lists of format elements
func (t TableBuffer2) Cfmt() []interface{} {
	if len(t.values) == 0 {
		// Empty values
		return []interface{}{}
	}

	separator := Format{
		Value: " | ",
		Fg:    237,
	}

	total := len(t.values) / 2
	response := make([]interface{}, total*4)
	for i := 0; i < total; i++ {
		j := i * 4
		k := t.values[i*2]
		v := t.values[i*2+1]

		k.Width = t.width1
		v.Width = t.width2

		response[j] = k
		response[j+1] = separator
		response[j+2] = v
		response[j+3] = "\n"
	}

	return response
}

// Print outputs table contents
func (t TableBuffer2) Print() {
	Print(t)
}
