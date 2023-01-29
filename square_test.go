package main

import (
	"testing"
)

func TestASquare(t *testing.T) {
	area, peri := square(8)
	want1, want2 := float64(64), float64(32)

	if area != want1 {
		t.Errorf("got %v as area, wanted %v as area", area, want1)
	}
	if peri != want2 {
		t.Errorf("got %v as perimeter, wanted %v as perimeter", peri, want2)
	}
}
