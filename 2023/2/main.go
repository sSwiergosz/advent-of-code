package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type limits struct {
	red   int
	blue  int
	green int
}

func isWithinLimit(input string, limit int) bool {
	splittedResult := strings.Split(input, " ")
	number, _ := strconv.Atoi(splittedResult[0])
	return number <= limit
}

func main() {
	guessesLimit := limits{12, 14, 13}
	readFile, err := os.Open("input.txt")

	re := regexp.MustCompile(`Game\s\d+:\s`)
	re2 := regexp.MustCompile(`\d+`)
	var gameId int
	var sum int

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		gameId, _ = strconv.Atoi(re2.FindString(line))
		line = re.ReplaceAllString(line, "")
		draws := strings.Split(line, "; ")
		isCorrect := true

		for _, draw := range draws {
			results := strings.Split(draw, ", ")

			for _, result := range results {
				if strings.Contains(result, "red") {
					if !isWithinLimit(result, guessesLimit.red) {
						isCorrect = false
						break
					}
				}

				if strings.Contains(result, "green") {
					if !isWithinLimit(result, guessesLimit.green) {
						isCorrect = false
						break
					}
				}

				if strings.Contains(result, "blue") {
					if !isWithinLimit(result, guessesLimit.blue) {
						isCorrect = false
						break
					}
				}
			}
		}
		if isCorrect {
			sum += gameId
		}
	}

	fmt.Println(sum)
}
