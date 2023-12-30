package main

import (
	"strconv"
	"strings"
)

func getLevelsRecursive(s []int, levels [][]int) [][]int {
	if len(s) <= 1 {
		return levels
	}
	freq := make(map[int]int)
	for _, v := range s {
		freq[v]++
	}
	if len(freq) == 1 {
		return levels
	}
	next := []int{}
	for i := 0; i < len(s)-1; i++ {
		next = append(next, s[i+1]-s[i])
	}
	return getLevelsRecursive(next, append([][]int{next}, levels...))
}

func Day(input string, part int) int {
	lines := strings.Split(input, "\n")
	allVals := [][]int{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		vals2 := []int{}
		for _, s := range parts {
			v, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			vals2 = append(vals2, v)
		}
		allVals = append(allVals, vals2)
	}

	if part == 1 {
		total := 0
		for i := 0; i < len(allVals); i++ {
			levels := getLevelsRecursive(allVals[i], [][]int{})
			add := 0
			for _, l := range levels {
				add = l[len(l)-1] + add
			}
			total += allVals[i][len(allVals[i])-1] + add
		}
		return total
	}

	if part == 2 {
		total := 0
		for i := 0; i < len(allVals); i++ {
			levels := getLevelsRecursive(allVals[i], [][]int{})
			sub := 0
			for _, l := range levels {
				sub = l[0] - sub
			}
			total += allVals[i][0] - sub
		}
		return total
	}

	return 0
}
