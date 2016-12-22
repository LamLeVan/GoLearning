package main

import (
	"fmt"
	"math"
)

func expression(x float64, n int) float64 {
	return math.Pow(x, float64(n)) / float64(n) * float64(sign(n))
}

func sign(n int) int {
	if n%2 == 0 {
		return 1
	}
	return -1
}

func sum(x float64, n int) float64 {
	var s float64
	for i := 1; i <= n; i++ {
		s += expression(x, i)
	}
	return s
}

func main() {
	fmt.Println(sum(2, 3))
}
