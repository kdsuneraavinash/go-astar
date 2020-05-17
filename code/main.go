package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// N denotes the number of cells on one side
// in the grid. (defaults to 4)
var N int = 5

// MaxN denotes the highest possible N value
// MaxN*MaxN will also be the size of a state
const MaxN int = 20

// Detect the rows/columns in a grid
func detectN(file string) int {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	data := string(buffer)
	firstLine := splitStrings(data, "\n")[0]
	cells := len(splitStrings(firstLine, "\t"))
	return cells
}

func writeToFile(fileName string, content []string) {
	// Write the solution to the file
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	for i := 0; i < len(content); i++ {
		_, err = file.WriteString(fmt.Sprintf("%v ", content[i]))
		if err != nil {
			log.Fatalln(err)
		}
	}
	file.Close()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	getInput := func(prompt string) string {
		fmt.Print(prompt)
		content, _ := reader.ReadString('\n')
		return strings.TrimSpace(content)
	}

	argStartFile := getInput("Start configuration file name: ")
	argGoalFile := getInput("Goal configuration file name: ")
	argOutFile := getInput("Output file name: ")
	argsUseSwapYn := getInput("Use misplaced tiles heuristic (y/N): ")
	argsUseSwap := strings.ToLower(argsUseSwapYn) == "y"

	N = detectN(argStartFile)
	if N > MaxN {
		log.Fatalf("N must not exceed %v.\n", MaxN)
	}
	start := StateFromFile(argStartFile)
	goal := StateFromFile(argGoalFile)
	fmt.Printf("N is set to %v.\n", N)
	strategy := ManhattanHeuristic
	if argsUseSwap {
		strategy = MisplacedHeuristic
		fmt.Println("Using misplaced tiles heuristic.")
	} else {
		fmt.Println("Using manhattan distance heuristic.")
	}
	solution, iterations := aStar(start, goal, strategy)
	fmt.Printf("%v iterations took place in finding the goal.\n", iterations)
	fmt.Printf("Solution path consists of %v steps.\n", len(solution))
	writeToFile(argOutFile, solution)
}
