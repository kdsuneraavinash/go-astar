package main

// HeuristicFrom Calculate distance from start state
type HeuristicFrom func(start State) int

// HeuristicTo Heuristic function to distance calculation between given end state
type HeuristicTo func(end State) HeuristicFrom

// Pair is a tuple of 2 numbers
type Pair struct {
	x, y int
}

// ManhattanHeuristic function for
// distance calculation using manhattan distance
// Splitting the function increased performance by 2x
func ManhattanHeuristic(end State) HeuristicFrom {
	points := make(map[int]Pair)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			ind := index(i, j)
			points[end[ind]] = Pair{i, j}
		}
	}

	return func(start State) int {
		totalDist := 0
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				ind := index(i, j)
				if !start.empty(ind) {
					a := Pair{i, j}
					b := points[start[ind]]
					dist := abs(a.x-b.x) + abs(a.y-b.y)
					totalDist += dist
				}
			}
		}
		return totalDist
	}
}

// MisplacedHeuristic function for
// distance calculation using misplaced tiles
func MisplacedHeuristic(end State) HeuristicFrom {
	return func(start State) int {
		misplaced := 0
		for i := 0; i < N*N; i++ {
			if start[i] != end[i] {
				misplaced++
			}
		}
		return misplaced
	}
}
