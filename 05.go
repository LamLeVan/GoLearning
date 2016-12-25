//Nhap 3 canh cua tam giac tinh S = (p(p-a)(p-b)(p-c))1/2
//voi P= (a+b+c)/2

package main

import (
	"fmt"
	"math"
)

func isTriangle(x, y, z float64) bool {
	return x+y > z || x+z > y || y+z > x
}

func perimeterHalf(x, y, z float64) float64 {
	return (x + y + z) / 2
}

// func substraction(p func(x, y, z), subs float64) float64 {
// 	return p - subs
// }
func substraction(p, subs float64) float64 {
	return p - subs
}

func intergrated(x, y, z float64) float64 {
	var p = perimeterHalf(x, y, z)
	return p * substraction(p, x) * substraction(p, y) * substraction(p, z)
}

func expression(x, y, z float64) float64 {
	return math.Sqrt(intergrated(x, y, z))
}

func main() {
	var x, y, z float64 = 3, 4, 5
	if isTriangle(x, y, z) {
		fmt.Print(expression(x, y, z))
	} else {
		fmt.Print("Error")
	}
}
