package main

import (
	"fmt"
)

func main() {
	//const nihongo = "日本語\u0099"
	//const stuff = "abc"
	//for index, runeValue := range nihongo {
	//	fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	//}
	//
	//for index, runeValue := range stuff {
	//	fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	//}

	mask := 0b001
	bitVector := 0b100

	b := bitVector & mask

	a := bitVector & ^mask

	c := bitVector | mask

	fmt.Printf("%b\n", b)

	fmt.Printf("%b\n", a)

	fmt.Printf("%b\n", c)

}
