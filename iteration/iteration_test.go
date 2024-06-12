package iteration

import "testing"

func TestIteration(t *testing.T) {
	//Let's write a test for a function that repeats a character 5 times.
	// Comment #2 I did it the opposiste instead of repeating any character 5 times. I repeated the same word 5 times statically. Now change that
	got := Iteration(5)
	expected := "ppppp"

	if got != expected {
		t.Errorf("\nexpected 	%v \ngot 		%v", expected, got)
	} else {
		t.Logf("\nexpected 	%v \ngot 		%v", expected, got)
	}
}

func BenchmarkIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iteration(5)
	}
}