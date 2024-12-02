package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var left []int
	var right []int
	var totalSum float64

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		strArray := strings.Split(fileScanner.Text(), " ")
		leftNum, err := strconv.Atoi(strArray[0])

		if err != nil {
			panic(err)
		}

		rightNum, err2 := strconv.Atoi(strArray[len(strArray)-1])

		if err2 != nil {
			panic(err2)
		}

		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	slices.Sort(left)
	slices.Sort(right)

	for i := 0; i < len(left); i++ {
		totalSum += math.Abs(float64(left[i]) - float64(right[i]))
	}

	fmt.Println(totalSum)

	readFile.Close()
}
