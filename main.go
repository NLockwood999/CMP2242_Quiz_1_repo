package main

import (
	"fmt"
	"math"
)

// Create a struct type named 'triangle' that has two fields named
// base and height both of type float64.
type triangle struct {
	base   float64
	height float64
}

// Create a method on type 'triangle' named 'area' that calculates and
// returns the area of a triangle.
func (t triangle) Area() float64 {
	return (t.height * t.base) / 2
}

// Create a method on type 'triangle' named 'perimeter' that calculates and
// returns the perimeter of a triangle.
func (t triangle) Perimeter() float64 {
	h := math.Sqrt(math.Pow(t.base, 2) + math.Pow(t.height, 2))
	return t.base + t.height + h
}

// Create a struct type named 'circle' that has one field named 'radius' of type float64.
type circle struct {
	radius float64
}

// Create a method on type 'circle' named 'area' that calculates and
// returns the area of a circle.
func (c circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// Create a method on type 'circle' named 'perimeter' that calculates and
// returns the area of a circle.
func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Write a Go function named 'square' that accepts a parameter of type float64
// that represents the side of the triangle. The 'square' function should return
// the area and the perimeter of the traingle when called.
func square(z float64) (float64, float64) {
	area := math.Pow(z, 2)
	perimeter := z * 4
	return area, perimeter
}

// Inside the main function create a variable of type triangle
func main() {
	tri1 := triangle{
		base:   10,
		height: 20,
	}
	cir1 := circle{
		radius: 8,
	}

	fmt.Println("The Area of the tri1 obj is: ", tri1.Area())
	fmt.Println("The Perimeter of the tri1 obj is: ", tri1.Perimeter())

	fmt.Println("The Area of the cir1 obj is: ", cir1.Area())
	fmt.Println("The Perimeter of the cir1 obj is: ", cir1.Perimeter())

	ans1, ans2 := square(8)
	fmt.Println("The results of square function are: ", "A=", ans1, "P=", ans2)

}
