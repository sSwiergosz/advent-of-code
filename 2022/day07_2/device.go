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

var TOTAL_DISK_SPACE = 70000000
var MIN_DISK_SPACE = 30000000

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

	totalSum := dirs["."]

	fmt.Print("Your total sum: ", totalSum, "\n")

	diff := TOTAL_DISK_SPACE - totalSum

	fmt.Print("Your free space on disk: ", diff, "\n")

	needed := MIN_DISK_SPACE - diff

	fmt.Print("Size you need to freed: ", needed, "\n")

	var candidates []int
	for _, v := range dirs {
		if v > needed {
			candidates = append(candidates, v)
		}
	}

	smallest := candidates[0]

	for _, v := range candidates {
		if v < smallest {
			smallest = v
		}
	}

	fmt.Print("You need to delete file with size: ", smallest)
}
