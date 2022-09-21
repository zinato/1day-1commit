package table

import (
	"errors"
	"fmt"
)

const (
	ErrDivisionZero    = "division by zero"
	ErrUnknownOperator = "unknown operator"
)

func DoMath(x, y int, op string) (int, error) {
	switch op {
	case "+":
		return x + y, nil
	case "-":
		return x - y, nil
	case "*":
		return x * y, nil
	case "/":
		if y == 0 {
			return 0, errors.New(ErrDivisionZero)
		}
		return x / y, nil
	default:
		return 0, fmt.Errorf("%v: %v", ErrUnknownOperator, op)
	}
}
