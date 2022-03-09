package chapter4

import (
	"reflect"
	"testing"
)

func TestBST(t *testing.T) {
	bst := BST{}

	bst.Add(5)
	bst.Add(4)
	bst.Add(7)
	bst.Add(8)

	depth := bst.Depth()
	if depth != 3 {
		t.Errorf("depth = %v, expected = %v", depth, 3)
	}

	arr := bst.LevelOrderTraversal()
	if expected := []int{5, 4, 7, 8}; !reflect.DeepEqual(arr, expected) {
		t.Errorf("levelOrderTraversal = %v, expected = %v", arr, expected)
	}

}
