package main

import "testing"

func TestHello(t *testing.T) {
	// what's the diff b/w t.Run and not using it
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func assertCorrectMessage(t testing.TB, got, want string) { //shorten b/c 2 args have the same type
	t.Helper() // tells go to fail at function call and not a specfic line
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
