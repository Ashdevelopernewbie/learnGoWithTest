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

func GetAllTails(numbersToSum ...[]int) []int {
	var tails []int
	for _, number := range numbersToSum {
		if len(number) <= 0 {
			tails = append(tails, 0)
		} else {
			tail := number[1:]
			tails = append(tails, Sum(tail))
		}
	}
	return tails
}

func GetAllHeads(numbertosum ...[]int) []int {
	var heads []int
	for _, number := range numbertosum {
		head := number[:1]
		heads = append(heads, Sum(head))
	}
	return heads
}
