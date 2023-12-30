package main

import (
	"fmt"
	"strings"
)

type Node struct {
	L   *Node
	R   *Node
	Val string
}

func (n *Node) Go(inst rune) *Node {
	switch inst {
	case 'R':
		return n.R
	case 'L':
		return n.L
	default:
		panic("unknown instruction")
	}
}

func (n *Node) String() string {
	return fmt.Sprintf("%s: (%s, %s)", n.Val, n.L.Val, n.R.Val)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func Day(input string, part int) int {
	lines := strings.Split(input, "\n")
	nodes := make(map[string]*Node)
	instructions := []rune(lines[0])
	for i := 2; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		parts := strings.Split(string(lines[i]), "=")
		thisV := strings.TrimSpace(parts[0])
		n := strings.Split(strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(parts[1]), "(", ""), ")", ""), ",")
		leftV := strings.TrimSpace(n[0])
		rightR := strings.TrimSpace(n[1])
		left, ok := nodes[leftV]
		if !ok {
			left = &Node{
				Val: leftV,
			}
			nodes[leftV] = left
		}
		right, ok := nodes[rightR]
		if !ok {
			right = &Node{
				Val: rightR,
			}
			nodes[rightR] = right
		}
		if this, ok := nodes[thisV]; !ok {
			nodes[thisV] = &Node{
				L:   left,
				R:   right,
				Val: thisV,
			}
		} else {
			this.L = left
			this.R = right
		}
	}

	if part == 1 {
		pos := nodes["AAA"]
		i := 0
		for {
			pos = pos.Go(instructions[i%len(instructions)])
			if pos.Val == "ZZZ" {
				return i + 1
			}
			i++
		}
	}

	if part == 2 {
		positions := []*Node{}
		for _, node := range nodes {
			if node.Val[2] == 'A' {
				positions = append(positions, node)
			}
		}
		firstZs := []int{}
		for _, node := range positions {
			i := 0
			for {
				node = node.Go(instructions[i%len(instructions)])
				if node.Val[2] == 'Z' {
					break
				}
				i++
			}
			firstZs = append(firstZs, i+1)
		}

		if len(firstZs) == 1 {
			return firstZs[0]
		}

		return LCM(firstZs[0], firstZs[1], firstZs...)
	}

	return 0
}
