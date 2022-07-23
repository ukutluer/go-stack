package stack

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Peek() (T, error) {
	if !s.IsEmpty() {
		return s.items[len(s.items)-1], nil
	}
	var result T
	return result, errors.New("THERE IS NO ITEM IN STACK")
}

func (s *Stack[T]) Pop() (T, error) {
	if !s.IsEmpty() {
		willBePoppedItem := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
		return willBePoppedItem, nil
	}
	var result T
	return result, errors.New("THERE IS NO ITEM IN STACK")

}

func (s *Stack[T]) IsEmpty() bool {
	return !(len(s.items) > 0)
}

func (s *Stack[T]) Print() {
	fmt.Printf("Stack has %d items.", len(s.items))
	fmt.Println(s.items)
}
func NewStack[T any]() *Stack[T] {
	stack := &Stack[T]{}
	return stack
}
