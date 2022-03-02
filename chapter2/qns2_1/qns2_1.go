package qns2_1

type Node struct {
	data int
	next *Node
}

func (head *Node) append(data int) {
	end := Node{data, nil}
	ptr := head
	for ptr.next != nil {
		ptr = ptr.next
	}
	ptr.next = &end
}

func RemoveDups1(head *Node) {
	if head == nil {
		return
	}
	cache := make(map[int]bool)
	cache[head.data] = true
	ptr := head
	for ptr != nil && ptr.next != nil {
		_, exists := cache[ptr.next.data]
		if exists {
			ptr.next = ptr.next.next
			continue
		} else {
			cache[ptr.next.data] = true
		}
		ptr = ptr.next
	}
}

func RemoveDups2(head *Node) {
	curr := head
	for curr != nil {
		prev := curr
		ptr := curr.next
		for ptr != nil {
			if ptr.data == curr.data {
				prev.next = prev.next.next
			}
			prev = ptr
			ptr = ptr.next
		}
		curr = curr.next
	}
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
