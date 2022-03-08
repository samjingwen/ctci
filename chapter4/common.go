package chapter4

import "errors"

type Queue []interface{}

func (q *Queue) Enqueue(val interface{}) {
	*q = append(*q, val)
}

func (q *Queue) Deque() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	val := (*q)[0]
	*q = (*q)[1:]
	return val, nil
}

func (q *Queue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	return (*q)[0], nil
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
