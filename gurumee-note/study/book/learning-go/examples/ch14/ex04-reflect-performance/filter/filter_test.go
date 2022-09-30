package filter

import (
	"fmt"
	"testing"
)

func filterShorterThan3(s string) bool {
	return len(s) <= 3
}

func BenchmarkFilterReflect(b *testing.B) {
	names := []string{"Andrew", "Bob", "Clara", "Hortense"}
	out := Filter(names, filterShorterThan3)
	fmt.Println("reflect: ", out)
}

func BenchmarkFilterString(b *testing.B) {
	names := []string{"Andrew", "Bob", "Clara", "Hortense"}
	out := FilterString(names, filterShorterThan3)
	fmt.Println("reflect: ", out)
}
