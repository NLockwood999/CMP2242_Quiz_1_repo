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
		height: 30,
		base:   15,
	}

	fmt.Println("The Area is: ", tri1.Area())
	fmt.Printf("Using%+v \n", tri1)
	fmt.Println("The Perimeter is: ", tri1.Perimeter())

	fmt.Println("\nThe Area is: ", tri2.Area())
	fmt.Printf("Using%+v \n", tri2)
	fmt.Println("The Perimeter is: ", tri2.Perimeter())
}
