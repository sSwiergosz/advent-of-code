package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	scanner := bufio.NewScanner(file)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func main() {
	lines, _ := readLines("input.txt")
	var count = 0

	for _, l := range lines {
		s := strings.Split(l, ",")

		elf1, elf2 := s[0], s[1]

		assignmentRange1 := strings.Split(elf1, "-")
		assignmentRange2 := strings.Split(elf2, "-")

		firstStart, _ := strconv.Atoi(assignmentRange1[0])
		firstEnd, _ := strconv.Atoi(assignmentRange1[1])
		secondStart, _ := strconv.Atoi(assignmentRange2[0])
		secondEnd, _ := strconv.Atoi(assignmentRange2[1])

		if secondStart >= firstStart && secondStart <= firstEnd || firstStart >= secondStart && secondEnd >= firstStart {
			count++
		}
	}

	fmt.Print(count)
}
