package main

import (
	"fmt"
	"math"
)

func sh(x float64) float64 {
	return math.Exp(x)/2 - math.Exp(-x)/2
}

func main() {
	fmt.Print(sh(10))
}
