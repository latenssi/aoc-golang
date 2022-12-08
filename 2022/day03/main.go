package main

import (
	"strings"
	"unicode"
)

func Day(input string, part int) int {
	sum := 0
	if part == 1 {
		for _, sack := range strings.Split(input, "\n") {
			if len(sack) == 0 {
				continue
			}
			contents := ""
			for i, r := range sack {
				contains := strings.ContainsRune(contents, r)
				firstHalf := i < len(sack)/2
				if !contains && firstHalf {
					contents += string(r)
				} else if contains && !firstHalf {
					// We found a duplicate, calculate priority
					sum += priority(r)
					break
				}
			}
		}
	}
	if part == 2 {
		var counts map[rune]int
		for i, sack := range strings.Split(input, "\n") {
			if len(sack) == 0 {
				continue
			}
			// First sack, initialize map
			if i%3 == 0 {
				counts = make(map[rune]int)
			}
			for _, r := range sack {
				// Increment count only once per sack
				if counts[r] == i%3 {
					counts[r] += 1
				}
			}
			// Last sack, calculate priority
			if i%3 == 2 {
				for r, count := range counts {
					if count == 3 {
						sum += priority(r)
					}
				}
			}
		}
	}

	return sum
}

func priority(r rune) int {
	if unicode.IsUpper(r) {
		return int(r) - 64 + 26
	}
	return int(r) - 96
}
