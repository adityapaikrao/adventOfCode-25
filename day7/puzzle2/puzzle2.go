package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseInput() [][]rune {
	grid := make([][]rune, 0)

	path := "../puzzle1.in"
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)

		row := make([]rune, 0)
		for _, char := range text {
			row = append(row, char)
		}

		grid = append(grid, row)
	}

	return grid
}

func runBeam(i int, j int, grid [][]rune, memo map[[2]int]int) int {
	// Base Cases
	i++

	if i >= len(grid) {
		return 1 // found one valid path
	}

	if j < 0 || j >= len(grid[i]) {
		return 0
	}

	gridPos := [2]int{i, j}
	if val, exists := memo[gridPos]; exists {
		return val
	}
	numTimelines := 0

	switch grid[i][j] {
	case '^':
		leftTimelines := runBeam(i, j-1, grid, memo)
		rightTimelines := runBeam(i, j+1, grid, memo)
		numTimelines = leftTimelines + rightTimelines

	default:
		numTimelines = runBeam(i, j, grid, memo)
	}

	memo[gridPos] = numTimelines

	return numTimelines

}

func countTimelines(grid [][]rune) int {
	start_i := 0
	start_j := 0

	for i, row := range grid {
		for j, char := range row {
			if char == 'S' {
				start_i = i
				start_j = j
				break
			}
		}
	}

	memo := make(map[[2]int]int)
	return runBeam(start_i, start_j, grid, memo)
	// runBeam(start_i, start_j, grid, &memo) ->
	//  dont need to pass in as reference maps & slices are already passed by reference

}

func main() {
	grid := parseInput()

	fmt.Println(countTimelines(grid))
}
