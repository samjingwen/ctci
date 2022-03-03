package qns2_4

import . "ctci/chapter2"

// 3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1
// 3 -> 2 -> 8 -> 5 -> 10 -> 5 -> 1 swap(1, 5)
// 3 -> 2 -> 1 -> 5 -> 10 -> 5 -> 8 swap(2, 6)

func Partition1(head *Node, partition int) *Node {
	ptr := head
	index := -1
	for ptr != nil {
		if ptr.Data < partition {
			index++
			swap(ptr, head, index)
		}
		ptr = ptr.Next
	}
	return head
}

func Partition2(head *Node, partition int) *Node {
	ptr := head
	var beforeStart *Node
	var beforeEnd *Node
	var afterStart *Node
	var afterEnd *Node

	for ptr != nil {
		next := ptr.Next
		ptr.Next = nil

		if ptr.Data < partition {
			if beforeStart == nil {
				beforeStart = ptr
				beforeEnd = ptr
			} else {
				beforeEnd.Next = ptr
				beforeEnd = beforeEnd.Next
			}
		} else {
			if afterStart == nil {
				afterStart = ptr
				afterEnd = ptr
			} else {
				afterEnd.Next = ptr
				afterEnd = afterEnd.Next
			}
		}
		ptr = next
	}

	beforeEnd.Next = afterStart
	return beforeStart
}

func Partition3(head *Node, partition int) *Node {
	start := head
	end := head
	ptr := head

	for ptr != nil {
		next := ptr.Next
		if ptr.Data < partition {
			ptr.Next = start
			start = ptr
		} else {
			end.Next = ptr
			end = ptr
		}
		ptr = next
	}
	end.Next = nil
	return start
}

func swap(node *Node, head *Node, index int) {
	ptr := head
	count := 0
	for ptr != nil && count < index {
		ptr = ptr.Next
		count++
	}
	ptr.Data, node.Data = node.Data, ptr.Data
}
