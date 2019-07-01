package main

import (
	"fmt"
	"math"
	"time"
)

func pi(n int) float64 {
	N := float64(n)
	res := make(chan float64)
	nth := 8.
	for i := 0.; i < nth; i += 1. {
		go func(id float64) {
			sum := 0.
			for i := id; i < N; i += nth {
				x := (i + .5) / N
				sum += 4. / (1. + x*x)
			}
			res <- sum
		}(i)
	}
	sum := 0.
	for i := 0.; i < nth; i += 1. {
		sum += <-res
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
