package main

import (
	Queue "aoc/day9/Queue"
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
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

func reachedTarget(state []int, target []int) bool {
	for i := range state {
		if state[i] != target[i] {
			return false
		}
	}
	return true
}

func processMachine(target []int, buttonsRow [][]int) int {
	levels := 0
	q := Queue.NewQueue() // stores the (state)
	q.Push(make([]int, len(target)))

	for !q.IsEmpty() && levels < 10 {
		size := q.Size()
		// fmt.Println("At level", levels)
		// q.PrintQueue()

		for size > 0 {
			state := q.Popleft().([]int)
			// fmt.Println("checking for state", state)
			if reachedTarget(state, target) {
				// fmt.Printf("%v presses for target: %v \n", levels, target)
				return levels
			}

			for _, button := range buttonsRow {
				newState := slices.Clone(state)
				for _, index := range button {
					newState[index] ^= 1
				}
				q.Push(newState)
			}
			size -= 1

		}

		levels += 1

	}
	fmt.Printf("reached limit :( %v", levels)
	return levels

}

func getButtonPresses(targets [][]int, buttons [][][]int) int {
	numPresses := 0

	for i := range targets {
		numPresses += processMachine(targets[i], buttons[i])
	}

	return numPresses
}

func main() {
	targets, button, _, err := parseInput()
	if err != nil {
		panic(err)
	}

	// fmt.Println("target:", targets)
	// fmt.Println("button:", button)
	// fmt.Println("joltages:", joltages)

	numButtonPresses := getButtonPresses(targets, button)
	fmt.Println("numButtonPresses:", numButtonPresses)

}
