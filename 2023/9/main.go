package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var diffs [][]int
var sum int

func isZeroSlice(arr []int) bool {
	for _, el := range arr {
		if el != 0 {
			return false
		}
	}

	return true
}

func calcDiff(arr []int) {
	var diff []int

	for i, el := range arr {
		if i+1 == len(arr) {
			break
		}

		currVal := el
		nextVal := arr[i+1]

		diff = append(diff, nextVal-currVal)
	}

	if !isZeroSlice(diff) {
		diffs = append(diffs, diff)
		calcDiff(diff)
	} else {
		diffs = append(diffs, diff)
	}
}

func convertToNumeric(arr []string) []int {
	var resVal []int
	for _, el := range arr {
		numVal, _ := strconv.Atoi(el)
		resVal = append(resVal, numVal)
	}

	return resVal
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		l := strings.Split(line, " ")
		numL := convertToNumeric(l)

		diffs = append(diffs, numL)
		calcDiff(numL)
	}

	for i := len(diffs) - 1; i >= 0; i-- {
		if i-1 < 0 {
			sum += diffs[i][len(diffs[i])-1]
			break
		}
		if isZeroSlice(diffs[i-1]) {
			sum += diffs[i][len(diffs[i])-1]
			continue
		}
		res := diffs[i][len(diffs[i])-1] + diffs[i-1][len(diffs[i-1])-1]
		diffs[i-1] = append(diffs[i-1], res)
	}

	fmt.Println(sum)
}
