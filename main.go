package main

import (
	"fmt"

	"github.com/ukutluer/go-stack/stack"
)

func main() {
	s := stack.NewStack[string]()
	s.Push("1")
	s.Push("2")
	s.Push("3")
	s.Print()
	item, _ := s.Peek()
	fmt.Println(item)
	item, _ = s.Pop()
	fmt.Println(item)
	s.Print()

	item, _ = s.Pop()
	fmt.Println(item)
	s.Print()

	item, _ = s.Pop()
	fmt.Println(item)
	s.Print()

	item, _ = s.Pop()
	fmt.Println(item)
	s.Print()

	fmt.Println(s.IsEmpty())
}
