package main

import (
	"sort"
	"strconv"
	"strings"
)

func Day(input string, part int) int {
	values := []int{}

	for _, g := range strings.Split(input, "\n\n") {
		if len(g) == 0 {
			continue
		}
		gValue := 0
		for _, val := range strings.Split(g, "\n") {
			if len(val) == 0 {
				continue
			}
			v, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			gValue += v
		}
		values = append(values, gValue)
	}

	sort.Ints(values)

	if part == 1 {
		return values[len(values)-1]
	}

	if part == 2 {
		sum := 0
		for i := len(values) - 1; i >= len(values)-3; i-- {
			sum += values[i]
		}
		return sum
	}

	return 0
}
