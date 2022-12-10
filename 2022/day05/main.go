package main

import (
	"regexp"
	"strconv"

	"github.com/latenssi/aoc-golang/util"
)

type Stack struct {
	items []rune
}

func (s *Stack) Len() int {
	return len(s.items)
}

func (s *Stack) Insert(item rune, index int) {
	s.items = append(s.items, 0)
	copy(s.items[index+1:], s.items[index:])
	s.items[index] = item
}

func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() rune {
	if s.Len() == 0 {
		return 0
	}
	item := s.items[s.Len()-1]
	s.items = s.items[:s.Len()-1]
	return item
}

func Move(from, to *Stack, count int, asStack bool) {
	if !asStack {
		for i := 0; i < count; i++ {
			to.Push(from.Pop())

		}
	} else {
		tmp := Stack{}
		Move(from, &tmp, count, false)
		Move(&tmp, to, count, false)
	}
}

const spotWidth = 4

func Day(input string, part int) string {
	stacks := make([]Stack, 9)
	r := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	for _, line := range util.Lines(input) {
		if len(line) == 0 {
			continue
		}

		if line[0] == 'm' {
			// Parse operation
			op := r.FindStringSubmatch(line)
			if len(op) != 4 {
				panic("invalid operation")
			}

			count, err := strconv.Atoi(op[1])
			if err != nil {
				panic(err)
			}

			from, err := strconv.Atoi(op[2])
			if err != nil {
				panic(err)
			}

			to, err := strconv.Atoi(op[3])
			if err != nil {
				panic(err)
			}

			Move(&stacks[from-1], &stacks[to-1], count, part == 2)
		} else {
			// Parse initial state
			for i := 0; i < len(line); i++ {
				if line[i] == '[' {
					stacks[i/spotWidth].Insert(rune(line[i+1]), 0)
				}
			}
		}
	}

	result := ""
	for _, s := range stacks {
		if s.Len() == 0 {
			continue
		}
		result += string(s.Pop())
	}

	return result
}
