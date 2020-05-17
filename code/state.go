package main

import (
	"io/ioutil"
	"log"
	"strconv"
)

// State struct which will hold the state tuple.
// State tuple is an immutable flattened grid with
// 0 for empty positions.
type State [MaxN * MaxN]int

// StateFromFile creates state object from the
// content of the file of the given file name
func StateFromFile(file string) State {
	// Read file content as a string
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	data := string(buffer)

	// Expand data tuple
	state := State{}
	for i, row := range splitStrings(data, "\n") {
		for j, cell := range splitStrings(row, "\t") {
			number, err := strconv.Atoi(cell)
			if err != nil {
				// Assign 0 if reading failed
				number = 0
			}
			state[index(i, j)] = number
		}
	}

	return state
}

// Function to check if index is empty
func (state *State) empty(index int) bool {
	return state[index] == 0
}

// swap function to swap indices of the given state.
// Will return a new state.
func (state State) swap(ai, aj, bi, bj int) State {
	aInd := index(ai, aj)
	bInd := index(bi, bj)
	if !state.empty(aInd) {
		log.Fatal("Swap position was already occupied.")
	}
	state[aInd], state[bInd] = state[bInd], state[aInd]
	return state
}

// nextMoves is a helper function to generate next moves.
// This will create a channel and use emitNextMoves.
func (state *State) nextMoves() []State {
	nextStates := make([]State, 0)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			ind := index(i, j)
			if state.empty(ind) {
				if i > 0 {
					nextStates = append(nextStates, state.swap(i, j, i-1, j))
				}
				if j > 0 {
					nextStates = append(nextStates, state.swap(i, j, i, j-1))
				}
				if i < N-1 {
					nextStates = append(nextStates, state.swap(i, j, i+1, j))
				}
				if j < N-1 {
					nextStates = append(nextStates, state.swap(i, j, i, j+1))
				}
			}
		}
	}
	return nextStates
}

// Detect the action required to go from current state to the next
func (state *State) detectAction(nextState State) (int, string) {
	var emptyPosition, secondPosition = -1, -1
	for i := 0; i < N*N; i++ {
		if state[i] != nextState[i] {
			if state.empty(i) {
				emptyPosition = i
			} else {
				secondPosition = i
			}
		}
	}
	if emptyPosition == -1 || secondPosition == -1 {
		log.Fatal("Invalid state action detected.")
	}
	ai, aj := position(emptyPosition)
	bi, bj := position(secondPosition)

	var direction string
	switch {
	case ai == bi-1 && aj == bj:
		direction = "up"
	case ai == bi+1 && aj == bj:
		direction = "down"
	case ai == bi && aj == bj-1:
		direction = "left"
	case ai == bi && aj == bj+1:
		direction = "right"
	}
	return state[secondPosition], direction
}
