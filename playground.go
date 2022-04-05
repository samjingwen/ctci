package main

import (
	"fmt"
)

type color struct {
	r, g, b uint8
}

func main() {
	emptyBoard := make([][]rune, 8)
	for i := 0; i < 8; i++ {
		emptyBoard[i] = make([]rune, 8)
	}

	fmt.Println(emptyBoard)
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
