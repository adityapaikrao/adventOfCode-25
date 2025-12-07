package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
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

func countSplits(grid [][]rune) int {
	numSplits := 0
	colPos := mapset.NewSet[int]()

	for i, row := range grid[0] {
		if row == 'S' {
			colPos.Add(i)
			break
		}
	}

	for i := 1; i < len(grid); i++ {
		for j, char := range grid[i] {
			if colPos.Contains(j) && char == '^' {
				colPos.Add(j - 1)
				colPos.Add(j + 1)
				numSplits++

				colPos.Remove(j)
			}
		}
	}

	return numSplits
}

func main() {
	grid := parseInput()

	// fmt.Println(grid)
	fmt.Println(countSplits(grid))
}
