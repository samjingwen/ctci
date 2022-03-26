package main

import (
	"fmt"
)

func main() {
	var x uint32 = 0xaaaaaaaa
	var y uint32 = 0x55555555
	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", y)
	fmt.Printf("%b\n", 0xa)
	fmt.Printf("%b\n", 5)
}

func getBit(num uint32, i int) bool {
	return (num & (1 << i)) != 0
}
