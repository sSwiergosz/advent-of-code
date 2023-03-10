package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile(name string) string {
	content, err := os.ReadFile(name)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	return string(content)
}

func allCharactersDifferent(str string) bool {
	for i := 0; i < len(str); i++ {
		contains := strings.Contains(str[i+1:], string(str[i]))

		if contains {
			return false
		}
	}

	return true
}

func main() {
	file := readFile("input.txt")

	for i := 0; i < len(file); i += 1 {
		if i+14 > len(file) {
			return
		}

		r := string(file[i : i+14])

		if allCharactersDifferent(r) {
			fmt.Print(i + 14)
			return
		}
	}
}
