package sum

import (
	"math"
	"testing"
)

type testCase struct {
	a, b, sum int64
}

func TestSum(t *testing.T) {
	for _, input := range []testCase{
		{a: 2, b: 2, sum: 4},
		{a: 32, b: 2, sum: 34},
		{a : 15, b: -17, sum: -2},
		{a: math.MaxInt64, b: 1, sum: math.MinInt64},
	} {
		if out := Sum(input.a, input.b); out != input.sum {
			t.Errorf("Wrong calculation of %d + %d", input.a, input.b)
		}
	}
}
