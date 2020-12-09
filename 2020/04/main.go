package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("passports.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	passports := scanBlankLines(file)
	fmt.Println("count: ", countPassports(passports))
}

func scanBlankLines(file io.Reader) []string {

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var result []string
	var head, tail int
	for ; head < len(lines); head++ {
		if len(strings.TrimSpace(lines[head])) == 0 {
			result = append(result, strings.Join(lines[tail:head], " "))
			tail = head + 1
		}
	}
	if head > tail {
		result = append(result, strings.Join(lines[tail:head], " "))
	}

	return result
}

func countPassports(passports []string) int {
	var count int
	for _, p := range passports {
		count = count + isValid(p)
	}
	return count
}

func isValid(passport string) int {
	if !strings.Contains(passport, "byr:") {
		return 0
	}
	// if !strings.Contains(passport, "cid:") { return 0 }
	if !strings.Contains(passport, "ecl:") {
		return 0
	}
	if !strings.Contains(passport, "eyr:") {
		return 0
	}
	if !strings.Contains(passport, "hcl:") {
		return 0
	}
	if !strings.Contains(passport, "hgt:") {
		return 0
	}
	if !strings.Contains(passport, "iyr:") {
		return 0
	}
	if !strings.Contains(passport, "pid:") {
		return 0
	}

	fields := strings.Fields(passport)
	// fmt.Printf("%#v\n", fields)
	for _, f := range fields {
		if !validField(f) {
			return 0
		}
	}

	return 1
}

func validField(field string) bool {
	kevyal := strings.Split(field, ":")
	key := kevyal[0]
	val := kevyal[1]

	switch key {
	case "byr":
		if len(val) != 4 {
			return false
		}
		if y, err := strconv.Atoi(val); err != nil || y < 1920 || y > 2002 {
			return false
		}
	case "iyr":
		if len(val) != 4 {
			return false
		}
		if y, err := strconv.Atoi(val); err != nil || y < 2010 || y > 2020 {
			return false
		}
	case "eyr":
		if len(val) != 4 {
			return false
		}
		if y, err := strconv.Atoi(val); err != nil || y < 2020 || y > 2030 {
			return false
		}
	case "hgt":
		if val[len(val)-2:] == "cm" {
			if x, err := strconv.Atoi(val[:len(val)-2]); err != nil || x < 150 || x > 193 {
				return false
			}
			break
		}
		if val[len(val)-2:] == "in" {
			if x, err := strconv.Atoi(val[:len(val)-2]); err != nil || x < 59 || x > 76 {
				return false
			}
			break
		}
		return false
	case "hcl":
		if m, err := regexp.MatchString("^#[0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f]$", val); err != nil || m == false {
			return false
		}
	case "ecl":
		var found bool
		for _, c := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
			if val == c {
				found = true
			}
		}
		if !found {
			return false
		}
	case "pid":
		if m, err := regexp.MatchString(`^\d\d\d\d\d\d\d\d\d$`, val); err != nil || m == false {
			return false
		}

	}

	return true
}
