package main

import (
	"strings"
)

var beats = map[rune]rune{
	'A': 'C',
	'B': 'A',
	'C': 'B',
}

var loses = map[rune]rune{
	'A': 'B',
	'B': 'C',
	'C': 'A',
}

var part_1_mapping = map[rune]rune{
	'X': 'A',
	'Y': 'B',
	'Z': 'C',
}

var scores = map[rune]int{
	'A': 1,
	'B': 2,
	'C': 3,
}

func Day(input string, part int) int {
	score := 0
	for _, playStr := range strings.Split(input, "\n") {
		if len(playStr) == 0 {
			continue
		}
		play := []rune(playStr)
		their := play[0]
		mine := part_1_mapping[play[2]]

		if part == 2 {
			switch play[2] {
			case 'X':
				// we need to lose
				mine = beats[their]
			case 'Y':
				// we need to draw
				mine = their
			case 'Z':
				// we need to win
				mine = loses[their]
			}
		}

		// Base score
		score += scores[mine]

		if their == mine {
			// Draw
			score += 3
		} else if beats[mine] == their {
			// Win
			score += 6
		}
	}

	return score
}
