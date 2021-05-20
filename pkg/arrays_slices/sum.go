package arrays_slices

// Sum takes in a slice of integers and returns the sum of the integers.
func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

// SumAll takes a varying number of slices, returns a new slice containing the totals for each slice passed in
func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

//SumAllTails calculates the totals of the "tails" of each slice.
// The tail of a collection is all the items apart from the first one (the "head")
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) != 0 {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		} else {
			sums = append(sums, 0)
		}
	}

	return sums
}
