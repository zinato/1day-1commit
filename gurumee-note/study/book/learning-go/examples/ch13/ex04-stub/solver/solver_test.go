package solver

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
)

var errorInvalidExpression = errors.New("invalid expression: ( 2 + 2 * 10")

type MathSolverStub struct{}

func (ms MathSolverStub) Resolve(ctx context.Context, expr string) (float64, error) {
	switch expr {
	case "2 + 2 * 10":
		return 22, nil
	case "( 2 + 2 ) * 10":
		return 40, nil
	case "( 2 + 2 * 10":
		return 0, errorInvalidExpression
	}

	return 0, nil
}

func TestProcessorProcessExpression(t *testing.T) {
	p := Processor{MathSolverStub{}}
	in := strings.NewReader(`2 + 2 * 10
( 2 + 2 ) * 10
( 2 + 2 * 10`)

	data := []float64{22, 40, 0, 0}
	for idx, expected := range data {
		result, err := p.ProcessExpression(context.Background(), in)
		fmt.Println(idx, result, err)
		if err != nil {
			if errors.Is(err, errorInvalidExpression) {
				continue
			}

			t.Error(err)
		}

		if result != expected {
			t.Errorf("%d test - Expected result: %f, got: %f", idx, expected, result)
		}
	}
}
