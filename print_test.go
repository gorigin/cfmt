package cfmt

import (
	"testing"
)

func BenchmarkReplaceStrings(b *testing.B) {
	a := []interface{}{
		"foo",
		"bar",
		"\n",
	}

	for i := 0; i < b.N; i++ {
		replace(a)
	}
}

func BenchmarkReplaceEmpty(b *testing.B) {
	a := []interface{}{}

	for i := 0; i < b.N; i++ {
		replace(a)
	}
}

func BenchmarkReplaceFormat1(b *testing.B) {
	a := []interface{}{
		"foo",
		Format{Value: "bar", Bold: true, Fg: 15},
		"\n",
	}

	for i := 0; i < b.N; i++ {
		replace(a)
	}
}

func BenchmarkReplaceFormat2(b *testing.B) {
	a := []interface{}{
		"foo",
		Format{Value: "bar", Bold: true, Fg: 15},
		Format{Value: "baz", Bold: true, Fg: 15},
	}

	for i := 0; i < b.N; i++ {
		replace(a)
	}
}
