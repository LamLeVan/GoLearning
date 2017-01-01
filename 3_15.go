package main

import (
	"fmt"
	"math"
)

func numbers(a int) []int {
	array := []int{}
	for a > 0 {
		array = append(array, a%10)
		a = a / 10
	}
	return array
}

func power3(a int) int {
	return int(math.Pow(float64(a), 3))
}

func isResult(a int) bool {
	s := 0
	array := numbers(a)
	for _, value := range array {
		s += power3(value)
	}
	return s == a
}

func main() {
	fmt.Println(numbers(371))
	for i := 1; i < 1000; i++ {
		if isResult(i) {
			fmt.Println(i)
		}
	}
}
