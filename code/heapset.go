package main

import "container/heap"

// HeapMap which will encapsulate a PriorityQueue and a Map
type HeapMap struct {
	heap PriorityQueue
	set  map[State]*PrioritizedState
}

// NewHeapMap will create empty heap map
func NewHeapMap() *HeapMap {
	heap := NewPriorityQueue()
	set := make(map[State]*PrioritizedState)
	return &HeapMap{heap: heap, set: set}
}

// Add will add an element to the heap with given priority
func (heapMap *HeapMap) Add(state State, priority int) {
	prioritizedStart := &PrioritizedState{state: state, priority: priority}
	heap.Push(&heapMap.heap, prioritizedStart)
	heapMap.set[state] = prioritizedStart
}

// Update will update and heapify the heap
func (heapMap *HeapMap) Update(state State, priority int) {
	prioritizedState := heapMap.set[state]
	prioritizedState.priority = priority
	heap.Fix(&heapMap.heap, prioritizedState.index)
}

// Pop will give the smallest element in heap
func (heapMap *HeapMap) Pop() *PrioritizedState {
	prioritizedState := heap.Pop(&heapMap.heap).(*PrioritizedState)
	delete(heapMap.set, prioritizedState.state)
	return prioritizedState
}

// Exists will check if state exists
func (heapMap *HeapMap) Exists(state State) bool {
	_, exists := heapMap.set[state]
	return exists
}

// Empty will check if heap is empty
func (heapMap *HeapMap) Empty() bool {
	return len(heapMap.heap) == 0
}
