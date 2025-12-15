package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	Queue "aoc/day9/Queue"
)

func parseInput() [][2]int {
	filePath := "../puzzle1.in"
	// filePath := "test.in" // Uncomment for testing
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

// Compress coordinates to reduce grid size
// Returns compressed points, x mapping, y mapping
func compressCoordinates(points [][2]int) ([][2]int, []int, []int) {
	xSet := make(map[int]bool)
	ySet := make(map[int]bool)

	for _, p := range points {
		xSet[p[0]] = true
		ySet[p[1]] = true
	}

	// Convert to sorted slices
	xVals := make([]int, 0, len(xSet))
	yVals := make([]int, 0, len(ySet))
	for x := range xSet {
		xVals = append(xVals, x)
	}
	for y := range ySet {
		yVals = append(yVals, y)
	}
	sort.Ints(xVals)
	sort.Ints(yVals)

	// Create mapping from original to compressed
	xToIdx := make(map[int]int)
	yToIdx := make(map[int]int)
	for i, x := range xVals {
		xToIdx[x] = i + 1 // +1 for padding
	}
	for i, y := range yVals {
		yToIdx[y] = i + 1 // +1 for padding
	}

	// Compress points
	compressed := make([][2]int, len(points))
	for i, p := range points {
		compressed[i] = [2]int{xToIdx[p[0]], yToIdx[p[1]]}
	}

	return compressed, xVals, yVals
}

/*Check if the given coordinate is within bounds of the grid*/
func inbound(x int, y int, N int, M int) bool {
	return x >= 0 && x < N && y >= 0 && y < M
}

/*Mark the boundary edges between consecutive red points*/
func markBoundaries(points [][2]int, N int, M int) [][]int {
	grid := make([][]int, N)
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

/*Uses BFS from outside to mark exterior, then interior becomes the valid area*/
func fillGrid(grid [][]int) [][]int {
	q := Queue.NewQueue()
	q.Push([2]int{0, 0})

	grid[0][0] = -1 // mark as visited (exterior)

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

	// -1 (exterior) -> 0, everything else (boundary + interior) -> 1
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

/*Compute prefix sum of the grid for O(1) rectangle sum queries*/
func computePrefixSum(grid [][]int) [][]int {
	N := len(grid)
	M := len(grid[0])

	// Compute prefix sum
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

/*Check if rectangle from (x1,y1) to (x2,y2) is completely inside the polygon*/
func isValidRect(x1, y1, x2, y2 int, prefix [][]int) bool {
	minX := min(x1, x2)
	minY := min(y1, y2)
	maxX := max(x1, x2)
	maxY := max(y1, y2)

	// Expected number of cells in the rectangle
	expectedCells := (maxX - minX + 1) * (maxY - minY + 1)

	// Actual cells that are inside the polygon
	actualCells := prefix[maxX][maxY] - prefix[minX-1][maxY] - prefix[maxX][minY-1] + prefix[minX-1][minY-1]

	return actualCells == expectedCells
}

func getMaxArea(points [][2]int) int {
	maxArea := 0

	// Compress coordinates
	compressed, xVals, yVals := compressCoordinates(points)

	N := len(xVals) + 2 // +2 for padding on both sides
	M := len(yVals) + 2

	fmt.Println("Compressed grid size:", N, "x", M)

	// Mark boundaries and fill
	grid := markBoundaries(compressed, N, M)
	grid = fillGrid(grid)

	// Compute prefix sum
	prefix := computePrefixSum(grid)

	// Check all pairs of red points
	for i := 0; i < len(compressed); i++ {
		for j := i + 1; j < len(compressed); j++ {
			p1 := compressed[i]
			p2 := compressed[j]

			if isValidRect(p1[0], p1[1], p2[0], p2[1], prefix) {
				// Calculate actual area using original coordinates
				origP1 := points[i]
				origP2 := points[j]
				width := abs(origP1[0]-origP2[0]) + 1
				height := abs(origP1[1]-origP2[1]) + 1
				area := width * height

				if area > maxArea {
					maxArea = area
					fmt.Printf("New max: points %v and %v, area %d\n", origP1, origP2, area)
				}
			}
		}
	}

	return maxArea
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	points := parseInput()
	fmt.Printf("Loaded %d points\n", len(points))
	fmt.Println("Max area:", getMaxArea(points))
}
