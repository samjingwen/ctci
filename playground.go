package main

import "fmt"

func main() {
	const nihongo = "日本語\u0099"
	const stuff = "abc"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}

	for index, runeValue := range stuff {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}

}
