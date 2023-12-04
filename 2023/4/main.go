package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func trimSpacesFromStringSlice(slice []string) []string {
	var result []string

	for _, str := range slice {
		trimmedStr := strings.ReplaceAll(str, " ", "")
		result = append(result, trimmedStr)
	}

	return result
}

func main() {
	readFile, err := os.Open("input.txt")
	re := regexp.MustCompile(`Card\s\d+:\s`)
	var sum int

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		hits := 0
		cardValue := 0
		l := fileScanner.Text()
		cl := re.ReplaceAllString(l, "")
		coupon := strings.Split(cl, "|")
		winning := strings.Fields(strings.TrimSpace(coupon[0]))
		guesses := strings.Fields(strings.TrimSpace(coupon[1]))

		clearedWinning := trimSpacesFromStringSlice(winning)
		clearedGuesses := trimSpacesFromStringSlice(guesses)

		for _, g := range clearedGuesses {
			for _, w := range clearedWinning {
				if w == g {
					hits++
				}
			}
		}

		for i := 0; i < hits; i++ {
			if i == 0 {
				cardValue += 1
			} else {
				cardValue = cardValue * 2
			}
		}
		sum += cardValue
		hits = 0
		cardValue = 0
	}
	fmt.Println(sum)
}
