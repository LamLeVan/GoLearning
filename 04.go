package main

import (
	"fmt"
	"math"
)

//tinh chu ky dao dong cua con lac don
const GA = 9.18 // accelerate of gravity
const PI float64 = 3.1416

func sqrt(x float64) float64 {
	return math.Sqrt(x / GA)
}

func expression(x float64) float64 {
	return 2 * PI * sqrt(x)
}

func main() {
	fmt.Print(expression(10))
}
