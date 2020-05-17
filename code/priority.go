package main

import "container/heap"

// An PrioritizedState is something managed in a priority queue.
type PrioritizedState struct {
	state    State
	priority int
	index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*PrioritizedState

// Len will give the number of elements in the queue
func (queue PriorityQueue) Len() int { return len(queue) }

// Checks if ith element is less than jth element
func (queue PriorityQueue) Less(i, j int) bool {
	return queue[i].priority < queue[j].priority
}

// Swaps ith and jth elements
func (queue PriorityQueue) Swap(i, j int) {
	queue[i], queue[j] = queue[j], queue[i]
	queue[i].index = i
	queue[j].index = j
}

// Push method to put an element
func (queue *PriorityQueue) Push(element interface{}) {
	length := len(*queue)
	item := element.(*PrioritizedState)
	item.index = length
	*queue = append(*queue, item)
}

// Pop element from the queue
func (queue *PriorityQueue) Pop() interface{} {
	oldQueue := *queue
	length := len(oldQueue)
	item := oldQueue[length-1]
	oldQueue[length-1] = nil
	item.index = -1
	*queue = oldQueue[0 : length-1]
	return item
}

// NewPriorityQueue will create a new queue for states with priority
func NewPriorityQueue() PriorityQueue {
	queue := make(PriorityQueue, 0)
	heap.Init(&queue)
	return queue
}
