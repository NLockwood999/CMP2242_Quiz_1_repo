package main

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	got := Greeting()
	expected := "Hello, world!"
	//comparison between got and expected
	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}
