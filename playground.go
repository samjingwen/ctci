package main

import "fmt"

type Stack []int

type BST struct {
	node *Node
}

type Node struct {
	data int
}

func main() {

	queue := make(chan *Node)
	queue <- &Node{1}
	curr := <-queue
	fmt.Println(*curr)
}

func doStuff(node *Node) {
	*node = Node{1}
}
