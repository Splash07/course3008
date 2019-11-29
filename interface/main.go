package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	Width float64
	Height float64
}
func (r Rect) Area() float64 {
	return 100
}

func (r Rec) Perimeter() float64 {
	return 100
}



type Circle struct {
	Diameter float64
}
func (c Circle) Area() float64 {
	return 200
}



func PrintArea (s Shape) {
	fmt.Println(s.Area())
}

func main {
	
}