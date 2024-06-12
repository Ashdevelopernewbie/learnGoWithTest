package main

import (
	"fmt"
	arraysandslices "learnGoWithTests/arraysAndSlices"
)

func main() {
	out := arraysandslices.SumAll([]int{1, 2}, []int{3, 4})
	fmt.Println(out)
}
