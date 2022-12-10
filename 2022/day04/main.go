package main

import (
	"strconv"
	"strings"

	"github.com/latenssi/aoc-golang/util"
)

type Pair struct {
	part1 PairPart
	part2 PairPart
}

func (p *Pair) FullyContains() bool {
	return p.part1.FullyContains(p.part2) || p.part2.FullyContains(p.part1)
}

func (p Pair) Ovelaps() bool {
	return p.FullyContains() || p.part1.Ovelaps(p.part2)
}

type PairPart []int

func (p PairPart) FullyContains(o PairPart) bool {
	return p[0] <= o[0] && p[1] >= o[1]
}

func (p PairPart) Ovelaps(o PairPart) bool {
	return (p[0] <= o[0] && p[1] >= o[0]) || (p[0] <= o[1] && p[1] >= o[1])
}

func Day(input string, part int) int {
	pairs := make([]Pair, 0)
	for _, line := range util.Lines(input) {
		if len(line) == 0 {
			continue
		}
		pp := make([]int, 0, 4)
		for _, parts := range strings.Split(line, ",") {
			for _, part := range strings.Split(parts, "-") {
				v, err := strconv.Atoi(part)
				if err != nil {
					panic(err)
				}
				pp = append(pp, v)
			}
		}
		pairs = append(pairs, Pair{PairPart{pp[0], pp[1]}, PairPart{pp[2], pp[3]}})
	}
	count := 0
	for _, p := range pairs {
		if part == 1 && p.FullyContains() || part == 2 && p.Ovelaps() {
			count++
		}
	}
	return count
}
