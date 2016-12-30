package main

import (
	"fmt"
	"math"
)

func power(k int) float64 {
	return math.Pow(float64(k), float64(1)/float64(k))
}

func result(k int) float64 {
	var s float64 = 10
	for i := 1; i < k; i++ {
		s += float64(1) / power(i)
	}
	return s
}

func main() {
	fmt.Println(result(10))
}
