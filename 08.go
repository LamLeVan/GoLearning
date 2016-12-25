package main

import (
	"fmt"
	"math"
)

func pow2(x float64) float64 {
	return x * x
}

func expression(x float64) float64 {
	return pow2(x) - math.Sin(x)
}

func main() {
	fmt.Print(expression(10))
}
