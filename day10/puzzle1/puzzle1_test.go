package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseInputFromReader_Sample(t *testing.T) {
	input := strings.TrimSpace(`
[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}
`)

	target, buttons, joltages, err := parseInputFromReader(strings.NewReader(input))
	if err != nil {
		t.Fatalf("parseInputFromReader returned error: %v", err)
	}

	wantTarget := [][]int{
		{0, 1, 1, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 1, 1, 0, 1},
	}
	if !reflect.DeepEqual(target, wantTarget) {
		t.Fatalf("target mismatch\n got: %#v\nwant: %#v", target, wantTarget)
	}

	wantButtons := [][][]int{
		{{3}, {1, 3}, {2}, {2, 3}, {0, 2}, {0, 1}},
		{{0, 2, 3, 4}, {2, 3}, {0, 4}, {0, 1, 2}, {1, 2, 3, 4}},
		{{0, 1, 2, 3, 4}, {0, 3, 4}, {0, 1, 2, 4, 5}, {1, 2}},
	}
	if !reflect.DeepEqual(buttons, wantButtons) {
		t.Fatalf("buttons mismatch\n got: %#v\nwant: %#v", buttons, wantButtons)
	}

	wantJoltages := [][]int{
		{3, 5, 4, 7},
		{7, 5, 12, 7, 2},
		{10, 11, 11, 5, 10, 5},
	}
	if !reflect.DeepEqual(joltages, wantJoltages) {
		t.Fatalf("joltages mismatch\n got: %#v\nwant: %#v", joltages, wantJoltages)
	}
}
