package main

import "testing"

func TestCArea(t *testing.T) {

	got := circle{radius: 8}
	want := got.Area()
	if want != 201.06192982974676 {
		t.Errorf("got %v as area, wanted %v as area", got, want)
	}
}

func TestCPerimeter(t *testing.T) {

	got := circle{radius: 8}
	want := got.Perimeter()
	if want != 50.26548245743669 {
		t.Errorf("got %v as perimeter, wanted %v as perimeter", got, want)
	}
}
