package arraysandslices

func Sum(num []int) int {
	var total int
	for _, digits := range num {
		total += digits
	}
	return total
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] += Sum(numbers)
	}
	return sums
}
