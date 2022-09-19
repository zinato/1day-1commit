package adder_test

import (
	"learning-go/examples/ch13/ex01-test-basic/adder"
	"testing"
)

func Test_Add(t *testing.T) {
	expected := 3
	result := adder.Add(1, 2)

	if expected != result {
		t.Errorf("incorrect result: expected %v, got %v", expected, result)
	}
}
