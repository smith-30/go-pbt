package ex_rapid

import (
	"math"
	"testing"

	"pgregory.net/rapid"
)

func TestSqrt(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		fgen := rapid.Float64Range(1, math.MaxFloat64)
		v := fgen.Draw(t, "v").(float64)
		if !(math.Sqrt(v) >= 1) {
			t.Fatalf("%v\n", math.Sqrt(v))
		}
	})
}
