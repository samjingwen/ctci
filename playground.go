package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	ints := arr[:-5]
	fmt.Println(ints)
}
