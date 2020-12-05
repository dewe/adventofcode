package main

import (
	"fmt"
	"testing"

	"github.com/andreyvit/diff"
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
		{
			name:     "two coords",
			coords:   []coordinate{{0, 0}, {2, 2}},
			size:     3,
			expected: "[[0 0 .] [0 . 1] [. 1 1]]",
		},
		{
			name:     "two coords",
			coords:   []coordinate{{1, 1}, {1, 6}, {8, 3}, {3, 4}, {5, 5}, {8, 9}},
			size:     10,
			expected: "[[0 0 0 0 0 . 2 2 2 2] [0 0 0 0 0 . 2 2 2 2] [0 0 0 3 3 4 2 2 2 2] [0 0 3 3 3 4 2 2 2 2] [. . 3 3 3 4 2 2 2 2] [1 1 . 3 4 4 4 4 2 2] [1 1 1 . 4 4 4 4 . .] [1 1 1 . 4 4 4 5 5 5] [1 1 1 . 4 4 5 5 5 5] [1 1 1 . 5 5 5 5 5 5]]",
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

func TestMakeDistanceTable(t *testing.T) {

	tests := []struct {
		name     string
		coords   []coordinate
		expected string
	}{
		{
			name:     "one coord",
			coords:   []coordinate{{1, 1}},
			expected: "[[a a] [a A]]",
		},
		{
			name:     "example coords",
			coords:   []coordinate{{1, 1}, {1, 6}, {8, 3}, {3, 4}, {5, 5}, {8, 9}},
			expected: "[[a a a a a . c c c] [a A a a a . c c c] [a a a d d e c c c] [a a d d d e c c C] [. . d D d e e c c] [b b . d e E e e c] [b B b . e e e e .] [b b b . e e e f f] [b b b . e e f f f] [b b b . f f f f F]]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			distTable := makeDistanceTable(tt.coords)

			prettyTable(distTable)

			if got, want := fmt.Sprintf("%v", distTable), tt.expected; got != want {
				t.Fatalf("distTable = %v, want %v", got, want)
			}
		})
	}

}

func TestRemoveEdgeValues(t *testing.T) {
	distTable := makeDistanceTable([]coordinate{{1, 1}, {1, 6}, {8, 3}, {3, 4}, {5, 5}, {8, 9}})
	expected := "[[. . . . . . . . .] [. . . . . . . . .] [. . . d d e . . .] [. . d d d e . . .] [. . d D d e e . .] [. . . d e E e e .] [. . . . e e e e .] [. . . . e e e . .] [. . . . e e . . .] [. . . . . . . . .]]"

	removed := removeEdgeVales(distTable)

	if got, want := fmt.Sprint(removed), expected; got != want {
		t.Fatalf("removeEdgeVales(), result not as expected:\n%v", diff.LineDiff(want, got))
	}
}

func TestCountAreas(t *testing.T) {
	// count leftovers in map, return map
}

func TestClosest2(t *testing.T) {

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
			closest: "a",
		},
		{
			name:    "same coord",
			loc:     coordinate{1, 1},
			coords:  []coordinate{{0, 1}, {1, 1}},
			closest: "B",
		},
		{
			name:    "two coords",
			loc:     coordinate{3, 3},
			coords:  []coordinate{{2, 2}, {3, 4}},
			closest: "b",
		},
		{
			name:    "three coords",
			loc:     coordinate{3, 3},
			coords:  []coordinate{{2, 2}, {3, 4}, {4, 4}},
			closest: "b",
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
			if got, want := closest2(tt.loc, tt.coords), tt.closest; got != want {
				t.Fatalf("closest() = %q, want %q", got, want)
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

func prettyTable(table [][]string, coords []coordinate) {

	fmt.Println("----------------")

	for x := range table {
		for y := range table[0] {
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
