package main

import "fmt"

// Create a struct type named 'circle' that has one field named 'radius' of type
// float64.
type circle struct {
	radius float64
}

// Create a method on type 'circle' named 'area' that calculates and
// returns the area of a circle.
func (c circle) Area() float64 {
	a := 3.14
	return (c.radius * c.radius) * float64(a)
}

// Create a method on type 'triangle' named 'perimeter' that calculates and
// returns the perimeter of a triangle.
func (c circle) Perimeter() float64 {
	b := 3.14
	return (2 * float64(b)) * c.radius
}

// Inside the main function create a variable of type circle
func main() {
	cir1 := circle{
		radius: 15,
	}

	cir2 := circle{
		radius: 30,
	}

	fmt.Println("The Area is: ", cir1.Area())
	fmt.Printf("Using%+v \n", cir1)
	fmt.Println("The Perimeter is: ", cir1.Perimeter())

	fmt.Println("\nThe Area is: ", cir2.Area())
	fmt.Printf("Using%+v \n", cir2)
	fmt.Println("The Perimeter is: ", cir2.Perimeter())
}
