package arraysandslices

func Sum(num []int) int {
	var total int
	for _, digits := range num {
		total += digits
	}
	return total
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}
