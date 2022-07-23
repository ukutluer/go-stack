package main

import (
	"fmt"

	"github.com/ukutluer/go-stack/stack"
)

func main() {
	//stackItem := make(stack.Stack.items[int], 0)

	stack2 := stack.NewStack[string]()
	stack2.Push("1")
	stack2.Push("2")
	stack2.Push("3")
	stack2.Print()
	item, _ := stack2.Peek()
	fmt.Println(item)
	item, _ = stack2.Pop()
	fmt.Println(item)
	stack2.Print()

	item, _ = stack2.Pop()
	fmt.Println(item)
	stack2.Print()

	item, _ = stack2.Pop()
	fmt.Println(item)
	stack2.Print()

	item, _ = stack2.Pop()
	fmt.Println(item)
	stack2.Print()

	fmt.Println(stack2.IsEmpty())
}
