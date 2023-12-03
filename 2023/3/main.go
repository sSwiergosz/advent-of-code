package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var lines []string
var sum int

func isValidSymbol(char rune) bool {
	return char != '.' && !unicode.IsDigit(char)
}

func hasAdjacentSymbol(row, col int) bool {
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i >= 0 && i < len(lines) && j >= 0 && j < len(lines[i]) && !(i == row && j == col) {
				if isValidSymbol(rune(lines[i][j])) {
					return true
				}
			}
		}
	}
	return false
}

func extractNumber(row, col int) (string, bool) {
	var num string

	// Check horizontally in rows
	for j := col; j < len(lines[row]) && unicode.IsDigit(rune(lines[row][j])); j++ {
		num += string(lines[row][j])
	}

	// // If there's a digit to the left, check if it's part of the same number
	for j := col - 1; j >= 0 && unicode.IsDigit(rune(lines[row][j])); j-- {
		num = string(lines[row][j]) + num
	}

	return num, len(num) > 0
}

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
	}

	isVisitedNumber := false

	for rowIndex, line := range lines {
		for columnIndex, char := range line {
			if unicode.IsDigit(char) && hasAdjacentSymbol(rowIndex, columnIndex) {
				num, valid := extractNumber(rowIndex, columnIndex)

				if valid && !isVisitedNumber {
					numValue, _ := strconv.Atoi(num)
					fmt.Println(numValue)
					sum += numValue
					isVisitedNumber = true
				}
			} else {
				isVisitedNumber = false
			}
		}
	}

	fmt.Println(sum)
}
