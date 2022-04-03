package main

import (
	"fmt"
)

func main() {
	//var builder strings.Builder

	str := "az⌘cbza"
	runeList := []rune(str)

	fmt.Printf("%v\n", len(str))
	fmt.Printf("%v\n", len(runeList))

	for i, ch := range []rune(str) {
		fmt.Printf("%v, %+q\n", i, ch)

	}
	fmt.Printf("%v\n", str[5:])

}
