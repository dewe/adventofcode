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

func makeDistanceTable(coords []coordinate) [][]string {
	_, max := minAndMax(coords)
	sizex, sizey := max.x+1, max.y+1

	distTable := make([][]string, sizey)
	for y := range distTable {
		distTable[y] = make([]string, sizex)
	}

	for y := range distTable {
		for x := range distTable[y] {
			// c := closest(coordinate{x, y}, coords)
			// base := int('a')

			// println(c, base, string(base+c))

			distTable[y][x] = closest2(coordinate{x, y}, coords)
		}
	}

	return distTable
}

func removeEdgeVales(distTable [][]string) [][]string {
	removeMap := map[string]int{}

	for _, s := range distTable[0] {
		if s != "." {
			removeMap[s]++
		}
	}

	for _, s := range distTable[len(distTable)-1] {
		if s != "." {
			removeMap[s]++
		}
	}

	for y := range distTable {
		for x := range distTable[y] {
			s := strings.ToLower(distTable[y][x])
			if removeMap[s] > 0 {
				distTable[y][x] = "."
			}
		}
	}

	return distTable
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

func closest2(loc coordinate, coords []coordinate) string {
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

	index := closest[0]

	base := int('a')
	s := string(base + index)

	if coords[index] == loc {
		s = strings.ToUpper(s)
	}

	return s
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
