package main

import "strings"

// Absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Get index of the flattned grid
func index(i int, j int) int {
	return i*N + j
}

// Get row and column index of a flattned index
func position(n int) (int, int) {
	return n / N, n % N
}

// String trimming function
func splitStrings(str string, with string) []string {
	return strings.Split(strings.TrimSpace(str), with)
}
