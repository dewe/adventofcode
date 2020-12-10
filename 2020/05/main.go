package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	codes,_ := readStrings("codes.txt")

	var ids []int
	for _, code := range codes {
		ids = append(ids, seatId(code))
	}
	sort.Ints(ids)
	fmt.Println(ids[len(ids)-1])
	fmt.Println(findMissing(ids))
}

func readStrings(path string) ([]string, error) {
	file, _ := os.Open(path)
	defer file.Close()

	var result []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, scanner.Err()
}

func seatId(code string) int {
	row := binaryPartition(code[0:7], 'F')
	col := binaryPartition(code[7:10], 'L')
	return row*8 + col
}

func binaryPartition(code string, head rune) int {
	var start int
	var end int = int(math.Pow(2, float64(len(code)))) - 1

	for _, x := range code {
		if x == head {
			end = start + (end-start)/2
		} else {
			start = start + (end-start)/2 + 1
		}
	}

	return start
}

func findMissing(ints []int) int {
	for i := 0; i < len(ints); i++ {
		if ints[i+1] - ints[i] == 2 {
			return ints[i] + 1
		}
	}
	return 0
}