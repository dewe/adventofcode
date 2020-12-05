package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	input := mustReadInts("input.txt")

	fmt.Println("A1:", sumAll(input))

	fmt.Println("A2:", findFirstRepeat(input))
}

func mustReadInts(name string) []int {
	// read file into []int
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	ints, err := ReadInts(file)
	if err != nil {
		panic(err)
	}

	return ints
}

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) ([]int, error) {
	var result []int

	scanner := bufio.NewScanner(r)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {

		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}

		result = append(result, x)
	}

	return result, scanner.Err()
}

func sumAll(ints []int) int {
	var sum int
	for _, v := range ints {
		sum += v
	}
	return sum
}

func findFirstRepeat(ints []int) int {
	m := make(map[int]bool)
	m[0] = true
	i, sum := 0, 0
	for {
		sum += ints[i]

		if m[sum] {
			return sum
		}

		m[sum] = true
		i = (i + 1) % len(ints)
	}
}
