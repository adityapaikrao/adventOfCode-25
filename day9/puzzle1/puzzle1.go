package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func calculateArea(point1 [2]int, point2 [2]int) int {
	area := 1
	for i := range point1 {
		area *= int(math.Abs(float64(point1[i]-point2[i]) + 1))
	}
	return area
}

func parseInput() [][2]int {
	filePath := "../puzzle1.in"
	points := make([][2]int, 0)

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		textList := strings.Split(text, ",")
		coords := [2]int{}

		for i, point := range textList {
			point, err := strconv.Atoi(point)
			if err != nil {
				err = fmt.Errorf("could not convert %v to int", point)
				panic(err)
			}
			coords[i] = point
		}

		points = append(points, coords)
	}

	return points
}

func getMaxArea(points [][2]int) int {
	maxArea := 0
	for i, point1 := range points {
		for j := i + 1; j < len(points); j++ {
			point2 := points[j]
			area := calculateArea(point1, point2)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func main() {
	points := parseInput()

	// fmt.Println(points)
	fmt.Println(getMaxArea(points))
}
