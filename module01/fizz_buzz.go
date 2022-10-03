package module01

import (
	"fmt"
	"strconv"
	"strings"
)

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.

func IsFizzBuzz(n int) string {
	str := ""

	if n%3 == 0 {
		str += "Fizz "
	}

	if n%5 == 0 {
		str += "Buzz "
	}

	if len(str) > 0 {
		return strings.TrimRight(str, " ")
	}

	return strconv.Itoa(n)
}

func FizzBuzz(n int) {

	for i := 1; i < n; i++ {
		fmt.Print(IsFizzBuzz(i), ", ")
	}

	fmt.Print(IsFizzBuzz(n), "\n")

}
