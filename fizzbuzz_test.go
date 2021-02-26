package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	got := fizzBuzz(1)
	want := 1

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
