package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pw struct {
	policy   policy
	password string
}

type policy struct {
	min    int
	max    int
	letter string
}

func main() {
	pws, _ := readPwFile()
	fmt.Println("valid: ", countValid(pws))

	ss, _ := ReadStr("passwords.txt")
	fmt.Println("valid toboggan: ", countValidToboggan(ss))
}

func readPwFile() ([]pw, error) {
	f, _ := os.Open("passwords.txt")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var result []pw
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		p := pw{
			parsePolicy(parts[0]),
			parts[1],
		}
		result = append(result, p)
	}

	return result, scanner.Err()
}

func parsePolicy(s string) policy {
	parts := strings.Fields(s)
	limits := strings.Split(parts[0], "-")
	min, _ := strconv.Atoi(limits[0])
	max, _ := strconv.Atoi(limits[1])
	return policy{min, max, parts[1]}
}

func countValid(pws []pw) int {
	var count int
	for _, pw := range pws {
		if pw.isValid() {
			count++
		}
	}
	return count
}

func (p *pw) isValid() bool {
	n := strings.Count(p.password, p.policy.letter)
	return p.policy.min <= n && n <= p.policy.max
}

func countValidToboggan(pws []string) int {
	var count int
	for _, s := range pws {
		if IsValid(s) {
			count++
		}
	}
	return count
}