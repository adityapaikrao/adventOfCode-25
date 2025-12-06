package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseInput() ([][]rune, error) {
	path := "../puzzle1.in"

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		grid = append(grid, []rune(text))
	}

	return grid, nil
}

func solveProblem2(grid [][]rune) (int, error) {
	if len(grid) == 0 {
		return 0, nil
	}

	// Find the maximum row length
	maxLen := 0
	for _, row := range grid {
		if len(row) > maxLen {
			maxLen = len(row)
		}
	}

	// Pad all rows to the same length
	for i := range grid {
		for len(grid[i]) < maxLen {
			grid[i] = append(grid[i], ' ')
		}
	}

	totalSum := 0
	numRows := len(grid)
	operatorRow := numRows - 1

	// Process columns from right to left
	col := maxLen - 1
	for col >= 0 {
		// Skip any space-only columns (separators between problems)
		for col >= 0 {
			isAllSpace := true
			for r := 0; r < numRows; r++ {
				if grid[r][col] != ' ' {
					isAllSpace = false
					break
				}
			}
			if !isAllSpace {
				break
			}
			col--
		}
		if col < 0 {
			break
		}

		// Found a non-space column - this is the start of a problem group
		// First, find the operator by scanning the operator row for * or +
		var operator rune = ' '
		scanCol := col
		for scanCol >= 0 {
			ch := grid[operatorRow][scanCol]
			if ch == '*' || ch == '+' {
				operator = ch
				break
			}
			// Check if we've hit a separator column
			isAllSpace := true
			for r := 0; r < numRows; r++ {
				if grid[r][scanCol] != ' ' {
					isAllSpace = false
					break
				}
			}
			if isAllSpace {
				break
			}
			scanCol--
		}

		// Collect all numbers for this problem
		// Each column with digits forms one number (reading top to bottom)
		// Continue until we hit a column that is all spaces (in the data rows)
		var numbers []int

		for col >= 0 {
			// Check if this column is a separator (all spaces in data rows AND operator row)
			isAllSpace := true
			for r := 0; r < numRows; r++ {
				if grid[r][col] != ' ' {
					isAllSpace = false
					break
				}
			}
			if isAllSpace {
				break
			}

			// Check if this column has any digits (not just operator)
			hasDigits := false
			for r := 0; r < operatorRow; r++ {
				ch := grid[r][col]
				if ch >= '0' && ch <= '9' {
					hasDigits = true
					break
				}
			}

			if hasDigits {
				// This column forms a number
				// Read top to bottom:
				// - Leading spaces (before first digit) are ignored
				// - Spaces in the middle (between digits) become 0s
				// - Trailing spaces (after last digit) are ignored

				// First, find the last digit position
				lastDigitRow := -1
				for r := operatorRow - 1; r >= 0; r-- {
					ch := grid[r][col]
					if ch >= '0' && ch <= '9' {
						lastDigitRow = r
						break
					}
				}

				numStr := ""
				seenDigit := false
				for r := 0; r < operatorRow; r++ {
					ch := grid[r][col]
					if ch >= '0' && ch <= '9' {
						numStr += string(ch)
						seenDigit = true
					} else if ch == ' ' && seenDigit && r < lastDigitRow {
						// Space in the middle (after first digit, before last digit) - becomes 0
						numStr += "0"
					}
					// Leading spaces and trailing spaces are ignored
				}
				if numStr != "" {
					num, err := strconv.Atoi(numStr)
					if err != nil {
						return 0, fmt.Errorf("cannot convert %s to int: %w", numStr, err)
					}
					numbers = append(numbers, num)
				}
			}
			col--
		}

		// Calculate result for this problem
		if len(numbers) > 0 {
			result := 0
			switch operator {
			case '*':
				result = 1
				for _, n := range numbers {
					result *= n
				}
			case '+':
				for _, n := range numbers {
					result += n
				}
			}

			totalSum += result
		}
	}

	return totalSum, nil
}

func main() {
	grid, err := parseInput()
	if err != nil {
		panic(fmt.Errorf("could not read file: %v", err))
	}

	result, err := solveProblem2(grid)
	if err != nil {
		panic(err)
	}
	fmt.Println("Total:", result)
}
