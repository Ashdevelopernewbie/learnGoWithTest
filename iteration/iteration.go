package iteration

func Iteration(count int, word string) string {
	var newWord string
	for i := 0; i < count; i++ {
		newWord += word
	}
	return newWord
}
