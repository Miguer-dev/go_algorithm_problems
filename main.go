package main

import (
	"AlgorithmProblems/binary_search"
	"fmt"
)

func main() {
	fmt.Print(binary_search.SearchMatrix([][]int{[]int{1, 3, 5, 7}, []int{10, 11, 16, 20}, []int{23, 30, 34, 60}}, 79))
}
