package qns2_3

import . "ctci/chapter2"

func DeleteMiddleNode(node *Node) {
	curr := node
	for curr != nil {
		curr.Data = curr.Next.Data
		if curr.Next.Next == nil {
			curr.Next = nil
		}
		curr = curr.Next
	}
}
