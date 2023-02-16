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

func calculatePlay(opp string, outcome string) string {
	if outcome == "X" && opp == "Rock" || outcome == "Y" && opp == "Scissors" || outcome == "Z" && opp == "Paper" {
		return "Scissors"
	}

	if outcome == "X" && opp == "Scissors" || outcome == "Y" && opp == "Paper" || outcome == "Z" && opp == "Rock" {
		return "Paper"
	}

	if outcome == "X" && opp == "Paper" || outcome == "Y" && opp == "Rock" || outcome == "Z" && opp == "Scissors" {
		return "Rock"
	}

	return ""
}

func pointForPlay(play string) int {
	switch play {
	case "Rock":
		return 1
	case "Paper":
		return 2
	default:
		return 3
	}
}

func roundPoints(opp string, outcome string) int {
	player := calculatePlay(opp, outcome)
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
	lines, _ := readLines("input.txt")

	var score int = 0

	for _, line := range lines {
		play := strings.Fields(line)

		opp := play[0]
		player := play[1]
		score += roundPoints(oppPlays[opp], player)
	}

	fmt.Printf("Your total score is: %v", score)
}
