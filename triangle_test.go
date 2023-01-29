package main

import "testing"

func TestTArea(t *testing.T) {
	got := triangle{height: 20, base: 10}
	want := got.Area()
	if want != 100 {
		t.Errorf("got %v as area, wanted %v as area", got, want)
	}
}

func TestTPerimeter(t *testing.T) {
	got := triangle{height: 20, base: 10}
	want := got.Perimeter()
	if want != 52.3606797749979 {
		t.Errorf("got %v as area, wanted %v as area", got, want)
	}
}
