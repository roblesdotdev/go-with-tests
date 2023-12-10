package main

import "testing"

func TestGreet(t *testing.T) {
	t.Run("Greet with name", func(t *testing.T) {

		got := Greet("Chuck")
		want := "Hello, Chuck!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Greet with empty string", func(t *testing.T) {
		got := Greet("")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)

	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
