package stubs

import (
	"context"
	"errors"
	"strings"
	"testing"
)

type MathSolverStub struct{}

func (ms MathSolverStub) Resolve(ctx context.Context, expr string) (float64, error) {
	switch expr {
	case "2 + 2 * 10":
		return 22, nil
	case "( 2 + 2 ) * 10":
		return 40, nil
	case "( 2 + 2 * 10":
		return 0, errors.New("invalid expression: ( 2 + 2 * 10")
	}
	return 0, nil
}

func TestProcessorProcessExpression(t *testing.T) {
	p := Processer{MathSolverStub{}}
	in := strings.NewReader("2 + 2 * 10\n( 2 + 2 ) * 10\n( 2 + 2 * 10")
	data := []float64{22, 40, 0}
	hasErr := []bool{false, false, true}
	for i, d := range data {
		result, err := p.ProcessExpression(context.Background(), in)
		if err != nil && !hasErr[i] {
			t.Error(err)
		}
		if result != d {
			t.Errorf("Expected result %f, got %f", d, result)
		}
	}
}
