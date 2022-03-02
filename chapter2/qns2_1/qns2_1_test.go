package qns2_1

import "testing"

type args struct {
	head *Node
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"", args{nil}, nil},
	{"", args{makeNodeFromSlice([]int{1})}, []int{1}},
	{"", args{makeNodeFromSlice([]int{1, 1})}, []int{1}},
	{"", args{makeNodeFromSlice([]int{1, 1, 1})}, []int{1}},
	{"", args{makeNodeFromSlice([]int{1, 1, 1, 1})}, []int{1}},
	{"", args{makeNodeFromSlice([]int{1, 2, 3, 4, 5})}, []int{1, 2, 3, 4, 5}},
	{"", args{makeNodeFromSlice([]int{1, 1, 3, 4, 5})}, []int{1, 3, 4, 5}},
	{"", args{makeNodeFromSlice([]int{1, 2, 1, 4, 5})}, []int{1, 2, 4, 5}},
	{"", args{makeNodeFromSlice([]int{1, 2, 3, 4, 4})}, []int{1, 2, 3, 4}},
	{"", args{makeNodeFromSlice([]int{1, 1, 3, 4, 5, 1})}, []int{1, 3, 4, 5}},
	{"", args{makeNodeFromSlice([]int{1, 1, 3, 4, 5, 4})}, []int{1, 3, 4, 5}},
	{"", args{makeNodeFromSlice([]int{1, 2, 3, 4, 5, 5, 5})}, []int{1, 2, 3, 4, 5}},
}

func TestRemoveDups1(t *testing.T) {
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			if RemoveDups1(test.args.head); !test.args.head.equals(test.want) {
				t.Fatalf("RemoveDups1() = %v, want %v", makeSliceFromNode(test.args.head), test.want)
			}
		})
	}
}

func TestRemoveDups2(t *testing.T) {
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			if RemoveDups2(test.args.head); !test.args.head.equals(test.want) {
				t.Fatalf("RemoveDups2() = %v, want %v", makeSliceFromNode(test.args.head), test.want)
			}
		})
	}
}
