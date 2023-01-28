package main

import "log" 
"math"

//Create a struct type named 'triangle' that has two fields named
//  base and height both of type float64.
type triangle struct {
	base float64
	height float64
}
//Create a method on type 'triangle' named 'area' that calculates and 
//returns the area of a triangle.
func (s triangle) Areatriangle() (float64, float64) {
return (s.height * s.base)/2
}

funct main() {
	
}