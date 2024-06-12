package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{5, 9, 4, 3, 7, 0, 5}

		got := Sum(numbers)
		expected := 33

		if got != expected {
			t.Errorf("expected %d got %d, given %v", expected, got, numbers)
		} else {
			t.Logf("expected %d got %d, given %v", expected, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4})
	expected := []int{3, 7}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %d got %d", expected, got)
	} else {
		t.Logf("expected %d got %d", expected, got)
	}

}
