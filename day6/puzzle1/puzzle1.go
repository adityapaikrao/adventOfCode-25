package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*Converts a string slice to int slice */
func sliceAtoi(strSlice []string) ([]int, error) {

	intSlice := make([]int, len(strSlice))

	for i, char := range strSlice {
		num, err := strconv.Atoi(char)
		if err != nil {
			return intSlice, err
		}
		intSlice[i] = num
	}
	return intSlice, nil
}

func parseInput() ([]string, [][]int, error) {
	path := "../puzzle1.in"
	file, err := os.Open(path)

	operations := []string{}
	values := make([][]int, 0)

	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		textList := strings.Fields(text)

		if textList[0] == "*" || textList[0] == "+" {

			operations = append(operations, textList...) // ... unpacts the text_list into string chars
		} else {
			intValues, err := sliceAtoi(textList)
			if err != nil {
				return operations, values, err
			}
			values = append(values, intValues)
		}
	}

	return operations, values, nil

}

func solveProblem(operations []string, values [][]int) int {
	totalSum := 0

	for i, op := range operations {
		curr := 0
		if op == "*" {
			curr = 1
		}

		for _, row := range values {
			// fmt.Printf("curr idx: %v and values[idx] %v", i, row)
			if op == "*" {
				curr *= row[i]
			} else {
				curr += row[i]
			}
		}

		totalSum += curr
	}

	return totalSum
}
func main() {
	operations, values, err := parseInput()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	// fmt.Println("Operations:", operations)
	// fmt.Println("values:", values)

	fmt.Print(solveProblem(operations, values))
}
