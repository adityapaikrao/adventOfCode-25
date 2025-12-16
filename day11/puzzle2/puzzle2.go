package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseInput() map[string][]string {
	filePath := "../puzzle.in"

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	adj := make(map[string][]string, 0)

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		textList := strings.Fields(text)

		src := textList[0][:len(textList[0])-1]
		adj[src] = make([]string, 0)

		for i := 1; i < len(textList); i++ {
			adj[src] = append(adj[src], textList[i])
		}
	}

	return adj
}

func makeKey(node string, visBoth [2]bool) string {
	return fmt.Sprintf("%s|%t|%t", node, visBoth[0], visBoth[1])
}

func dfs(node string, adj map[string][]string, memo *map[string]int, visBoth [2]bool) int {
	mapKey := makeKey(node, visBoth)
	if val, ok := (*memo)[mapKey]; ok {
		return val
	}
	if node == "out" {
		if visBoth[0] && visBoth[1] {
			return 1
		} else {
			return 0
		}
	}

	paths := 0

	for _, dst := range adj[node] {
		newVis := visBoth
		switch dst {
		case "dac":
			newVis[0] = true
		case "fft":
			newVis[1] = true
		}
		paths += dfs(dst, adj, memo, newVis)
	}

	(*memo)[mapKey] = paths
	return paths
}

func getNumberOfPaths(adj map[string][]string) int {
	memo := make(map[string]int)
	visBoth := [2]bool{false, false}

	return dfs("svr", adj, &memo, visBoth)

}

func main() {
	adj := parseInput()
	fmt.Println("Number of Paths:", getNumberOfPaths(adj))
}
