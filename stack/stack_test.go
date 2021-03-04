package stack

import "testing"

func TestNew(t *testing.T) {
	stack := New()
	if stack == nil {
		t.Errorf("New returned unexpected value. Wanted: Stack interface, Got: nil")
		return
	}
	if stack.Size() != 0 {
		t.Errorf("New returned stack of unexpected size. Wanted size: 0, Got: %d", stack.Size())
		return
	}
}

func TestStackImpl_Push(t *testing.T) {
	stack := New()
	stack.Push("hello")

	if stack.Size() != 1 {
		t.Errorf("Push changed stack size unexpectedly. Wanted size: 1, Got: %d", stack.Size())
		return
	}
}

func TestStackImpl_Pop_EmptyStack(t *testing.T) {
	stack := New()
	popped := stack.Pop()
	if popped != nil {
		t.Errorf("Pop returned unexpected value. Wanted: %+v, Got: %+v", nil, popped)
		return
	}
}

func TestStackImpl_Pop(t *testing.T) {
	var popped interface{}
	firstValue := "hello"
	secondValue := "world"
	stack := New()

	stack.Push(firstValue)
	stack.Push(secondValue)

	popped = stack.Pop()
	if popped != secondValue {
		t.Errorf("Pop returned unexpected value. Wanted: %+v, Got: %+v", secondValue, popped)
		return
	}
	if stack.Size() != 1 {
		t.Errorf("Pop resulted in unexpected Stack size. Wanted: %d, Got: %d", 1, stack.Size())
		return
	}

	popped = stack.Pop()
	if popped != firstValue {
		t.Errorf("Pop returned unexpected value. Wanted: %+v, Got: %+v", firstValue, popped)
		return
	}
	if stack.Size() != 0 {
		t.Errorf("Pop resulted in unexpected Stack size. Wanted: %d, Got: %d", 0, stack.Size())
		return
	}
}

func TestStackImpl_Peek_EmptyStack(t *testing.T) {
	stack := New()
	peek := stack.Peek()
	if peek != nil {
		t.Errorf("Peek returned unexpected value. Wanted: %+v, Got: %+v", nil, peek)
		return
	}
}

func TestStackImpl_Peek(t *testing.T) {
	firstValue := "hello"
	secondValue := "world"
	stack := New()

	stack.Push(firstValue)
	peek := stack.Peek()
	if peek != firstValue {
		t.Errorf("Peek returned unexpected value. Wanted: %+v, Got: %+v", firstValue, peek)
		return
	}

	stack.Push(secondValue)
	peek = stack.Peek()
	if peek != secondValue {
		t.Errorf("Peek returned unexpected value. Wanted: %+v, Got: %+v", secondValue, peek)
		return
	}
}

func TestStackImpl_Size(t *testing.T) {
	firstValue := "hello"
	secondValue := "world"
	stack := New()

	size := stack.Size()
	if size != 0 {
		t.Errorf("Size returned unexpected size. Wanted: %+v, Got: %+v", 0, size)
		return
	}

	stack.Push(firstValue)
	size = stack.Size()
	if size != 1 {
		t.Errorf("Size returned unexpected size. Wanted: %+v, Got: %+v", 1, size)
		return
	}

	stack.Push(secondValue)
	size = stack.Size()
	if size != 2 {
		t.Errorf("Size returned unexpected size. Wanted: %+v, Got: %+v", 2, size)
		return
	}
}
