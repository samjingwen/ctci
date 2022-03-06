package chapter3

import "errors"

type Stack []int

func (stack *Stack) Append(val int) {
	*stack = append(*stack, val)
}

func (stack *Stack) Pop() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	val := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return val, nil
}

func (stack *Stack) Peek() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	return (*stack)[len(*stack)-1], nil
}

func (stack *Stack) IsEmpty() bool {
	return len(*stack) == 0
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
