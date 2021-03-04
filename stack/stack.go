package stack

import "sync"

// Stack represents a basic interface that implements the usual stack methods.
type Stack interface {
	// Push pushes to the top of the stack.
	Push(elem interface{})
	// Pop removes the most recent element of the stack and returns it.
	// Returns nil if the stack is empty.
	Pop() interface{}
	// Peek returns the most recent element of the stack.
	// Returns nil if the stack is empty.
	Peek() interface{}
	// Size returns the count of elements in the stack.
	Size() int
}

type stackImpl struct {
	mutex *sync.RWMutex
	data  []interface{}
}

// New returns a thread-safe implementation of Stack.
func New() Stack {
	return &stackImpl{
		mutex: &sync.RWMutex{},
		data:  []interface{}{},
	}
}

func (s *stackImpl) Push(elem interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.data = append(s.data, elem)
}

func (s *stackImpl) Pop() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.data) == 0 {
		return nil
	}

	elem := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return elem
}

func (s *stackImpl) Peek() interface{} {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if len(s.data) == 0 {
		return nil
	}
	return s.data[len(s.data)-1]
}

func (s *stackImpl) Size() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return len(s.data)
}
