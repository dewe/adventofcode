package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	coords := readCoords("input.txt")
	fmt.Println("len", len(coords))

	maxArea(coords)
}

type coordinate struct {
	x int
	y int
}

func readCoords(name string) []coordinate {
	result := []coordinate{}

	file, _ := os.Open(name)
	s := bufio.NewScanner(file)
	for s.Scan() {
		result = append(result, parse(s.Text()))
	}
	return result
}

func parse(s string) coordinate {
	parts := strings.Split(s, ", ")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return coordinate{x, y}
}

func maxArea(coords []coordinate) int {
	// find boundaries, min-max coordinates
	min, max := minAndMax(coords)

	fmt.Println("min", min)
	fmt.Println("max", max)

	// set up distance table
	locations := make([][]string, max.x)
	for i := range locations {
		locations[i] = make([]string, max.y)
	}

	locations = distanceTable(coords, locations)

	// find infinite areas around boundary edges
	// remove infinite coordinates
	// find max area among remaining coordinates

	return 0
}

func distanceTable(coords []coordinate, table [][]string) [][]string {

	// calculate closest manhattan distance for all locations
	xmin, xmax := 0, len(table)
	ymin, ymax := 0, len(table[0])

	for x := xmin; x < xmax; x++ {
		for y := ymin; y < ymax; y++ {
			table[x][y] = closest(coordinate{x, y}, coords)
		}
	}

	return table
}

// return the index of the closest coord, or "." if multiple close
func closest(loc coordinate, coords []coordinate) string {
	mindist := 10000000000
	closest := []int{}

	for index, coord := range coords {
		d := distance(loc, coord)

		if d == mindist {
			closest = append(closest, index)
		}

		if d < mindist {
			closest = []int{index}
			mindist = d
		}
	}

	if len(closest) != 1 {
		return "."
	}

	return fmt.Sprint(closest[0])

}

func minAndMax(coords []coordinate) (coordinate, coordinate) {
	var xmin, xmax = coords[0].x, coords[0].x
	var ymin, ymax = coords[0].y, coords[0].y

	for _, c := range coords {
		if c.x < xmin {
			xmin = c.x
		}
		if c.x > xmax {
			xmax = c.x
		}
		if c.y < ymin {
			ymin = c.y
		}
		if c.y > ymax {
			ymax = c.y
		}
	}

	return coordinate{xmin, ymin}, coordinate{xmax, ymax}
}

func distance(a, b coordinate) int {
	dx := abs(a.x - b.x)
	dy := abs(a.y - b.y)

	return dx + dy
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
