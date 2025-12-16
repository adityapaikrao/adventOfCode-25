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

func dfs(node string, adj map[string][]string, memo *map[string]int) int {
	if val, ok := (*memo)[node]; ok {
		return val
	}
	if node == "out" {
		return 1
	}

	paths := 0

	for _, dst := range adj[node] {
		paths += dfs(dst, adj, memo)
	}

	(*memo)[node] = paths
	return paths
}

func getNumberOfPaths(adj map[string][]string) int {
	memo := make(map[string]int)

	return dfs("you", adj, &memo)

}

func main() {
	adj := parseInput()
	fmt.Println("Number of Paths:", getNumberOfPaths(adj))
}
