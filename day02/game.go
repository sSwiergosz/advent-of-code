package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func pointForPlay(play string) int {
	if play == "Rock" {
		return 1
	}

	if play == "Paper" {
		return 2
	}

	return 3
}

func roundPoints(opp string, player string) int {
	score := pointForPlay(player)

	if opp == player {
		return score + 3
	}

	if opp == "Rock" && player == "Scissors" || opp == "Paper" && player == "Rock" || opp == "Scissors" && player == "Paper" {
		return score + 0
	}

	if player == "Rock" && opp == "Scissors" || player == "Paper" && opp == "Rock" || player == "Scissors" && opp == "Paper" {
		return score + 6
	}

	return score
}

func main() {
	oppPlays := map[string]string{"A": "Rock", "B": "Paper", "C": "Scissors"}
	yourPlays := map[string]string{"X": "Rock", "Y": "Paper", "Z": "Scissors"}
	lines, _ := readLines("input.txt")

	var score int = 0

	for _, line := range lines {
		play := strings.Fields(line)

		opp := play[0]
		player := play[1]
		score += roundPoints(oppPlays[opp], yourPlays[player])
	}

	fmt.Printf("Your total score is: %v", score)
}
