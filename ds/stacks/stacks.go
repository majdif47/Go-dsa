package stacks

import "errors"

type Stack struct {
  elements []int
}

func NewStack() *Stack {
  return &Stack{
    elements: []int{},
  }
}



func (s *Stack) Push(x int) {
  s.elements = append(s.elements, x)
}

func (s *Stack) Pop() (int, error) {
  if s.IsEmpty() {
    return 0, errors.New("Stack is empty")
  }
  top := s.elements[len(s.elements)-1]
  s.elements = s.elements[:len(s.elements)-1]
  return top, nil
}

func (s *Stack) Size() int {
  return len(s.elements)
}

func (s *Stack) Peek() (int, error) {
  if s.IsEmpty() {
    return 0, errors.New("Stack is empty")
  }
  return s.elements[len(s.elements)-1], nil
}

func (s *Stack) IsEmpty() bool {
  return len(s.elements) == 0
}

func (s *Stack) Empty() {
  s.elements = nil
}



