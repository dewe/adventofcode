package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	slope, _ := readFile("topology.txt")
	fmt.Println("hits 1, 1: ", run(slope, 1, 1))
	fmt.Println("hits 3, 1: ", run(slope, 3, 1))
	fmt.Println("hits 5, 1: ", run(slope, 5, 1))
	fmt.Println("hits 7, 1: ", run(slope, 7, 1))
	fmt.Println("hits 1, 2: ", run(slope, 1, 2))
}

func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, scanner.Err()
}

func run(slope []string, right, down int) int {
	var x, hits int
	for y := 0; y < len(slope); y = y + down {
		row := slope[y]		
		hits = hits + moduloHit(row, x)
		x = x + right
	}
	return hits
}

func moduloHit(row string, x int) int {
	mod := len(row)
	if row[x%mod] == '#' {
		return 1
	}
	return 0
}
