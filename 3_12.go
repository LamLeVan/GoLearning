package main

import "fmt"

func factorial(a int) int64 {
	var i int64 = 1
	for j := 1; j <= a; j++ {
		i = i * int64(j)
	}
	return i
}

func result() int64 {
	max := 1000
	for i := 1; i < max; i++ {
		if factorial(i) > int64(max) {
			return int64(i)
		}
	}
	return 0
}

func main() {
	fmt.Println(result())
}
