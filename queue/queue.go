package queue

import "sync"

// Queue represents a basic interface that implements the usual queue methods.
type Queue interface {
	// Enqueue pushes the given element at the end of the queue.
	Enqueue(elem interface{})
	// Dequeue removes the element from the head of the queue and returns it.
	// Returns nil if queue has no elements.
	Dequeue() interface{}
	// IsEmpty returns true if the size of the queue is 0, otherwise returns false.
	IsEmpty() bool
	// Peek returns the element at the head of the queue.
	// Returns nil if queue has no elements.
	Peek() interface{}
}

type queueImpl struct {
	mutex *sync.RWMutex
	data  []interface{}
}

// New returns a thread-safe implementation of Queue.
func New() Queue {
	return &queueImpl{
		mutex: &sync.RWMutex{},
		data:  []interface{}{},
	}
}

func (q *queueImpl) Enqueue(elem interface{}) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.data = append(q.data, elem)
}

func (q *queueImpl) Dequeue() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if len(q.data) == 0 {
		return nil
	}

	first := q.data[0]
	q.data = q.data[1:]
	return first
}

func (q *queueImpl) IsEmpty() bool {
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	return len(q.data) == 0
}

func (q *queueImpl) Peek() interface{} {
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	if len(q.data) == 0 {
		return nil
	}
	return q.data[0]
}
