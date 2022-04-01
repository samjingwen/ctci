package main

import "fmt"

func main() {
	var subsets [][]int
	subsets = append(subsets, []int{})

	newSet := make([][]int, len(subsets))
	copy(newSet, subsets)
	fmt.Println(newSet)
}
