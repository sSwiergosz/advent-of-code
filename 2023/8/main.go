package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	// Process the input
	instructions := strings.Split(input[0], "")
	coordMap := make(map[string][2]string)
	for i := 2; i < len(input); i++ {
		parts := strings.Split(input[i], " = ")
		coords := strings.Split(strings.Trim(parts[1], "()"), ", ")
		coordMap[parts[0]] = [2]string{coords[0], coords[1]}
	}

	// Part 1 logic
	steps := 0
	arrived := false
	currentCoord := "AAA"
	instructionIndex := 0

	for !arrived {
		nextStep := instructions[instructionIndex]
		currentCoordObj := coordMap[currentCoord]
		var nextCoord string
		if nextStep == "L" {
			nextCoord = currentCoordObj[0]
		} else {
			nextCoord = currentCoordObj[1]
		}

		if nextCoord == "ZZZ" {
			arrived = true
		} else {
			currentCoord = nextCoord
			instructionIndex++
			if instructionIndex >= len(instructions) {
				instructionIndex = 0
			}
		}
		steps++
	}

	// Print the result
	fmt.Println("Number of steps:", steps)
}
