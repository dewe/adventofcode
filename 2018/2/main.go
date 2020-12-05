package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ids := readLines("input.txt")

	c := checksum(ids)
	fmt.Println("Checksum:", c)

	boxes := findBoxes(ids)
	fmt.Println("Common:", common(boxes[0], boxes[1]))
}

func readLines(name string) []string {
	result := []string{}

	file, _ := os.Open(name)
	s := bufio.NewScanner(file)
	for s.Scan() {
		result = append(result, s.Text())
	}
	return result
}

func xplicates(id string) (int, int) {

	counts := countRunes(id)

	var dupl, tripl int
	for _, val := range counts {

		if val == 2 {
			dupl = 1
		}

		if val == 3 {
			tripl = 1
		}
	}

	return dupl, tripl
}

func countRunes(s string) map[rune]int {
	m := map[rune]int{}

	for _, rune := range s {
		m[rune]++
	}

	return m
}

func checksum(ids []string) int {
	var dupl, tripl int

	for _, id := range ids {
		d, t := xplicates(id)
		dupl += d
		tripl += t
	}

	return dupl * tripl
}

func findBoxes(ids []string) []string {
	for _, s1 := range ids {
		for _, s2 := range ids {
			if stringDiff(s1, s2) == 1 {
				return []string{s1, s2}
			}
		}
	}

	return []string{}
}

func stringDiff(s1 string, s2 string) int {
	var diff int
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff++
		}
	}
	return diff
}

func common(s1, s2 string) string {
	common := []byte{}

	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			common = append(common, s1[i])
		}
	}

	return string(common)
}
