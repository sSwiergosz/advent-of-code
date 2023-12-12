package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func count(cfg string, nums []int) int {
	if cfg == "" {
		if len(nums) == 0 {
			return 1
		}
		return 0
	}

	if len(nums) == 0 {
		for _, char := range cfg {
			if char == '#' {
				return 0
			}
		}
		return 1
	}

	result := 0

	if cfg[0] == '.' || cfg[0] == '?' {
		result += count(cfg[1:], nums)
	}

	if cfg[0] == '#' || cfg[0] == '?' {
		if nums[0] <= len(cfg) && !strings.Contains(cfg[:nums[0]], ".") &&
			(nums[0] == len(cfg) || (nums[0] < len(cfg) && cfg[nums[0]] != '#')) {
			var newCfg string
			if nums[0] < len(cfg) {
				newCfg = cfg[nums[0]+1:]
			} else {
				newCfg = ""
			}
			result += count(newCfg, nums[1:])
		}
	}

	return result
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		cfg := parts[0]
		numsStr := strings.Split(parts[1], ",")

		nums := make([]int, len(numsStr))
		for i, n := range numsStr {
			num, err := strconv.Atoi(n)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Invalid number: %v\n", err)
				os.Exit(1)
			}
			nums[i] = num
		}

		total += count(cfg, nums)
	}

	fmt.Println(total)
}
