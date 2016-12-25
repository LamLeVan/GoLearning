package main

import (
	"fmt"
	"math"
)

func power2(x float64) float64 {
	return x * x
}

func expression(x float64) float64 {
	return 1 + math.Log(power2(x))
}

func main() {
	fmt.Print(expression(10))
}
