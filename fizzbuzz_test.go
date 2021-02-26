package main

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	FizzBuzzProvider := []struct {
			Got int
			Want string
	}{
		{Got: 1, Want: "1"},
		{Got: 3, Want: "Fizz"},
		{Got: 5, Want: "Buzz"},
		{Got: 15, Want: "FizzBuzz"},
	}

	for _, FizzBuzzInput := range FizzBuzzProvider {
		got := fizzBuzz(FizzBuzzInput.Got)
		if FizzBuzzInput.Want != got {
			t.Errorf("got %q want %q", FizzBuzzInput.Got, FizzBuzzInput.Want)
		}
	}
}
