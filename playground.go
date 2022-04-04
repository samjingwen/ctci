package main

import (
	"fmt"
	"sort"
)

type color struct {
	r, g, b uint8
}

func main() {
	coins := []int{1, 2, 5}

	sort.Sort(SortByDesc(coins))

	fmt.Println(coins)
}

type SortByDesc []int

func (arr SortByDesc) Len() int {
	return len(arr)
}

func (arr SortByDesc) Less(i, j int) bool {
	return arr[i] > arr[j]
}

func (arr SortByDesc) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
