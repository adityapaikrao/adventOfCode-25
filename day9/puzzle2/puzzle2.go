package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	Queue "aoc/day9/Queue"
)

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
		pointCoords := strings.Split(text, ",")
		point := [2]int{}

		for i, value := range pointCoords {
			num, err := strconv.Atoi(value)
			if err != nil {
				err = fmt.Errorf("could not convert %v to int", value)
				panic(err)
			}

			point[i] = num
		}

		points = append(points, point)
	}

	return points
}

func calculateArea(point1 [2]int, point2 [2]int) int {
	area := 1
	for i := range point1 {
		area *= int(math.Abs(float64(point1[i]-point2[i]))) + 1
	}

	return area
}

/*Normalize points to optimize performance**/
func normalizePoints(points [][2]int, padding int) ([][2]int, int, int) {
	minX := points[0][0]
	minY := points[0][1]

	maxX := points[0][0]
	maxY := points[0][1]

	for _, point := range points {

		// update Xs
		if point[0] < minX {
			minX = point[0]
		}

		if point[0] > maxX {
			maxX = point[0]
		}

		// update Ys
		if point[1] < minY {
			minY = point[1]
		}

		if point[1] > maxY {
			maxY = point[1]
		}

	}

	for i := range points {
		points[i][0] += padding - minX
		points[i][1] += padding - minY
	}

	M := (maxY - minY) + (2 * padding) + 1
	N := maxX - minX + (2 * padding) + 1
	return points, N, M

}

/*Plots the points on a grid & returns the grid with 1s(reds) & 2s(greens) */
func markBoundaries(points [][2]int, N int, M int) [][]int {
	grid := make([][]int, N) // N * M grid
	for j := range N {
		grid[j] = make([]int, M)
	}

	numPoints := len(points)

	for k := range numPoints {
		curr := points[k]
		next := points[(k+1)%numPoints]

		if curr[0] == next[0] {
			for j := min(curr[1], next[1]); j <= max(curr[1], next[1]); j++ {
				grid[curr[0]][j] = 1
			}
		} else {
			for i := min(curr[0], next[0]); i <= max(curr[0], next[0]); i++ {
				grid[i][curr[1]] = 1
			}
		}
	}

	return grid

}

/*Check if the given coordinate is within bounds of the grid*/
func inbound(x int, y int, N int, M int) bool {
	if x < 0 || x >= N {
		return false
	} else if y < 0 || y >= M {
		return false
	} else {
		return true
	}
}

/*Uses the boundary of the points to fill up the green & red tiles with 1s (red) & 2s (green)*/
func fillGrid(grid [][]int) [][]int {
	q := Queue.NewQueue() // queue for BFS
	q.Push([2]int{0, 0})

	grid[0][0] = -1 // mark as visited

	N := len(grid)
	M := len(grid[0])

	directions := [][2]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}

	for !q.IsEmpty() {
		point := q.Popleft().([2]int)

		for _, offset := range directions {
			newX := point[0] + offset[0]
			newY := point[1] + offset[1]

			if inbound(newX, newY, N, M) && grid[newX][newY] == 0 {
				q.Push([2]int{newX, newY})
				grid[newX][newY] = -1
			}
		}
	}

	// set negative elems (out of boundary to zero)
	for i := range N {
		for j := range M {
			if grid[i][j] == -1 {
				grid[i][j] = 0
			} else {
				grid[i][j] = 1
			}
		}
	}
	return grid

}

/*Compute prefix sum of the grid*/
func computePrefixSum(grid [][]int) [][]int {
	N := len(grid)
	M := len(grid[0])

	prefix := make([][]int, N)
	for i := range N {
		prefix[i] = make([]int, M)
	}

	for i := 1; i < N; i++ {
		for j := 1; j < M; j++ {
			prefix[i][j] = prefix[i-1][j] + prefix[i][j-1] - prefix[i-1][j-1] + grid[i][j]
		}
	}

	return prefix
}

/*Checks if rectangle formed by point1, point2 is contained within the green & red area*/
func isValid(point1 [2]int, point2 [2]int, prefix [][]int) bool {
	expectedArea := calculateArea(point1, point2)

	minX := min(point1[0], point2[0])
	minY := min(point1[1], point2[1])
	maxX := max(point1[0], point2[0])
	maxY := max(point1[1], point2[1])

	// To get sum from (minX, minY) to (maxX, maxY) inclusive, we need:
	// prefix[maxX][maxY] - prefix[minX-1][maxY] - prefix[maxX][minY-1] + prefix[minX-1][minY-1]
	actualArea := prefix[maxX][maxY] - prefix[maxX][minY-1] - prefix[minX-1][maxY] + prefix[minX-1][minY-1]

	return actualArea == expectedArea

}

// Shoelace formula to calculate polygon area (returns 2x the area to avoid float)
func shoelaceArea(points [][2]int) int {
	n := len(points)
	area := 0
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		area += points[i][0] * points[j][1]
		area -= points[j][0] * points[i][1]
	}
	if area < 0 {
		area = -area
	}
	return area / 2
}

// Calculate perimeter (boundary points) - Manhattan distance for grid movements
func calcPerimeter(points [][2]int) int {
	n := len(points)
	perim := 0
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		dx := points[j][0] - points[i][0]
		dy := points[j][1] - points[i][1]
		if dx < 0 {
			dx = -dx
		}
		if dy < 0 {
			dy = -dy
		}
		perim += dx + dy
	}
	return perim
}

// Pick's theorem: A = i + b/2 - 1
// So: i = A - b/2 + 1
// Total points (interior + boundary) = i + b = A + b/2 + 1
func totalPolygonArea(points [][2]int) int {
	A := shoelaceArea(points)
	b := calcPerimeter(points)
	return A + b/2 + 1
}

func getMaxArea(points [][2]int) int {
	fmt.Println("intial points: \n", points)

	// Use Shoelace + Pick's theorem for O(n) solution
	area := totalPolygonArea(points)
	fmt.Println("Total area (using Shoelace + Pick's theorem):", area)

	return area
}

func main() {
	points := parseInput()

	fmt.Println(getMaxArea(points))
}
