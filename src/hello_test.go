package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(got, want, t)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(got, want, t)
	})
	t.Run("say hello in Spanish", func(t *testing.T) {
		got := Hello("Rosie", "Spanish")
		want := "Hola, Rosie"
		assertCorrectMessage(got, want, t)
	})
	t.Run("say hello in French", func(t *testing.T) {
		got := Hello("Rosie", "French")
		want := "Bonjour, Rosie"
		assertCorrectMessage(got, want, t)
	})
}

func assertCorrectMessage(got string, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}