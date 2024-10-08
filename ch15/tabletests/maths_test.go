package maths

import "testing"

func TestDoMath(t *testing.T) {
	data := []struct {
		name     string
		x        int
		y        int
		op       string
		expected int
		errMsg   string
	}{
		{"addition", 2, 2, "+", 4, ""},
		{"subtraction", 2, 2, "-", 0, ""},
		{"multiplication", 2, 2, "*", 4, ""},
		{"division", 2, 2, "/", 1, ""},
		{"bad_division", 2, 0, "/", 0, `division by zero`},
		{"bad_op", 2, 2, "?", 0, `unknown operator ?`},
		{"another_mult", 2, 3, "*", 6, ""},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			actual, err := DoMath(d.x, d.y, d.op)
			if d.expected != actual {
				t.Errorf("expected: %d actual: %d", d.expected, actual)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("expected: `%s` actual: `%s`", d.errMsg, errMsg)
			}
		})
	}
}
