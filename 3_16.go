package main

import "fmt"

func isPrime(a int) bool {
	for i := 2; i < a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func result() {
	for i := 1; i < 1000; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}

func main() {
	result()
}
