package ex_rapid

import (
	"fmt"
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

func TestCustomGen(t *testing.T) {
	type point struct {
		x int
		y int
	}

	gen := rapid.Custom(func(t *rapid.T) point {
		return point{
			x: rapid.Int().Draw(t, "x").(int),
			y: rapid.Int().Draw(t, "y").(int),
		}
	})

	rapid.Check(t, func(t *rapid.T) {
		v := gen.Draw(t, "v").(point)
		fmt.Printf("%#v\n", v)
	})
}
