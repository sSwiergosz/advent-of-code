package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	scanner := bufio.NewScanner(file)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func visibleFromLeft(currentHeight int, currentHeightIndex int, row string) bool {
	for i := 0; i < currentHeightIndex; i++ {
		currentHeightNum, _ := strconv.Atoi(string(row[i]))

		if currentHeightNum >= currentHeight {
			return false
		}
	}

	return true
}

func visibleFromRight(currentHeight int, currentHeightIndex int, row string) bool {
	for i := len(row) - 1; i > currentHeightIndex; i-- {
		currentElement, _ := strconv.Atoi(string(row[i]))

		if currentHeight <= currentElement {
			return false
		}
	}

	return true
}

func visibleInRow(currentHeight int, i int, row string) bool {
	return visibleFromLeft(currentHeight, i, row) || visibleFromRight(currentHeight, i, row)
}

func visibleFromTop(rowIndex int, lines []string, columnIndex int, currentHeight int) bool {
	for i := 0; i < rowIndex; i++ {
		currentElementInColumn, _ := strconv.Atoi(string(lines[i][columnIndex]))
		if currentElementInColumn >= currentHeight {
			return false
		}
	}

	return true
}

func visibleFromBottom(rowIndex int, lines []string, columnIndex int, currentHeight int) bool {
	for i := len(lines) - 1; i > rowIndex; i-- {
		currentElementInColumn, _ := strconv.Atoi(string(lines[i][columnIndex]))

		if currentHeight <= currentElementInColumn {
			return false
		}
	}

	return true
}

func visibleInColumn(rowIndex int, lines []string, columnIndex int, currentHeight int) bool {
	return visibleFromTop(rowIndex, lines, columnIndex, currentHeight) || visibleFromBottom(rowIndex, lines, columnIndex, currentHeight)
}

func main() {
	lines, _ := readLines("input.txt")
	treesCount := 0

	for rowIndex, row := range lines {
		for i := 0; i < len(row); i++ {
			currentHeight, _ := strconv.Atoi(string(row[i]))

			isVisibleInRow := visibleInRow(currentHeight, i, row)
			isVisibleInColumn := visibleInColumn(rowIndex, lines, i, currentHeight)

			if isVisibleInRow || isVisibleInColumn {
				treesCount++
			}
		}
	}

	fmt.Println("Total visible trees:", treesCount)
}
