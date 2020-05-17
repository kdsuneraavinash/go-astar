# Golang A* Algorithm Implementation for Solving Modified N-Puzzle

## Problem

Suppose the standard n-puzzle problem is modified so that there are two empty locations (instead of one) allowing you to pick any one tile adjacent to any one of the two empty locations and move it to the adjacent empty location. You are given a starting configuration and are supposed to rearrange the tiles by move them into the empty locations as necessary to end up in a given goal configuration.

Following shows an example of a starting and a goal configuration.

Start:

```text
1	4	-	7
9	2	3	5
6	-	10	13
8	11	14	12
```

Goal:

```text
1	4	7	5
9	2	3	-
-	11	10	13
6	8	14	12
```

Sample Output:

```text
(7,left), (11, up), (8,right), (6, down), (5, up)
```

Your program should be able to accept as command line input, two tab delimited files of the form provided together with this assignment as sample starting configuration and goal configuration. The output should be a text file containing a sequence of moves of the form [(tile_number, move)].

## Implementation

This is an implementation of the basic heuristic search using a* to solve the problem. This implementation is done using golang.
