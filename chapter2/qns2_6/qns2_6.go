package qns2_6

import . "ctci/chapter2"

func Palindrome(head *Node) bool {
	ptr := head

	var iter func(*Node) bool
	iter = func(node *Node) bool {
		if node == nil {
			return true
		}

		res := iter(node.Next) && ptr.Data == node.Data
		ptr = ptr.Next
		return res
	}
	return iter(ptr)
}
