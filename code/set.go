package main

// Set is a datastructure supporting O(1) add/remove/check
type Set struct {
	data map[State]struct{}
}

// Add element to set
func (set *Set) Add(state *State) {
	set.data[*state] = struct{}{}
}

// Remove element from set
func (set *Set) Remove(state *State) {
	delete(set.data, *state)
}

// Exists checks whether set element exists
func (set *Set) Exists(state *State) bool {
	_, exists := set.data[*state]
	return exists
}

// NewSet will create an empty set
func NewSet() Set {
	return Set{data: make(map[State]struct{})}
}
