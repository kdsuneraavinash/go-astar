package main

import (
	"fmt"
	"log"
)

func aStar(start State, goal State, heuristicTo HeuristicTo) ([]string, int) {
	// Required Data Structures:
	// openHeapMap: Heap to make the item grabbing easier.
	// 			 	Also a map to keep track of elements that are open.
	// 				This will also keep pointers to open heap items.
	//				This will aid in updating priorities.
	// closedSet:	Set to keep track of closed elements.
	//				openMap keys are closedSet elements are disjoint.
	// memoizedG:	Map to memoize calculated g(x) values.
	// parent:		Keep track of the parent node of each element.
	openHeapMap := NewHeapMap()
	closedSet := NewSet()
	memoizedG := make(map[State]int)
	parent := make(map[State]*State)
	heuristicFrom := heuristicTo(goal)

	// Initial heuristic value is calculated and inserted into all structures.
	// Pointer of the inserted heap element is added to the map.
	fStart := heuristicFrom(start)
	memoizedG[start] = fStart
	openHeapMap.Add(start, fStart)

	iterations := 0
	// Loop until heap is empty
	for !openHeapMap.Empty() {
		iterations++

		// Get the smallest state from queue and move into closed
		state := openHeapMap.Pop().state
		closedSet.Add(&state)

		// Check if goal is reached
		if state == goal {
			break
		}

		// G value when traversed from current state
		gNextState := memoizedG[state] + 1

		// Loop for each possibility
		for _, successor := range state.nextMoves() {
			// Check if successor is already in opened/closed.
			// If they exists in either set, check whether current path is better.
			// If this is a completely new state, add it to the heap.
			existsInOpen := openHeapMap.Exists(successor)
			existsInClosed := closedSet.Exists(&successor)

			if existsInOpen || existsInClosed {
				if gNextState < memoizedG[successor] {
					// Current path is a better path.
					// If this is currently on progress, update the value.
					// Otherwise, re-add to the heap.
					memoizedG[successor] = gNextState
					parent[successor] = &state
					fSuccessor := gNextState + heuristicFrom(successor)
					if existsInOpen {
						openHeapMap.Update(successor, fSuccessor)
					} else {
						closedSet.Remove(&successor)
						openHeapMap.Add(successor, fSuccessor)
					}
				}
			} else {
				// New node found.
				// Insert the element to open set.
				// Also record the path to the element and the node pointer.
				memoizedG[successor] = gNextState
				parent[successor] = &state
				fSuccessor := gNextState + heuristicFrom(successor)
				openHeapMap.Add(successor, fSuccessor)
			}
		}
	}

	goalParent, goalReached := parent[goal]

	// Goal is not reached
	if !goalReached {
		log.Fatalln("Goal is unreachable.")
		return []string{}, -1
	}

	// Trace the actions that took place
	nextNode := goal
	previousNode := goalParent
	previousNodeExists := true
	actions := make([]string, 0)
	for previousNodeExists {
		actionTile, actionDescription := previousNode.detectAction(nextNode)
		action := fmt.Sprintf("(%v, %v)", actionTile, actionDescription)
		actions = append(actions, action)
		nextNode = *previousNode
		previousNode, previousNodeExists = parent[*previousNode]
	}

	// Reverse the actions list
	for i, j := 0, len(actions)-1; i < j; i, j = i+1, j-1 {
		actions[i], actions[j] = actions[j], actions[i]
	}

	return actions, iterations
}
