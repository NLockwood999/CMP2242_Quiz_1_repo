package main

import "fmt"

// Create a struct type named 'triangle' that has two fields named
//
//	base and height both of type float64.
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
	return (t.height + t.height) + t.base
}

// Inside the 'main' function create a variable of type 'triangle'
func main() {
	tri1 := triangle{
		height: 20,
		base:   18,
	}

	tri2 := triangle{
		height: 20,
		base:   18,
	}

	fmt.Println(tri1.Area(), "\n")
	fmt.Println(tri1.Perimeter(), "\n")
	fmt.Println(tri2.Area(), "\n")
	fmt.Println(tri2.Perimeter(), "\n")
}
