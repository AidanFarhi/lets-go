package stubs

import (
	"context"
	"io"
)

type Processer struct {
	Solver MathSolver
}

func (p Processer) ProcessExpression(ctx context.Context, r io.Reader) (float64, error) {
	expression, err := readToNewLine(r)
	if len(expression) == 0 {
		return 0, err
	}
	answer, err := p.Solver.Resolve(ctx, expression)
	return answer, err
}

func readToNewLine(r io.Reader) (string, error) {
	var out []byte
	b := make([]byte, 1)
	for {
		_, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				return string(out), nil
			}
		}
		if b[0] == '\n' {
			break
		}
		out = append(out, b[0])
	}
	return string(out), nil
}

type MathSolver interface {
	Resolve(ctx context.Context, expression string) (float64, error)
}
