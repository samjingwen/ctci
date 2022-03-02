package main

import "fmt"

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

	//mask := 0b001
	//bitVector := 0b100
	//
	//b := bitVector & mask
	//
	//a := bitVector & ^mask
	//
	//c := bitVector | mask
	//
	//fmt.Printf("%b\n", b)
	//
	//fmt.Printf("%b\n", a)
	//
	//fmt.Printf("%b\n", c)

	head := makeNodeFromSlice([]int{1, 2, 3, 4, 5})
	arr := makeSliceFromNode(head)
	fmt.Println(arr)

	fmt.Println(head.equals([]int{1, 2, 3, 4, 5}))
	fmt.Println(head.equals([]int{1, 2, 3, 5}))
	fmt.Println(head.equals([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(head.equals([]int{}))

	head.append(6)
	fmt.Println(makeSliceFromNode(head))
}

type Node struct {
	data int
	next *Node
}

func makeNodeFromSlice(arr []int) *Node {
	head := Node{0, nil}
	ptr := &head
	for _, val := range arr {
		curr := Node{val, nil}
		ptr.next = &curr
		ptr = ptr.next
	}
	return head.next
}

func makeSliceFromNode(head *Node) []int {
	var arr []int
	ptr := head
	for ptr != nil {
		arr = append(arr, ptr.data)
		ptr = ptr.next
	}
	return arr
}

func (head *Node) equals(arr []int) bool {
	ptr := head
	for _, val := range arr {
		if ptr == nil || ptr.data != val {
			return false
		}
		ptr = ptr.next
	}
	if ptr != nil {
		return false
	}
	return true
}

func (head *Node) append(data int) {
	end := Node{data, nil}
	ptr := head
	for ptr.next != nil {
		ptr = ptr.next
	}
	ptr.next = &end
}
