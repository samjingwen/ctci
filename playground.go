package main

import (
	"fmt"
)

type color struct {
	r, g, b uint8
}

func main() {
	visited := make([][]color, 5)
	fmt.Println(visited)
	for i, _ := range visited {
		visited[i] = make([]color, 4)
	}
	fmt.Println(visited)
}
