package main

import (
	"aoc/day8/minheap"
	uf "aoc/day8/unionfind"
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"slices"
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

func makeConnections(pointSet *uf.UnionFind, minHeap *minheap.MinHeap, maxConnections int) {
	numConnections := 0

	for numConnections < maxConnections && minHeap.Len() > 0 {
		edge := heap.Pop(minHeap).([3]int)
		point1, point2 := edge[0], edge[1]

		pointSet.Union(point1, point2)
		numConnections++
	}
}

func main() {
	points := parseInput()
	nlargest := 3
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
	makeConnections(pointSet, minHeap, maxConnections)

	sizes := make([]int, 0)
	for id := range points {
		if pointSet.Find(id) == id {
			sizes = append(sizes, pointSet.ComponentSize(id))
		}
	}

	slices.Sort(sizes)
	slices.Reverse(sizes)

	fmt.Println("Component sizes:", sizes, "Num Components:", pointSet.NumComponents)

	answer := 1
	for i := range min(nlargest, len(sizes)) {
		answer *= sizes[i]
	}

	fmt.Println(answer)

}
