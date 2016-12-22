package main

import "fmt"

func numbers(a int) []int {
	array := []int{}
	for a > 0 {
		array = append(array, a%10)
		a = a / 10
	}
	return array
}

func calculate(array []int) (int, int) {
	s := 0
	i := 1
	for _, value := range array {
		s += value
		i *= value
	}
	return s, i
}

func result(a int) {
	b := numbers(a)
	fmt.Println(calculate(b))
}

func main() {
	result(125356)
}
