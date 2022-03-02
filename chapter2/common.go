package chapter2

type Node struct {
	Data int
	Next *Node
}

func (head *Node) Append(data int) {
	end := Node{data, nil}
	ptr := head
	for ptr.Next != nil {
		ptr = ptr.Next
	}
	ptr.Next = &end
}

func (head *Node) Equals(arr []int) bool {
	ptr := head
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

func MakeNodeFromSlice(arr []int) *Node {
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
