package main

import (
	"fmt"

	"pgregory.net/rapid"
)

func main() {
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

	for i := 0; i < 5; i++ {
		fmt.Println(gen.Example(i))
	}
}
