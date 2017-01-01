package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		if i < 10 {
			fmt.Printf("%v    ", i)
		} else {
			fmt.Printf("%v   ", i)
		}
		if (i+1)%10 == 0 {
			fmt.Println("")
		}
	}
}
