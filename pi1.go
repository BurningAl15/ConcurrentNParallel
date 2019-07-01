package main

import (
	"fmt"
	"math"
	"time"
)

func pi(n int) float64 {
	N := float64(n)

	sum := 0.

	for i := 0.; i < N; i += 1. {
		x := (i + .5) / N
		sum += 4. / (1. + x*x)
	}
	return sum / N
}

func main() {
	fmt.Println(math.Pi)
	ini := time.Now()
	p := pi(100000000)
	fin := time.Since(ini)
	fmt.Println(p, fin)
}
