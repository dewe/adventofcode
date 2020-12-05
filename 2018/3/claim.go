package main

import (
	"fmt"
)

type claim struct {
	id     string
	left   int
	top    int
	width  int
	height int
}

type tile struct {
	x int
	y int
}

// parse "#1 @ 1,3: 4x4" into claim
func newClaim(s string) claim {
	var id string
	var left, top, width, height int

	fmt.Sscanf(s, "#%s @ %d,%d: %dx%d", &id, &left, &top, &width, &height)

	return claim{id, left, top, width, height}
}

func (c claim) tiles() []tile {
	tiles := []tile{}
	for x := c.left; x < c.left+c.width; x++ {
		for y := c.top; y < c.top+c.height; y++ {
			tiles = append(tiles, tile{x, y})
		}
	}
	return tiles
}
