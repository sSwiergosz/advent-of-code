package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isNumber(v string) bool {
	if _, err := strconv.Atoi(v); err == nil {
		return true
	}

	return false
}

func filterNumbers(vals []string) []string {

	n := 0
	for _, val := range vals {
		if isNumber(val) {
			vals[n] = val
			n++
		}
	}

	vals = vals[:n]

	return vals
}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	sum := 0

	for fileScanner.Scan() {
		result := ""
		strArray := strings.Split(fileScanner.Text(), "")
		filteredArray := filterNumbers(strArray)

		switch l := len(filteredArray); l {
		case 0:
			result = "0"
		case 1:
			result = filteredArray[0] + filteredArray[0]
		default:
			result = filteredArray[0] + filteredArray[len(filteredArray)-1]
		}

		n, _ := strconv.Atoi(result)
		sum += n
	}

	fmt.Println(sum)

	readFile.Close()
}
