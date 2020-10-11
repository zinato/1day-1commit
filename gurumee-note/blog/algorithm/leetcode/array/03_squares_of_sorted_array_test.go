package array

import (
	"reflect"
	"sort"
	"testing"
)

func sortedSquares(A []int) []int {
	for i := 0; i < len(A); i++ {
		A[i] = A[i] * A[i]
	}

	sort.Ints(A)
	return A
}

func TestSortedSquares(t *testing.T) {
	input := []int{-7, -3, 2, 3, 11}
	output := sortedSquares(input)
	expected := []int{4, 9, 9, 49, 121}

	if reflect.DeepEqual(output, expected) {
		t.Log("test is completed")
	} else {
		t.Error("test is failed, it must be", expected, "but output is", output)
	}
}
