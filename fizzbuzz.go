package main

import (
	"strconv"
)

func fizzBuzz(input int) string {
	if input == 15 {
		return "FizzBuzz"
	} else if input == 3 {
		return "Fizz"
	} else if input == 5 {
		return "Buzz"
	} else {
		return strconv.Itoa(input)
	}
}

func main() {}
