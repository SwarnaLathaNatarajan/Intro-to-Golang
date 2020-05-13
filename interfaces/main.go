package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Square struct {
	h, w float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (s Square) area() float64 {
	return s.h * s.w
}

func getArea(s Shape) float64 {
	return s.area()
}
func main() {
	c, s := Circle{2}, Square{2, 3}
	fmt.Println(getArea(c), getArea(s))
}
