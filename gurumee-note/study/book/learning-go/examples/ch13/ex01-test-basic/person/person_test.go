package person

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewPerson(t *testing.T) {
	comparer := cmp.Comparer(func(x, y Person) bool {
		return x.Name == y.Name && x.Age == y.Age
	})

	expected := Person{
		Name: "Denis",
		Age:  37,
	}
	result := NewPerson("Denis", 37)
	if diff := cmp.Diff(expected, result, comparer); diff != "" {
		t.Errorf(diff)
	}
}
