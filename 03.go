package main

import (
	"fmt"
	"math"
)

// This func is not handle error, we will continue after then
func power5(ratio float64) float64 {
	return math.Pow(5, ratio)
}

func log(x, y float64) (float64, string) {
	if x <= 0 || y <= 0 {
		return 0, ""
	}
	return math.Log(y) / math.Log(x), ""
}

func log5(y float64) float64 {
	var a, _ = log(5, y)
	return a
}

func absolute(x, y float64) float64 {
	return math.Abs(x - y)
}

func arctg(x, y float64) float64 {
	return math.Atan(x + y)
}

func numeritor(x, y float64) float64 {
	return power5(x) + log5(absolute(x, y))
}

func demominator(x, y float64) float64 {
	return 1 + arctg(x, y)
}

func expression(x, y float64) float64 {
	return numeritor(x, y) / demominator(x, y)
}

func main() {
	fmt.Print(expression(10, 20))
}
