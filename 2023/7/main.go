package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cardList := "AKQJT98765432"
	var hands []Hand

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		cards := parts[0]
		bid, _ := strconv.Atoi(parts[1])

		cardCount := intoCharacterMap(cards)
		sortedCards := sortCards(cardCount, cardList)
		hands = append(hands, Hand{sortedCards, bid, cards})
	}

	sort.SliceStable(hands, func(i, j int) bool {
		return less(hands[i], hands[j], cardList)
	})

	result := 0
	for i, hand := range hands {
		result += (len(hands) - i) * hand.bid
	}

	fmt.Println("Result:", result)
}

type Hand struct {
	cards []CardCount
	bid   int
	raw   string
}

type CardCount struct {
	card  rune
	count int
}

func intoCharacterMap(word string) map[rune]int {
	cardCount := make(map[rune]int)
	for _, c := range word {
		cardCount[c]++
	}
	return cardCount
}

func sortCards(cardCount map[rune]int, cardList string) []CardCount {
	var sorted []CardCount
	for card, count := range cardCount {
		sorted = append(sorted, CardCount{card, count})
	}
	sort.Slice(sorted, func(i, j int) bool {
		ci := strings.IndexRune(cardList, sorted[i].card)
		cj := strings.IndexRune(cardList, sorted[j].card)
		if sorted[i].count == sorted[j].count {
			return ci < cj
		}
		return sorted[i].count > sorted[j].count
	})
	return sorted
}

func less(a, b Hand, cardList string) bool {
	if len(a.cards) != len(b.cards) {
		return len(a.cards) < len(b.cards)
	}
	for i := range a.cards {
		ac, bc := a.cards[i], b.cards[i]
		if ac.count != bc.count {
			return ac.count > bc.count
		}
		ai := strings.IndexRune(cardList, ac.card)
		bi := strings.IndexRune(cardList, bc.card)
		if ai != bi {
			return ai < bi
		}
	}
	return a.raw < b.raw
}
