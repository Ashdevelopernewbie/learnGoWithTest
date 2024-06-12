package arraysandslices

func Sum(num [7]int) int {
	var total int
	for i := 0; i < len(num); i++ {
		total += num[i]
	}
	return total
}
