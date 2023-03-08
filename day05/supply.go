package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack []int

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

func main() {
	m := make(map[int][]string)
	lines, _ := readLines("input.txt")

	for i, l := range lines {
		if i < 8 {
			counter := 0
			for j := 1; j <= len(l); j += 4 {
				if string(l[j]) != " " {
					m[counter] = append(m[counter], string(l[j]))
				}
				counter++
			}
		}

		if i > 9 {
			s := strings.Split(l, " ")

			times, _ := strconv.Atoi(s[1])
			from, _ := strconv.Atoi(s[3])
			to, _ := strconv.Atoi(s[5])

			for k := 1; k <= times; k++ {
				popped := m[from-1][0]
				m[from-1] = m[from-1][1:]

				m[to-1] = append([]string{popped}, m[to-1]...)
			}
		}
	}

	fmt.Print(m)
}
