package main

import "testing"

func TestIntIfNotFizzbuzz(t *testing.T) {
	got := fizzBuzz(1)
	want := "1"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestFizzFor3(t *testing.T) {
	got := fizzBuzz(3)
	want := "Fizz"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestFizzFor5(t *testing.T) {
	got := fizzBuzz(5)
	want := "Buzz"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
