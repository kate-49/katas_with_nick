package main

import "testing"

func checkEquality(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestIntIfNotFizzbuzz(t *testing.T) {
	got := fizzBuzz(1)
	want := "1"
	checkEquality(t, got, want)
}

func TestFizzFor3(t *testing.T) {
	got := fizzBuzz(3)
	want := "Fizz"
	checkEquality(t, got, want)
}

func TestFizzFor5(t *testing.T) {
	got := fizzBuzz(5)
	want := "Buzz"
	checkEquality(t, got, want)
}
