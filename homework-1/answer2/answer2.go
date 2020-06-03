package main

import (
	"fmt"
	"math"
)

var (
	a float64
	b float64
)

func main() {

	a, b = 2, 3
	c := math.Sqrt(a*a + b*b)
	p := a + b + c
	s := a * b / 2
	fmt.Printf("гипотенуза - %.1f, периметр - %.1f, площадь - %.1f\n", c, p, s)
}
	