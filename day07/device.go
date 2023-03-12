package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var path []string
var dirs map[string]int

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	path = append(path, ".")
	dirs = make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Split(line, " ")

		if args[1] == "ls" {
			continue
		}

		if args[0] == "$" {
			if args[1] == "cd" {
				if args[2] == "/" {
					path = path[:1]
				} else if args[2] == ".." {
					path = path[:len(path)-1]
				} else {
					path = append(path, args[2])
				}
			}
		} else {
			if args[0] == "dir" {

			} else {
				size, _ := strconv.Atoi(args[0])
				pathKey := strings.Join(path, "/")

				dirs[pathKey] += size

				if len(path) > 0 {
					for i := len(path) - 1; i > 0; i-- {
						parentKey := strings.Join(path[:i], "/")

						dirs[parentKey] += size
					}
				}
			}
		}
	}

	var totalSum int
	for _, v := range dirs {
		if v < 100000 {
			totalSum += v
		}
	}

	fmt.Print(totalSum)
}
