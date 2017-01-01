package main

import "fmt"

func swap(x *int, y *int) {
	var temp int
	temp = *x //Save the value at address
	*x = *y   // Put y into x
	*y = temp // Put temp into y
}

func main() {
	//Local variable definition
	var a int = 100
	var b int = 200
	swap(&a, &b)
	fmt.Printf("Before swap, value of a :%d\n", a)
	fmt.Printf("Before swap, value of b: %d\n", b)

	/*calling a function to swap the values
	 * &a indicates pointer to a ie. address of variable a and
	 * &b in
	 */
}
