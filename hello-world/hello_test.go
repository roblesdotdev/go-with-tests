package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Chuck")
	want := "Hello, Chuck!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
