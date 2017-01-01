package main

import "fmt"

type triangle struct {
	a, b, c float64
}

func (tr triangle) isTriangle() bool {
	if tr.a > 0 && tr.b > 0 && tr.c > 0 {
		if tr.a+tr.b > tr.c && tr.a+tr.c > tr.b && tr.b+tr.c > tr.a {
			return true
		}
	}
	return false
}

func (tr triangle) isSosceles() bool {
	if tr.isTriangle() {
		return tr.a == tr.b || tr.b == tr.c || tr.c == tr.a
	}
	return false
}

func (tr triangle) isEquilateral() bool {
	if tr.isTriangle() {
		return tr.a == tr.b && tr.a == tr.c
	}
	return false
}

func (tr triangle) isQuarate() bool {
	if tr.isTriangle() {
		if tr.a*tr.a+tr.b*tr.b == tr.c*tr.c {
			return true
		}
		if tr.a*tr.a+tr.c*tr.c == tr.b*tr.b {
			return true
		}
		if tr.b*tr.b+tr.c*tr.c == tr.a*tr.a {
			return true
		}
	}
	return false
}

func (tr triangle) toString() string {
	if tr.isEquilateral() {
		return "Tam giac deu"
	}
	if tr.isSosceles() {
		return "Tam giac can"
	}
	if tr.isQuarate() {
		return "Tam giac vuong"
	}
	if tr.isTriangle() {
		return "Tam giac thuong"
	}
	return "Khong phai tam giac"
}

func main() {
	tr := triangle{3, 3, 10}
	fmt.Println(tr.toString())
}
