package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	total := 0

	for _, i := range numbers {
		total += i
	}

	return total
}
