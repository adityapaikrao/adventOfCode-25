package minheap

import "fmt"

type MinHeap [][3]int // stores (point ID1, point ID2, distance)

func (minheap MinHeap) Len() int {
	return len(minheap)
}

func (minHeap MinHeap) Less(i, j int) bool {
	return minHeap[i][2] < minHeap[j][2] // compare on dist
}

func (minHeap MinHeap) Swap(i, j int) {
	minHeap[i], minHeap[j] = minHeap[j], minHeap[i]
}

func (minHeap *MinHeap) Push(x any) {
	val, ok := x.([3]int)
	if !ok {
		err := fmt.Errorf("Invalid type for x: %v")
		panic((err))
	}
	*minHeap = append(*minHeap, val) // what does
}

func (minHeap *MinHeap) Pop() any {
	oldHeap := *minHeap
	n := len(oldHeap)

	popEle := oldHeap[n-1]
	*minHeap = oldHeap[:n-1]

	return popEle
}
