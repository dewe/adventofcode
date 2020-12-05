package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("expensereport.txt")
	expenses, err := ReadInts(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(expenses)
	fmt.Println(FindTwo2020(expenses))
	fmt.Println(FindThree2020(expenses))
}

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

// FindTwo2020 finds the two ints that sums to 2020 and return their product
func FindTwo2020(expenses []int) int {

	for i := 0; i < len(expenses); i++ {
		for j := i + 1; j < len(expenses); j++ {
			m := expenses[i]
			n := expenses[j]
			if m+n == 2020 {
				return m * n
			}
		}
	}
	return 0
}

// FindThree2020 finds the two ints that sums to 2020 and return their product
func FindThree2020(expenses []int) int {

	for i := 0; i < len(expenses); i++ {
		for j := i + 1; j < len(expenses); j++ {
			for k := i + 2; k < len(expenses); k++ {
				m := expenses[i]
				n := expenses[j]
				o := expenses[k]
				if m+n+o == 2020 {
					return m * n *o
				}
			}
		}
	}
	return 0
}
