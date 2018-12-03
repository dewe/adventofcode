package main

import (
	"bufio"
	"fmt"
	"os"
)

type fabric map[tile]int

func main() {

	claims := readClaims("input.txt")

	fabric := claimFabric(claims)

	fmt.Println("Overlapping:", countOverlaps(fabric))

	fmt.Println("Non overlapping:", findFirstNonOverlap(claims, fabric))

}

func readClaims(name string) []claim {
	result := []claim{}

	file, _ := os.Open(name)
	s := bufio.NewScanner(file)
	for s.Scan() {
		result = append(result, newClaim(s.Text()))
	}
	return result
}

func claimFabric(claims []claim) fabric {
	fabric := fabric{}

	for _, c := range claims {
		for _, t := range c.tiles() {
			fabric[t]++
		}
	}

	return fabric
}

func countOverlaps(f fabric) int {
	var count int
	for _, val := range f {
		if val > 1 {
			count++
		}
	}
	return count
}

func findFirstNonOverlap(claims []claim, f fabric) string {
	for _, c := range claims {
		if allFree(c.tiles(), f) {
			return c.id
		}
	}
	return ""
}

func allFree(tiles []tile, f fabric) bool {
	free := true
	for _, t := range tiles {
		if f[t] > 1 {
			free = false
		}
	}
	return free
}
