package main

import "fmt"

const Pi = 3.14

func main() {
	const Flag = true
	fmt.Printf("%T(%v), %T(%v)", Pi, Pi, Flag, Flag)
}
