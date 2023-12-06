package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func processLine(regex *regexp.Regexp, line string, ints []int) []int {
	values := strings.Fields(regex.ReplaceAllString(line, ""))
	for _, str := range values {
		num, _ := strconv.Atoi(str)
		ints = append(ints, num)
	}
	return ints
}

func calculateSum(times, distances []int) int {
	sum := 0
	for i, time := range times {
		ways := 0
		for j := 0; j <= time; j++ {
			speed := j
			result := (time - speed) * speed
			if result > distances[i] {
				ways++
			}
		}
		if sum == 0 {
			sum += ways
		} else {
			sum *= ways
		}
	}
	return sum
}

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer readFile.Close()

	reTimes := regexp.MustCompile(`Time:\s`)
	reDistances := regexp.MustCompile(`Distance:\s`)

	var intTimes, intDistances []int

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.HasPrefix(line, "Time:") {
			intTimes = processLine(reTimes, line, intTimes)
		} else if strings.HasPrefix(line, "Distance:") {
			intDistances = processLine(reDistances, line, intDistances)
		}
	}

	sum := calculateSum(intTimes, intDistances)
	fmt.Println(sum)

	// Part 2
	strTimes := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(intTimes)), ""), "[]")
	strDistances := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(intDistances)), ""), "[]")

	var part2TimesArray []int
	var part2DistancesArray []int

	part2Times, _ := strconv.Atoi(strTimes)
	part2Distances, _ := strconv.Atoi(strDistances)

	part2TimesArray = append(part2TimesArray, part2Times)
	part2DistancesArray = append(part2DistancesArray, part2Distances)

	sum2 := calculateSum(part2TimesArray, part2DistancesArray)

	fmt.Println(sum2)
}
