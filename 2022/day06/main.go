package main

func FindEndIndexOfSequenceOfDifferentChars(input string, windowSize int) int {
	for i := 0; i < len(input)-windowSize; i++ {
		runeSet := make(map[rune]bool)
		for j := 0; j < windowSize; j++ {
			runeSet[rune(input[i+j])] = true
		}
		if len(runeSet) == windowSize {
			return i + windowSize
		}
	}
	return 0
}

func Day(input string, part int) int {
	if part == 1 {
		return FindEndIndexOfSequenceOfDifferentChars(input, 4)
	}
	if part == 2 {
		return FindEndIndexOfSequenceOfDifferentChars(input, 14)
	}
	return 0
}
