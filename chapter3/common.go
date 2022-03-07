package chapter3

import "errors"

type Stack struct {
	values []int
}

func NewStack(arr ...int) *Stack {
	return &Stack{arr}
}

func (stack *Stack) Push(val int) {
	stack.values = append(stack.values, val)
}

func (stack *Stack) Pop() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	val := stack.values[stack.Size()-1]
	stack.values = stack.values[:stack.Size()-1]
	return val, nil
}

func (stack *Stack) Peek() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	return stack.values[stack.Size()-1], nil
}

func (stack *Stack) IsEmpty() bool {
	return stack.Size() == 0
}

func (stack *Stack) Size() int {
	return len(stack.values)
}

func (stack *Stack) Values() []int {
	return stack.values
}

type Queue []int

func (queue *Queue) Enqueue(val int) {
	*queue = append(*queue, val)
}

func (queue *Queue) Deque() (int, error) {
	if len(*queue) == 0 {
		return 0, errors.New("queue is empty")
	}
	val := (*queue)[0]
	*queue = (*queue)[1:]
	return val, nil
}

func (queue *Queue) Peek() (int, error) {
	if len(*queue) == 0 {
		return 0, errors.New("queue is empty")
	}
	return (*queue)[0], nil
}

func (queue *Queue) IsEmpty() bool {
	return len(*queue) == 0
}
