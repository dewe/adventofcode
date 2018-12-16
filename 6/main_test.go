package main

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {

	coords := []coordinate{
		{1, 1},
		{1, 6},
		{8, 3},
		{3, 4},
		{5, 5},
		{8, 9},
	}

	if got, want := maxArea(coords), 17; got != want {
		t.Fatalf("maxArea = %v, want %v", got, want)
	}

}

func TestMinAndMax(t *testing.T) {

	tests := []struct {
		name  string
		input []coordinate
		min   coordinate
		max   coordinate
	}{
		{"identity", []coordinate{{1, 1}}, coordinate{1, 1}, coordinate{1, 1}},
		{"small-big", []coordinate{{1, 1}, {2, 3}}, coordinate{1, 1}, coordinate{2, 3}},
		{"crossing", []coordinate{{0, 4}, {5, 2}}, coordinate{0, 2}, coordinate{5, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			min, max := minAndMax(tt.input)

			if got, want := min, tt.min; got != want {
				t.Fatalf(" = %v, want %v", got, want)
			}
			if got, want := max, tt.max; got != want {
				t.Fatalf(" = %v, want %v", got, want)
			}
		})
	}
}

func TestDistance(t *testing.T) {

	tests := []struct {
		name     string
		a        coordinate
		b        coordinate
		distance int
	}{
		{"one", coordinate{1, 1}, coordinate{1, 2}, 1},
		{"twelve", coordinate{7, 1}, coordinate{1, 7}, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got, want := distance(tt.a, tt.b), tt.distance; got != want {
				t.Fatalf("distance() = %v, want %v", got, want)
			}
		})
	}
}

func TestDistanceTable(t *testing.T) {

	tests := []struct {
		name     string
		coords   []coordinate
		size     int
		expected string
	}{
		{
			name:     "one coord",
			coords:   []coordinate{{0, 0}},
			size:     2,
			expected: "[[0 0] [0 0]]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			locations := make([][]string, tt.size)
			for i := range locations {
				locations[i] = make([]string, tt.size)
			}

			distTable := distanceTable(tt.coords, locations)

			prettyTable(distTable)

			if got, want := fmt.Sprintf("%v", distTable), tt.expected; got != want {
				t.Fatalf("distTable = %v, want %v", got, want)
			}
		})
	}

}

func TestClosest(t *testing.T) {

	tests := []struct {
		name    string
		loc     coordinate
		coords  []coordinate
		closest string
	}{
		{
			name:    "just one coord",
			loc:     coordinate{0, 0},
			coords:  []coordinate{{1, 1}},
			closest: "0",
		},
		{
			name:    "two coords",
			loc:     coordinate{3, 3},
			coords:  []coordinate{{2, 2}, {3, 4}},
			closest: "1",
		},
		{
			name:    "three coords",
			loc:     coordinate{3, 3},
			coords:  []coordinate{{2, 2}, {3, 4}, {4, 4}},
			closest: "1",
		},
		{
			name:    "two equal of three coords",
			loc:     coordinate{3, 3},
			coords:  []coordinate{{2, 2}, {3, 4}, {2, 3}},
			closest: ".",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, want := closest(tt.loc, tt.coords), tt.closest; got != want {
				t.Fatalf("closest() = %q, want %q", got, want)
			}
		})
	}
}

func prettyTable(table [][]string) {

	fmt.Println("----------------")

	for y := range table[0] {
		for x := range table {
			fmt.Print(paddedString(table[x][y]))
		}
		fmt.Print("\n")
	}
	fmt.Println("---------------")
	fmt.Println()

}

func paddedString(s string) string {
	if s == "" {
		return "  "
	}
	return s + " "
}
