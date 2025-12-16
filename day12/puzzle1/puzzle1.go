package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(shapeCount int) (map[int][][]rune, map[int][]int, map[int][]int) {
	shapes := make(map[int][][]rune)
	regions := make(map[int][]int)
	regionShapes := make(map[int][]int)

	filePath := "../puzzle.in"
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for range shapeCount {
		isKey := true
		idx := -1
		for scanner.Scan() {
			text := strings.TrimSpace(scanner.Text())
			if text == "" {
				break
			}
			if isKey {
				num, err := strconv.Atoi(string(text[0]))
				if err != nil {
					err = fmt.Errorf("could not convert %v to int", text[0])
					panic(err)
				}
				shapes[num] = make([][]rune, 0)
				isKey = false
				idx = num
			} else {
				currRow := make([]rune, 0)
				for _, char := range text {
					currRow = append(currRow, char)
				}
				shapes[idx] = append(shapes[idx], currRow)
			}
		}
	}

	regionIdx := 0
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		textList := strings.Split(text, ":")

		regions[regionIdx] = make([]int, 0)
		regionShapes[regionIdx] = make([]int, 0)

		// read in dimensions
		for char := range strings.SplitSeq(textList[0], "x") {
			if num, err := strconv.Atoi(char); err == nil {
				regions[regionIdx] = append(regions[regionIdx], num)
			}
		}

		// read in num shapes
		values := strings.FieldsSeq(textList[1])
		for char := range values {
			num, err := strconv.Atoi(char)
			if err != nil {
				err = fmt.Errorf("could not convert %v to int", char)
				panic(err)
			}
			regionShapes[regionIdx] = append(regionShapes[regionIdx], num)
		}

		regionIdx++

	}

	return shapes, regions, regionShapes
}

func canFit(regionShape []int, regionArea int, shapes map[int][][]rune, shapeAreas map[int]int) bool {
	allShapeAreas := 0

	for shapeIdx, count := range regionShape {
		allShapeAreas += count * shapeAreas[shapeIdx]
	}

	return allShapeAreas <= regionArea
}

func getRegions(shapes map[int][][]rune, regions map[int][]int, regionShapes map[int][]int) (int, []int) {
	numRegions := 0
	possibleInvalidRegions := make([]int, 0)

	shapeAreas := make(map[int]int, 0)
	for idx, shape := range shapes {
		area := 0
		for _, row := range shape {
			for _, char := range row {
				if char == '#' {
					area++
				}
			}
		}
		shapeAreas[idx] = area
	}

	regionAreas := make(map[int]int, 0)
	for idx, region := range regions {
		area := 1
		for _, dim := range region {
			area *= dim
		}
		regionAreas[idx] = area
	}

	for i := range regions {
		if canFit(regionShapes[i], regionAreas[i], shapes, shapeAreas) {
			numRegions += 1
			possibleInvalidRegions = append(possibleInvalidRegions, i)
		}
	}

	return numRegions, possibleInvalidRegions
}

func main() {
	shapes, regions, regionShapes := parseInput(6)

	numValidregions, _ := getRegions(shapes, regions, regionShapes)
	fmt.Println("Number of Valid Regions:", numValidregions)
	// fmt.Println("Possible invalid Regions:", possibleInvalidRegions)
}
