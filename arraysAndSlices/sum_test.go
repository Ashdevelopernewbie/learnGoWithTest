package arraysandslices

import (
	"reflect"
	"testing"
)

func assertCorrectMessage(t testing.TB, got, expected []int) {

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("\nexpected 		%d\ngot			%d", expected, got)
	} else {
		t.Logf("\nexpected 		%d\ngot			%d", expected, got)
	}
}

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

	assertCorrectMessage(t, got, expected)
}

func TestGetAllTails(t *testing.T) {
	assertMessage := func(t testing.TB, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("\nexpected 		%d\ngot			%d", expected, got)
		} else {
			t.Logf("\nexpected 		%d\ngot			%d", expected, got)
		}
	}
	t.Run("get the tail of all slices", func(t *testing.T) {
		got := GetAllTails([]int{1, 2}, []int{3, 4})
		expected := []int{2, 4}

		assertMessage(t, got, expected)
	})
	t.Run("get the tail of all slices even if empty", func(t *testing.T) {
		got := GetAllTails([]int{}, []int{1, 2}, []int{3, 4})
		expected := []int{0, 2, 4}

		assertMessage(t, got, expected)
	})
}

func TestGetAllHeads(t *testing.T) {
	got := GetAllHeads([]int{1, 5, 7, 2, 2}, []int{3, 8, 4})
	expected := []int{1, 3}

	assertCorrectMessage(t, got, expected)
}
