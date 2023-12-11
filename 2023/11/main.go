package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid []string

	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	emptyRows, emptyCols := findEmptyRowsAndCols(grid)
	points := findGalaxies(grid)
	total := calculateTotalDistance(points, emptyRows, emptyCols)

	fmt.Println(total)
}

func findEmptyRowsAndCols(grid []string) ([]int, []int) {
	var emptyRows, emptyCols []int
	rowLen := len(grid[0])

	for r, row := range grid {
		if strings.Count(row, ".") == rowLen {
			emptyRows = append(emptyRows, r)
		}
	}

	for c := 0; c < rowLen; c++ {
		colEmpty := true
		for _, row := range grid {
			if row[c] != '.' {
				colEmpty = false
				break
			}
		}
		if colEmpty {
			emptyCols = append(emptyCols, c)
		}
	}
	return emptyRows, emptyCols
}

func findGalaxies(grid []string) [][2]int {
	var points [][2]int
	for r, row := range grid {
		for c, ch := range row {
			if ch == '#' {
				points = append(points, [2]int{r, c})
			}
		}
	}
	return points
}

func calculateTotalDistance(points [][2]int, emptyRows, emptyCols []int) int {
	total := 0
	for i, point1 := range points {
		for _, point2 := range points[:i] {
			total += distance(point1, point2, emptyRows, emptyCols)
		}
	}
	return total
}

func distance(p1, p2 [2]int, emptyRows, emptyCols []int) int {
	dist := 0
	for r := min(p1[0], p2[0]); r < max(p1[0], p2[0]); r++ {
		if slices.Contains(emptyRows, r) {
			dist += 2
		} else {
			dist++
		}
	}
	for c := min(p1[1], p2[1]); c < max(p1[1], p2[1]); c++ {
		if slices.Contains(emptyCols, c) {
			dist += 2
		} else {
			dist++
		}
	}
	return dist
}
