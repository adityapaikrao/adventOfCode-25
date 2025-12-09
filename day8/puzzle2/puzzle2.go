package main

import (
	"aoc/day8/minheap"
	uf "aoc/day8/unionfind"
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculateDistance(point1 [3]int, point2 [3]int) int {
	dist := 0
	for i := range point1 {
		diff := point1[i] - point2[i]
		dist += diff * diff
	}
	return dist
}

func parseInput() [][3]int {
	path := "../puzzle1.in"
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	points := make([][3]int, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		textList := strings.Split(text, ",")

		var curr_point [3]int
		for i, coord := range textList {
			if num, err := strconv.Atoi(coord); err == nil {
				curr_point[i] = num
			} else {
				panic(fmt.Errorf("could not convert %s to int", coord))
			}
		}
		points = append(points, curr_point)
	}
	return points

}

func makeConnections(pointSet *uf.UnionFind, minHeap *minheap.MinHeap, maxConnections int) [2]int {
	lastPair := [2]int{}

	for pointSet.NumComponents != 1 && minHeap.Len() > 0 {
		edge := heap.Pop(minHeap).([3]int)
		point1, point2 := edge[0], edge[1]

		pointSet.Union(point1, point2)

		lastPair = [2]int{point1, point2}
	}

	return lastPair
}

func main() {
	points := parseInput()
	maxConnections := 1000

	minHeap := &minheap.MinHeap{}
	heap.Init(minHeap)

	for i := range points {
		for j := i + 1; j < len(points); j++ {
			dist := calculateDistance(points[i], points[j])
			edge := [3]int{i, j, dist}
			heap.Push(minHeap, edge)
		}
	}

	pointSet := uf.NewUnionFind(len(points))
	lastPair := makeConnections(pointSet, minHeap, maxConnections)

	fmt.Println("The last points are:", points[lastPair[0]], points[lastPair[1]])
	fmt.Println("Product:", points[lastPair[0]][0]*points[lastPair[1]][0])
}
