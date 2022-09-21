package table

import (
	"fmt"
	"strings"
	"testing"
)

func TestDoMath(t *testing.T) {
	data := []struct {
		name     string
		num1     int
		num2     int
		op       string
		expected int
	}{
		{"addition", 2, 2, "+", 4},
		{"subtraction", 2, 2, "-", 0},
		{"multiplication", 2, 2, "*", 4},
		{"division", 2, 2, "/", 1},
		{"division zero", 2, 0, "/", 0},
		{"bad operation", 2, 2, "?", 0},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, err := DoMath(d.num1, d.num2, d.op)
			if result != d.expected {
				t.Errorf("Expected %d, got %d", d.expected, result)
			}
			if err != nil {
				if msg := err.Error(); msg == ErrDivisionZero || strings.Contains(msg, ErrUnknownOperator) {
					fmt.Printf("expected error: %v", msg)
				} else {
					t.Errorf(msg)
				}
			}
		})
	}
}
