package queue

import (
	"testing"
)

func TestNew(t *testing.T) {
	queue := New()
	if queue == nil {
		t.Errorf("New returned unexpected value. Wanted: Queue interface, Got: nil")
		return
	}
	if !queue.IsEmpty() {
		t.Errorf("New returned non-empty queue.")
		return
	}
}

func TestQueueImpl_Enqueue(t *testing.T) {
	queue := New()
	value := "hello"
	queue.Enqueue(value)

	if queue.IsEmpty() {
		t.Errorf("Enqueue failed to insert value into queue.")
		return
	}
	if peek := queue.Peek(); peek != value {
		t.Errorf("Enqueue inserted unexpected value into queue. Wanted: %+v, Got: %+v", value, peek)
		return
	}
}

func TestQueueImpl_Dequeue_EmptyQueue(t *testing.T) {
	queue := New()
	if value := queue.Dequeue(); value != nil {
		t.Errorf("Dequeue returned unexpected value. Wanted: %+v, Got: %+v", nil, value)
		return
	}
}

func TestQueueImpl_Dequeue(t *testing.T) {
	queue := New()

	firstValue := "hello"
	secondValue := "world"

	queue.Enqueue(firstValue)
	queue.Enqueue(secondValue)

	if value := queue.Dequeue(); value != firstValue {
		t.Errorf("Dequeue returned unexpected value. Wanted: %+v, Got: %+v", firstValue, value)
		return
	}

	if value := queue.Dequeue(); value != secondValue {
		t.Errorf("Dequeue returned unexpected value. Wanted: %+v, Got: %+v", secondValue, value)
		return
	}

	if empty := queue.IsEmpty(); !empty {
		t.Errorf("Dequeue changed queue size unexpectedly. Wanted: emoty queue, Got: non-empty queue")
		return
	}
}

func TestQueueImpl_IsEmpty(t *testing.T) {
	queue := New()

	if empty := queue.IsEmpty(); !empty {
		t.Errorf("IsEmpty returned unexpected value. Wanted: %t, Got: %t", true, empty)
		return
	}

	queue.Enqueue("someValue")

	if empty := queue.IsEmpty(); empty {
		t.Errorf("IsEmpty returned unexpected value. Wanted: %t, Got: %t", false, empty)
		return
	}
}

func TestQueueImpl_Peek(t *testing.T) {
	queue := New()

	firstValue := "hello"
	secondValue := "world"

	queue.Enqueue(firstValue)
	if value := queue.Peek(); value != firstValue {
		t.Errorf("Peek returned unexpected value. Wanted: %+v, Got: %+v", firstValue, value)
		return
	}

	queue.Enqueue(secondValue)
	if value := queue.Peek(); value != firstValue {
		t.Errorf("Peek returned unexpected value. Wanted: %+v, Got: %+v", firstValue, value)
		return
	}
}
