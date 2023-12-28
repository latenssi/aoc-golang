package main

import (
	"math"
	"strconv"
	"strings"
)

type Card struct {
	WinningNumbers []int
	PlayerNumbers  []int
}

func (c *Card) Matches() int {
	matches := 0
	for _, m := range c.WinningNumbers {
		for _, n := range c.PlayerNumbers {
			if m == n {
				matches++
			}
		}
	}
	return matches
}

func (c *Card) Points() int {
	return int(math.Pow(2, float64(c.Matches()-1)))
}

func CardFromInputLine(line string) (*Card, error) {
	card := &Card{}
	parts1 := strings.Split(line, ":")
	parts2 := strings.Split(strings.Trim(parts1[1], " "), "|")
	for _, n := range strings.Split(strings.Trim(parts2[0], " "), " ") {
		if n == "" {
			continue
		}
		v, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		card.WinningNumbers = append(card.WinningNumbers, v)
	}
	for _, n := range strings.Split(strings.Trim(parts2[1], " "), " ") {
		if n == "" {
			continue
		}
		v, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		card.PlayerNumbers = append(card.PlayerNumbers, v)
	}
	return card, nil
}

func Day(input string, part int) int {
	allCards := []*Card{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		card, err := CardFromInputLine(line)
		if err != nil {
			panic(err)
		}
		allCards = append(allCards, card)
	}

	if part == 1 {
		points := 0
		for _, card := range allCards {
			points += card.Points()
		}
		return points
	}

	if part == 2 {
		sum := 0
		cardCountsByIndex := map[int]int{}
		for i, card := range allCards {
			curCardCount, ok := cardCountsByIndex[i]
			if !ok {
				curCardCount = 1
				cardCountsByIndex[i] = 1
			}
			sum += curCardCount
			for j := 0; j < card.Matches(); j++ {
				oldCount := cardCountsByIndex[i+1+j]
				if oldCount == 0 {
					oldCount = 1
				}
				cardCountsByIndex[i+1+j] = oldCount + curCardCount
			}
		}
		return sum
	}

	return 0
}
