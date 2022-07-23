package stack

import (
	"testing"
)

func TestStackPeek(t *testing.T) {
	s := NewStack[string]()
	s.Push("1")
	s.Push("2")
	s.Push("3")
	expectedValue := "3"
	actualValue, _ := s.Peek()
	if actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestStackPop(t *testing.T) {
	s := NewStack[string]()
	s.Push("1")
	s.Push("2")
	s.Push("3")
	expectedValue := "3"
	actualValue, _ := s.Pop()
	if actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestStackPush(t *testing.T) {
	s := NewStack[string]()
	s.Push("1")

	expectedValue := "1"
	actualValue, _ := s.Peek()
	if actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestStackIsEmpty(t *testing.T) {
	s := NewStack[string]()
	s.Push("1")
	expectedValue := false
	actualValue := s.IsEmpty()

	if actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}
