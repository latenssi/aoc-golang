package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	Cards []int
	Bid   int
	Type  int
}

const (
	HighCard = iota + 1
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var cardsPart1 = []rune{
	' ', ' ', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A',
}

var cardsPart2 = []rune{
	' ', 'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', ' ', 'Q', 'K', 'A',
}

func parseHandType(h []int, handleJoker bool) (int, error) {
	freq := map[int]int{}
	jokerCount := 0
	for _, c := range h {
		freq[c]++
		if handleJoker && c == 1 {
			jokerCount++
		}
	}
	if len(freq) == 1 { // All cards have the same label
		return FiveOfAKind, nil
	}
	if len(freq) == 2 { // Hand has cards of two distinct labels
		if jokerCount > 0 {
			return FiveOfAKind, nil
		}
		for _, v := range freq {
			if v == 4 {
				return FourOfAKind, nil
			}
		}
		return FullHouse, nil
	}
	if len(freq) == 3 { // Hand has cards of three distinct labels
		if jokerCount >= 2 {
			return FourOfAKind, nil
		}
		for _, v := range freq {
			if v == 3 {
				if jokerCount == 1 {
					return FourOfAKind, nil
				}
				return ThreeOfAKind, nil
			}
		}
		if jokerCount > 0 {
			return FullHouse, nil
		}
		return TwoPairs, nil
	}
	if len(freq) == 4 { // Hand has cards of four distinct labels
		if jokerCount > 0 {
			return ThreeOfAKind, nil
		}
		return OnePair, nil
	}
	if len(freq) == 5 { // Hand has cards of five distinct labels
		if jokerCount > 0 {
			return OnePair, nil
		}
		return HighCard, nil
	}
	return 0, fmt.Errorf("invalid hand type: %v", h)
}

func Day(input string, part int) int {
	fmt.Println(part)
	hands := []Hand{}

	cards := cardsPart1
	if part == 2 {
		cards = cardsPart2
	}

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			panic("Invalid input")
		}
		hCards := []int{}
		for _, c := range parts[0] {
			for i, v := range cards {
				if c == v {
					hCards = append(hCards, i)
					break
				}
			}
		}
		hBid, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		hType, err := parseHandType(hCards, part == 2)
		if err != nil {
			panic(err)
		}
		hands = append(hands, Hand{hCards, hBid, hType})
	}

	sort.Slice(hands, func(i, j int) bool {
		// Sort by type
		if hands[i].Type < hands[j].Type {
			return true
		}
		if hands[i].Type > hands[j].Type {
			return false
		}
		// Sort by each card
		for k := 0; k < len(hands[i].Cards); k++ {
			if hands[i].Cards[k] < hands[j].Cards[k] {
				return true
			}
			if hands[i].Cards[k] > hands[j].Cards[k] {
				return false
			}
		}
		return true
	})

	total := 0
	for i, h := range hands {
		total += h.Bid * (i + 1)
	}

	return total
}
