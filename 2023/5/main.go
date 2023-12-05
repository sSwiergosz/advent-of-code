package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringArrayToIntArray(strArray []string) []int {
	var intArray []int

	for _, str := range strArray {
		intVal, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("Error converting string '%s' to int: %v\n", str, err)
		} else {
			intArray = append(intArray, intVal)
		}
	}

	return intArray
}

func readInput(input string) (string, []string) {
	file, err := os.Open(input)

	if err != nil {
		fmt.Println(err)
		panic(0)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var seeds string
	var blocks []string
	var block string

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "seeds:") {
			seeds = line
		} else {
			if line == "" {
				if len(block) > 0 {
					blocks = append(blocks, block)
					block = ""
				}
			} else {
				block += line + "\n"
			}
		}
	}

	if seeds != "" {
		blocks = append(blocks, block)
	}

	return seeds, blocks
}

func main() {
	seeds, blocks := readInput("input.txt")
	ranges := make([][3]int, 0)

	sl := strings.Split(strings.TrimSpace(strings.Split(seeds, ":")[1]), " ")
	numericSeeds := stringArrayToIntArray(sl)

	for _, b := range blocks {
		ranges = nil
		for i, line := range strings.Split(b, "\n") {
			if i == 0 {
				continue
			}

			if line == "" {
				continue
			}

			values := strings.Split(line, " ")
			a, _ := strconv.Atoi(values[0])
			b, _ := strconv.Atoi(values[1])
			c, _ := strconv.Atoi(values[2])

			ranges = append(ranges, [3]int{a, b, c})
		}

		newSeeds := []int{}
		for _, x := range numericSeeds {
			found := false
			for _, tuple := range ranges {
				a, b, c := tuple[0], tuple[1], tuple[2]
				if b <= x && x < b+c {
					newSeeds = append(newSeeds, x-b+a)
					found = true
					break
				}
			}

			if !found {
				newSeeds = append(newSeeds, x)
			}
		}
		numericSeeds = newSeeds
	}

	min := numericSeeds[0]
	for _, v := range numericSeeds {
		if v < min {
			min = v
		}
	}

	fmt.Println(min)
}
