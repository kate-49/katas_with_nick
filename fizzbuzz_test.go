package main

import (
	"testing"
)

func checkEquality(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestFizzBuzz(t *testing.T) {
	t.Run("TestIntIfNotFizzbuzz", func(t *testing.T) {
		got := fizzBuzz(1)
		want := "1"
		checkEquality(t, got, want)
	})

	t.Run("TestFizzFor3", func(t *testing.T) {
		got := fizzBuzz(3)
		want := "Fizz"
		checkEquality(t, got, want)
	})

	t.Run("TestFizzFor5", func(t *testing.T) {
		got := fizzBuzz(5)
		want := "Buzz"
		checkEquality(t, got, want)
	})

	t.Run("TestFizzFor15", func(t *testing.T) {
		got := fizzBuzz(15)
		want := "FizzBuzz"
		checkEquality(t, got, want)
	})

	t.Run("TestFizzFor30", func(t *testing.T) {
		got := fizzBuzz(30)
		want := "FizzBuzz"
		checkEquality(t, got, want)
	})
}
