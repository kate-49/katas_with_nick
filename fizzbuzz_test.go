package main

import "testing"

func returnIntIfNotFizzbuzz(t *testing.T) {
	got := fizzBuzz(1)
	want := 1

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func returnFizzFor3(t *testing.T) {
	got := fizzBuzz(3)
	want := "Fizz"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
