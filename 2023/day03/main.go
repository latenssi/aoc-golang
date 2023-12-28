package main

import (
	"strconv"
	"strings"
	"unicode"
)

type Number struct {
	Value    int
	ValueStr string
	Indexes  []int
}

func (n *Number) ParseValue() error {
	val, err := strconv.Atoi(n.ValueStr)
	if err != nil {
		return err
	}
	n.Value = val
	return nil
}

func (n *Number) AdjacentToIndex(index, gridW int) bool {
	for _, i := range n.Indexes {
		if i == index || i == index-1 || i == index+1 ||
			i == index-gridW || i == index-gridW-1 || i == index-gridW+1 ||
			i == index+gridW || i == index+gridW-1 || i == index+gridW+1 {
			return true
		}
	}

	return false
}

func Day(input string, part int) int {
	gridW := 0
	allNumbers := []Number{}
	var num *Number
	symbolIndexes := []int{}
	gearIndexes := []int{}
	for l, line := range strings.Split(input, "\n") {
		if gridW == 0 {
			gridW = len(line)
		}
		for i, r := range line {
			if unicode.IsDigit(r) {
				if num == nil {
					num = &Number{}
				}
				num.Indexes = append(num.Indexes, i+(l*gridW))
				num.ValueStr += string(r)
			} else {
				if num != nil {
					if err := num.ParseValue(); err != nil {
						panic(err)
					}
					allNumbers = append(allNumbers, *num)
					num = nil
				}
				if r != '.' {
					symbolIndexes = append(symbolIndexes, i+(l*gridW))
					if r == '*' {
						gearIndexes = append(gearIndexes, i+(l*gridW))
					}
				}
			}
		}
	}

	if part == 1 {
		sum := 0
		for _, n := range allNumbers {
			for _, i := range symbolIndexes {
				if n.AdjacentToIndex(i, gridW) {
					sum += n.Value
					break // only add once
				}
			}
		}
		return sum
	}

	if part == 2 {
		sum := 0
		for _, i := range gearIndexes {
			numbers := []Number{}
			for _, n := range allNumbers {
				if n.AdjacentToIndex(i, gridW) {
					numbers = append(numbers, n)
				}
			}
			if len(numbers) == 2 {
				sum += numbers[0].Value * numbers[1].Value
			}
		}
		return sum
	}

	return 0
}
