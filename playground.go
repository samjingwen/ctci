package main

import (
	"fmt"
)

func main() {

	fmt.Printf("%b\n", 1<<2)
	fmt.Printf("%v\n", getBit(0b1010101, 2))
}

func getBit(num uint32, i int) bool {
	return (num & (1 << i)) != 0
}
