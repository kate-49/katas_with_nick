package main

import (
	"strconv"
)

func fizzBuzz(input int) string {
	if input == 1 {
		return strconv.Itoa(input)
	} else {
		return "Fizz"
	}
}

func main() {}
