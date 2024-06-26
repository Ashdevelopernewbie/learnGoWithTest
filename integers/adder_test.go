package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d, got %d", expected, sum)
	} else {
		t.Logf("expected %d, got %d", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(4, 5)
	fmt.Println(sum)
	// Output: 9
}
