package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	maxValue := 10
	result := make([]int, 0, maxValue)
	for i := 0; i < maxValue; i++ {
		if i%2 == 0 {
			result = append(result, i)
		}
	}
	for i := range result {
		fmt.Printf("%d: %v\n", i, &result[i])
	}
	newSlice := result[1:3]
	newSlice2 := result[2:4]
	fmt.Printf("[:]: %s\n", getSliceHeader(&result))
	fmt.Printf("[1:3]: %s\n", getSliceHeader(&newSlice))
	fmt.Printf("[2:4]: %s\n", getSliceHeader(&newSlice2))
}

func getSliceHeader(slice *[]int) string {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(slice))
	return fmt.Sprintf("%+v", sh)
}
