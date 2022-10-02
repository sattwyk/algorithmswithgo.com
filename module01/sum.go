package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	num := 0

	for _, i := range numbers {
		num += i
	}

	return num
}
