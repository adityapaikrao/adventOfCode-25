package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseIntListToken(token string) ([]int, error) {
	clean := strings.TrimSpace(token)
	clean = strings.Trim(clean, "(){}")
	if clean == "" {
		return []int{}, nil
	}
	parts := strings.Split(clean, ",")
	res := make([]int, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		n, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("invalid int %q in token %q: %w", p, token, err)
		}
		res = append(res, n)
	}
	return res, nil
}

func parseInputFromReader(r io.Reader) ([][]int, [][][]int, [][]int, error) {
	target := make([][]int, 0)
	buttons := make([][][]int, 0)
	joltages := make([][]int, 0)

	scanner := bufio.NewScanner(r)
	// Allow longer lines than the default 64K.
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}
		textList := strings.Fields(text)
		if len(textList) < 2 {
			return nil, nil, nil, fmt.Errorf("invalid input line (expected >= 2 fields): %q", text)
		}

		currTarget := make([]int, 0, len(textList[0]))
		for _, char := range textList[0] {
			switch char {
			case '#':
				currTarget = append(currTarget, 1)
			case '.':
				currTarget = append(currTarget, 0)
			}
		}
		target = append(target, currTarget)

		buttonRow := make([][]int, 0, len(textList)-2)
		for i := 1; i < len(textList)-1; i++ {
			currButton, err := parseIntListToken(textList[i])
			if err != nil {
				return nil, nil, nil, err
			}
			buttonRow = append(buttonRow, currButton)
		}
		buttons = append(buttons, buttonRow)

		currJoltage, err := parseIntListToken(textList[len(textList)-1])
		if err != nil {
			return nil, nil, nil, err
		}
		joltages = append(joltages, currJoltage)
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, nil, err
	}

	return target, buttons, joltages, nil
}

func parseInput() ([][]int, [][][]int, [][]int, error) {
	filePath := "../puzzle.in"
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, nil, err
	}
	defer file.Close()

	return parseInputFromReader(file)
}

func processMachine(joltage []int, buttonsRow [][]int) int {
	numCounters := len(joltage)
	numButtons := len(buttonsRow)

	// Build a matrix: buttonMatrix[i][j] = 1 if button i affects counter j, else 0
	buttonMatrix := make([][]int, numButtons)
	for i := range numButtons {
		buttonMatrix[i] = make([]int, numCounters)
		for _, idx := range buttonsRow[i] {
			if idx < numCounters {
				buttonMatrix[i][idx] = 1
			}
		}
	}

	// Sort buttons by number of counters they affect (descending) - press buttons that affect more counters first
	// This helps find good solutions faster
	buttonOrder := make([]int, numButtons)
	for i := range numButtons {
		buttonOrder[i] = i
	}

	// Count how many counters each button affects
	buttonCounts := make([]int, numButtons)
	for i := range numButtons {
		for j := range numCounters {
			buttonCounts[i] += buttonMatrix[i][j]
		}
	}

	// Sort by count descending (buttons affecting more counters first)
	for i := 0; i < numButtons-1; i++ {
		for j := i + 1; j < numButtons; j++ {
			if buttonCounts[buttonOrder[j]] > buttonCounts[buttonOrder[i]] {
				buttonOrder[i], buttonOrder[j] = buttonOrder[j], buttonOrder[i]
			}
		}
	}

	bestResult := math.MaxInt

	// Calculate a lower bound: for each counter, we need at least ceil(joltage[j] / maxCoverage) presses
	// where maxCoverage is the max buttons that can affect a counter
	lowerBound := 0
	for j := range numCounters {
		maxCoverage := 0
		for i := range numButtons {
			if buttonMatrix[i][j] > maxCoverage {
				maxCoverage = buttonMatrix[i][j]
			}
		}
		if maxCoverage > 0 {
			lb := (joltage[j] + maxCoverage - 1) / maxCoverage
			if lb > lowerBound {
				lowerBound = lb
			}
		}
	}

	var solve func(orderIdx int, remaining []int, totalPresses int)
	solve = func(orderIdx int, remaining []int, totalPresses int) {
		// Pruning: if we've already exceeded best, stop
		if totalPresses >= bestResult {
			return
		}

		// Check if all constraints are satisfied
		allZero := true
		hasNegative := false
		maxRemaining := 0
		for _, r := range remaining {
			if r < 0 {
				hasNegative = true
				break
			}
			if r != 0 {
				allZero = false
				if r > maxRemaining {
					maxRemaining = r
				}
			}
		}

		if hasNegative {
			return // Invalid state - overshot
		}

		if allZero {
			if totalPresses < bestResult {
				bestResult = totalPresses
			}
			return
		}

		// Better lower bound pruning
		if totalPresses+maxRemaining >= bestResult {
			return
		}

		// If we've used all buttons but haven't reached target, fail
		if orderIdx >= numButtons {
			return
		}

		buttonIdx := buttonOrder[orderIdx]

		// Calculate maximum times we can press this button without exceeding any counter
		maxPresses := math.MaxInt
		affectsAny := false
		for j := range numCounters {
			if buttonMatrix[buttonIdx][j] > 0 {
				affectsAny = true
				maxForThis := remaining[j] / buttonMatrix[buttonIdx][j]
				if maxForThis < maxPresses {
					maxPresses = maxForThis
				}
			}
		}

		if !affectsAny {
			// This button doesn't affect any remaining counter, skip it
			solve(orderIdx+1, remaining, totalPresses)
			return
		}

		// Try different numbers of presses for this button (from max down to 0 for better pruning)
		for presses := maxPresses; presses >= 0; presses-- {
			// Calculate new remaining values
			newRemaining := make([]int, numCounters)
			copy(newRemaining, remaining)
			for j := range numCounters {
				newRemaining[j] -= presses * buttonMatrix[buttonIdx][j]
			}

			solve(orderIdx+1, newRemaining, totalPresses+presses)
		}
	}

	solve(0, joltage, 0)

	if bestResult == math.MaxInt {
		fmt.Printf("No solution for target: %v\n", joltage)
		return 0
	}

	fmt.Printf("%v presses for target: %v\n", bestResult, joltage)
	return bestResult
}

func getButtonPresses(joltages [][]int, buttons [][][]int) int {
	// n := len(joltages)
	// if n == 0 {
	// 	return 0
	// }

	// results := make(chan int, n)
	// for i := range joltages {
	// 	go func() {
	// 		results <- processMachine(joltages[i], buttons[i])
	// 	}()
	// }

	// total := 0
	// for range n {
	// 	total += <-results
	// }
	// return total
	numPresses := 0

	for i := range joltages {
		numPresses += processMachine(joltages[i], buttons[i])
	}

	return numPresses
}

func main() {
	_, button, joltages, err := parseInput()
	if err != nil {
		panic(err)
	}

	// fmt.Println("target:", targets)
	// fmt.Println("button:", button)
	// fmt.Println("joltages:", joltages)

	numButtonPresses := getButtonPresses(joltages, button)
	fmt.Println("numButtonPresses:", numButtonPresses)

}
