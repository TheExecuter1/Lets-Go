package main

import "testing"



func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := strin("Chris","")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := strin("","")
		want := "Hello, fuz"
		assertCorrectMessage(t, got, want)
	})

	t.Run("checking spanish lang type", func (t *testing.T){
		got := strin("Fuz", "Espanol")
		want := "hola, Fuz"
		assertCorrectMessage(t, got, want)
	})

	t.Run("checking french lang", func (t *testing.T){
		got := strin("Fuz", "French")
		want := "bonj, Fuz"
		assertCorrectMessage(t, got, want)
	})



}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

