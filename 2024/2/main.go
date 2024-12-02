package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	safeCount := 0

	for fileScanner.Scan() {
		strArray := strings.Split(fileScanner.Text(), " ")
		baseFirst, _ := strconv.Atoi(strArray[0])
		baseSecond, _ := strconv.Atoi(strArray[1])
		isIncreasing := baseSecond > baseFirst
		isSafe := true

		for i := 0; i < len(strArray)-1; i++ {
			first, _ := strconv.Atoi(strArray[i])
			second, _ := strconv.Atoi(strArray[i+1])
			difference := math.Abs(float64(first) - float64(second))

			if difference < 1 || difference > 3 {
				isSafe = false
				break
			}

			if (isIncreasing && first > second) || (!isIncreasing && first < second) {
				isSafe = false
				break
			}
		}

		if isSafe {
			safeCount++
		}

	}

	fmt.Println(safeCount)

	defer readFile.Close()
}
