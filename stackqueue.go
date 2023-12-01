package twostacksqueue

import (
	"container/list"
	"errors"
)

type Stack[T any] struct {
	*list.List
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{list.New()}
}

func (s *Stack[T]) Push(v T) {
	s.PushBack(v)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.Len() == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	e := s.Back()
	s.Remove(e)
	return e.Value.(T), nil
}

type Queue[T any] struct {
	in, out *Stack[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		in:  NewStack[T](),
		out: NewStack[T](),
	}
}

func (q *Queue[T]) Enqueue(v T) {
	q.in.Push(v)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.out.Len() == 0 {
		if q.in.Len() == 0 {
			var zero T
			return zero, errors.New("queue is empty")
		}
		for q.in.Len() > 0 {
			value, _ := q.in.Pop()
			q.out.Push(value)
		}
	}
	return q.out.Pop()
}
