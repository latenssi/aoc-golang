package main

import (
	"strings"
	"unicode"
)

func Day(input string, part int) int {
	sum := 0
	var counts map[rune]int

	for sackIdx, sack := range strings.Split(input, "\n") {
		if len(sack) == 0 {
			continue
		}

		if part == 1 {
			counts = make(map[rune]int)
			for charIdx, r := range sack {
				if charIdx < len(sack)/2 {
					// First half, increment count
					counts[r] = 1
				} else if counts[r] == 1 {
					// We found a duplicate, calculate priority
					sum += priority(r)
					break
				}
			}
		}

		if part == 2 {

			// First sack, initialize map
			if sackIdx%3 == 0 {
				counts = make(map[rune]int)
			}

			for _, r := range sack {
				// Increment count only once per sack
				if counts[r] == sackIdx%3 {
					counts[r] += 1
				}
			}

			// Last sack, calculate priority
			if sackIdx%3 == 2 {
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
