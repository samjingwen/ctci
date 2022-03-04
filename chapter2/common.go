package chapter2

import (
	"fmt"
	"strings"
)

type Node struct {
	Data int
	Next *Node
}

func (node *Node) String() string {
	ptr := node
	var builder strings.Builder
	fmt.Fprintf(&builder, "[")
	for ptr != nil {
		fmt.Fprintf(&builder, "%v", ptr.Data)
		if ptr.Next != nil {
			fmt.Fprintf(&builder, ", ")
		}
		ptr = ptr.Next
	}
	fmt.Fprintf(&builder, "]")
	return builder.String()
}

func (node *Node) Append(data int) {
	end := Node{data, nil}
	ptr := node
	for ptr.Next != nil {
		ptr = ptr.Next
	}
	ptr.Next = &end
}

func (node *Node) Equals(arr []int) bool {
	ptr := node
	for _, val := range arr {
		if ptr == nil || ptr.Data != val {
			return false
		}
		ptr = ptr.Next
	}
	if ptr != nil {
		return false
	}
	return true
}

func MakeNodeFromSlice(arr ...int) *Node {
	head := Node{0, nil}
	ptr := &head
	for _, val := range arr {
		curr := Node{val, nil}
		ptr.Next = &curr
		ptr = ptr.Next
	}
	return head.Next
}

func MakeSliceFromNode(head *Node) []int {
	var arr []int
	ptr := head
	for ptr != nil {
		arr = append(arr, ptr.Data)
		ptr = ptr.Next
	}
	return arr
}
