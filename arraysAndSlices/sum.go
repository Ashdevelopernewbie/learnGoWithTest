package arraysandslices

func Sum(num [7]int) int {
	var total int
	for _, digits := range num {
		total += digits
	}
	return total
}
