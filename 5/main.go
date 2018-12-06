package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"unicode"
)

func main() {

	// func ReadFile(filename string) ([]byte, error)
	buf, _ := ioutil.ReadFile("input.txt")
	result, passes := react(string(buf))
	fmt.Println("A1:", len(result), passes)

	chars := uniqueChars(string(buf))
	var min = len(result)
	for _, c := range chars {
		stripped := stripAndReact(string(buf), c)
		if len(stripped) >= min {
			continue
		}
		min = len(stripped)
	}

	fmt.Println("Shortest:", min)

}

func react(s string) (string, int) {
	var units = []byte(s)
	var reduced []byte
	var passes int

	for {
		passes++
		reduced = pass(units)

		if len(reduced) == len(units) {
			break
		}

		units = reduced
	}

	return string(reduced), passes
}

func pass(units []byte) []byte {

	for n := 0; n < len(units)-1; n++ {
		pair := units[n : n+2]
		lowered := bytes.ToLower(pair)

		if pair[0] != pair[1] && lowered[0] == lowered[1] {
			units = append(units[:n], units[n+2:]...)
		}
	}

	return units
}

func uniqueChars(s string) string {

	units := []byte(s)
	seen := make(map[byte]bool, len(units))
	result := []byte{}

	for _, b := range bytes.ToLower(units) {
		if seen[b] {
			continue
		}
		seen[b] = true
		result = append(result, b)
	}

	return string(result)
}

func strip(s string, b rune) string {
	var stripped []rune
	for _, c := range []rune(s) {
		if unicode.ToLower(c) != unicode.ToLower(b) {
			stripped = append(stripped, c)
		}
	}
	return string(stripped)
}

func stripAndReact(s string, b rune) string {
	result, _ := react(strip(s, b))
	return result
}
