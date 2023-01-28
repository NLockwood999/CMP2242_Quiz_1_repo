package main

import "fmt"

/*
Write a Go function named 'square' that accepts a parameter of type float64

	that represents the side of the triangle. The 'square' function should return
	the area and the perimeter of the traingle when called.
*/
func square(z float64) (x, y float64) {

	x = (z * z) //area of square
	y = (z * 4) //perimeter of square

	return x, y
}

func main() {

	area, perimeter := square(5)
	fmt.Println(area, perimeter)
}
