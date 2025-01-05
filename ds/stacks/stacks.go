package stacks

import "errors"

type Stack[T any] struct {
  elements []T
}

func NewStack[T any]() *Stack[T] {
  return &Stack[T]{
    elements: []T{},
  }
}



func (s *Stack[T]) Push(value T) {
  s.elements = append(s.elements, value)
}

func (s *Stack[T]) Pop() (T, error) {
  if s.IsEmpty() {
    var zeroValue T
    return zeroValue, errors.New("Stack is empty")
  }
  top := s.elements[len(s.elements)-1]
  s.elements = s.elements[:len(s.elements)-1]
  return top, nil
}

func (s *Stack[T]) Size() int {
  return len(s.elements)
}

func (s *Stack[T]) Peek() (T, error) {
  if s.IsEmpty() {
    var zeroValue T
    return zeroValue, errors.New("Stack is empty")
  }
  return s.elements[len(s.elements)-1], nil
}

func (s *Stack[T]) IsEmpty() bool {
  return len(s.elements) == 0
}

func (s *Stack[T]) Empty() {
  s.elements = nil
}



