package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var m int32 = -2147483648
	var n int32 = -2147483647

	var x int32 = -0b1111111_11111111_11111111_11111111
	var y int32 = -0b10000000_00000000_00000000_00000000

	fmt.Printf("m: %x, %v\n", *(*[4]byte)(unsafe.Pointer(&m)), m)
	fmt.Printf("n: %x, %v\n", *(*[4]byte)(unsafe.Pointer(&n)), n)
	fmt.Printf("x: %b, %v\n", x, x)
	fmt.Printf("y: %b, %v\n", y, y)

	var aa int32 = -1
	fmt.Printf("%b\n", uint32(aa))

	fmt.Printf("%b\n", uint32(n))
}
