package main

import "testing"

func TestHello(t *testing.T) {
	// use t run to run subtests
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("E", "Spanish")
		want := "Hola E"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("E", "French")
		want := "Bonjour E"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) { //shorten b/c 2 args have the same type
	t.Helper() // tells go to fail at function call and not a specfic line
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
