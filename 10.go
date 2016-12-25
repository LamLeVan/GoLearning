//Tinh ch(x)
package main

import (
	"fmt"
	"math"
)

func ch(x float64) float64 {
	return (math.Exp(x) + math.Exp(-x)) / 2
}

func main() {
	fmt.Print(ch(10))
}
