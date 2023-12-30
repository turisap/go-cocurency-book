package main

import (
	"sync"
)

type Queue struct {
	elements    []int
	front, rear int
	len         int
}

// NewQueue initializes an empty circular queue
// with the given capacity
func NewQueue(capacity int) *Queue {
	return &Queue{
		elements: make([]int, capacity),
		front:    0,  // Read from elements[front]
		rear:     -1, // Write to elements[rear]
		len:      0,
	}
}

// Enqueue adds a value to the queue. Returns false
// if queue is full
func (q *Queue) Enqueue(value int) bool {
	if q.len == len(q.elements) {
		return false
	}
	// Advance the write pointer, go around in a circle
	q.rear = (q.rear + 1) % len(q.elements)
	// Write the value
	q.elements[q.rear] = value
	q.len++
	return true
}

func (q *Queue) Dequeue() (int, bool) {
	if q.len == 0 {
		return 0, false
	}
	// Read the value at the read pointer
	data := q.elements[q.front]
	// Advance the read pointer, go around in a circle
	q.front = (q.front + 1) % len(q.elements)
	q.len--
	return data, true
}

func main() {
	lock := sync.Mutex{}
	fullCond := sync.NewCond(&lock)
	emptyCond := sync.NewCond(&lock)
	queue := NewQueue(10)

	producer := func() {
	}

	consumer := func() {
	}

	for i := 0; i < 3; i++ {
		go producer()
	}
	for i := 0; i < 3; i++ {
		go consumer()
	}
	select {} // Wait indefinitely
}
