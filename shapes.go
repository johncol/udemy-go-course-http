package main

import "fmt"

// HasArea x
type HasArea interface {
	area() float64
}

// NewSquare x
func NewSquare(sideLength float64) Square {
	return Square{
		sideLength: sideLength,
	}
}

// Square x
type Square struct {
	sideLength float64
}

func (square Square) area() float64 {
	return square.sideLength * square.sideLength
}

// NewTriangle x
func NewTriangle(height, base float64) Triangle {
	return Triangle{
		height: height,
		base:  base,
	}
}

// Triangle x
type Triangle struct {
	height float64
	base float64
}

func (triangle Triangle) area() float64 {
	return 0.5 * triangle.base * triangle.height
}

func printArea(thingWithArea HasArea) {
	fmt.Println("Area of shape is", thingWithArea.area())
}
