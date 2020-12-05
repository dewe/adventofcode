package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadStrings does stuff
func ReadStr(path string) ([]string, error) {
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

// IsValid does other stuff
func IsValid(password string) bool {
	parts := strings.Fields(password)
	pos1, _ := strconv.Atoi(strings.Split(parts[0], "-")[0])
	pos2, _ := strconv.Atoi(strings.Split(parts[0], "-")[1])
	b := strings.TrimRight(parts[1], ":")[0]
	pw := parts[2]

	return (pw[pos1-1] == b) != (pw[pos2-1] == b)
}
