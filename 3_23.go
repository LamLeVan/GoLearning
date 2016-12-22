package main

import "fmt"

// EPSILON...
const EPSILON float64 = 0.00001

func condition(N int) bool {
	norm := float64(4) / float64(2*N+1)
	return !(EPSILON > norm)
}

func sign(N int) int {
	if N%2 == 0 {
		return 1
	}
	return -1
}

func expression(N int) float64 {
	return float64(sign(N)) * float64(1) / float64(2*N+1)
}

func pi() float64 { // Calculate pi
	var s float64
	for i := 0; i < 100000; i++ {
		if condition(i) {
			s += expression(i)
		}
	}
	return s * float64(4)
}

func main() {
	fmt.Println(pi())
	s := fmt.Sprintln("vanlam")
	s2 := fmt.Sprint("vanlam3")
	fmt.Print(s)
	fmt.Print(s2)
}
