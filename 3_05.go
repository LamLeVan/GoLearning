package main

import "fmt"

func star(a int) {
	for i := 0; i < a; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func dollar(a int) {
	for i := 0; i < a; i++ {
		fmt.Print("$")
	}
	fmt.Println()
}

func starResult() {
	for i := 0; i < 10; i++ {
		star(i)
	}
}

func dollarResult() {
	for i := 10; i > 0; i-- {
		dollar(i)
	}
}

func starTriangle() {
	high := 10
	for i := 0; i < high; i++ {
		for j := high - 1; j > i; j-- {
			print(" ")
		}
		for j := 0; j < i*2+1; j++ {
			print("*")
		}
		println()
	}
}

func main() {
	starResult()
	fmt.Println()
	dollarResult()
	fmt.Println()
	starTriangle()
}
