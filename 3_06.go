package main

import "fmt"

// so tuyet hao: 6 = 3 + 1 + 2
func divisor(a int) []int { // E
	array := []int{}
	for i := 1; i < a; i++ {
		if a%i == 0 {
			array = append(array, i)
		}
	}
	return array
}

func sum(a []int) int {
	sum := 0
	for _, value := range a {
		sum += value
	}
	return sum
}

func isPerfectNumber(a int) bool {
	array := divisor(a)
	return a == sum(array)
}

func result() {
	for i := 1; i <= 10000; i++ {
		if isPerfectNumber(i) {
			fmt.Println(i)
		}
	}
}

func main() {
	result()
}
