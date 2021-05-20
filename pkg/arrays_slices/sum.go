package arrays_slices

// Sum takes in a slice of integers and returns the sum of the integers.
func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}
