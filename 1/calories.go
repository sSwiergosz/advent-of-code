package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func readLines(path string) ([]string, error) {
	calories, err := os.Open(path)
	scanner := bufio.NewScanner(calories)

	if err != nil {
		panic(err)
	}

	defer calories.Close()

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func createElvesCaloriesMap(calories []string) []int {
	var elves []int
	var count int

	for _, c := range calories {
		if c != "" {
			cal, _ := strconv.Atoi(c)
			count += cal
		} else {
			elves = append(elves, count)
			count = 0
		}
	}

	return elves
}

func sortElvesSlice(calories []string) []int {
	e := createElvesCaloriesMap(calories)

	sort.Slice(e, func(i, j int) bool {
		return e[i] > e[j]
	})

	return e
}

func main() {
	c, _ := readLines("calories.txt")
	e := sortElvesSlice(c)

	fmt.Printf("Max calories -> %v\n", e[0])
	fmt.Printf("TOP3 elves calories -> %v", e[0]+e[1]+e[2])
}
