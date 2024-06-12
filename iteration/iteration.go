package iteration

func Iteration(count int) string {
	word := "p"
	var newWord string
	for i := 0; i < count; i++ {
		newWord += word
	}
	return newWord
}
