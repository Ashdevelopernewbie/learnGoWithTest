package arraysandslices

func Sum(num []int) int {
	var total int
	for _, digits := range num {
		total += digits
	}
	return total
}
