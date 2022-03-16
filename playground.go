package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	var leftNodeSize int
	if leftNodeSize = 2; false {
		leftNodeSize = 3
	}
	fmt.Println(leftNodeSize)
}

func getSliceHeader(slice *[]int) string {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(slice))
	return fmt.Sprintf("%+v", sh)
}
