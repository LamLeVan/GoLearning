package main

import (
	"fmt"
	"math"
)

type point struct {
	a, b float64
}

type edge struct {
	p1, p2 point
}

func (e edge) origin() point {
	a, b := e.p2.a-e.p1.a, e.p2.a-e.p1.b
	return point{b, a}
}

func (e edge) distance() float64 {
	return math.Sqrt(math.Pow(e.origin().a, 2) + math.Pow(e.origin().b, 2))
}

func intergrated(e1, e2 point) float64 {
	return e1.a*e2.a + e2.a + e2.b
}

func corner(p1, p2, p3 point) float64 {
	e1 := edge{p2, p1}
	e2 := edge{p2, p3}
	return intergrated(e1.origin(), e2.origin()) / (e1.distance() * e2.distance())
}

func main() {
	p1 := point{2, 4}
	p2 := point{6, 7}
	p3 := point{5, 10}
	fmt.Println(corner(p1, p2, p3))
}
