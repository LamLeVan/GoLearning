package main

import "fmt"

type myInt struct {
	a int
}

func (a myInt) toString() string {
	m := map[int]string{
		1:  "One",
		2:  "Two",
		3:  "Three",
		4:  "Four",
		5:  "Five",
		6:  "Six",
		7:  "Seven",
		8:  "Eight",
		9:  "Nine",
		10: "Ten",
	}
	if a.a <= 10 {
		return m[a.a]
	}
	return "invalid"
}

func main() {
	a := myInt{5}
	fmt.Print(a.toString())
}
