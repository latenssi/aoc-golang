package main

import (
	"fmt"
	"strconv"
	"strings"
)

func readInputValuesSeparate(s string) ([]int, error) {
	result := []int{}
	for i, p := range strings.Split(s, " ") {
		if i == 0 {
			continue
		}
		s := strings.Trim(p, " ")
		if s == "" {
			continue
		}
		v, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		result = append(result, v)
	}
	return result, nil
}

func minWait(time, distance int) int {
	for wait := 0; wait < time/2; wait++ {
		if wait*(time-wait) > distance {
			return wait
		}
	}
	return 0
}

func Day(input string, part int) int {
	lines := strings.Split(input, "\n")
	if len(lines) < 2 {
		panic("Invalid input")
	}

	if part == 1 {
		times, err := readInputValuesSeparate(lines[0])
		if err != nil {
			panic(err)
		}
		distances, err := readInputValuesSeparate(lines[1])
		if err != nil {
			panic(err)
		}

		fmt.Println(times, distances)

		total := 1

		for i := 0; i < len(times); i++ {
			time := times[i]
			distance := distances[i]
			ways := time - minWait(time, distance)*2 + 1
			total *= ways
		}

		return total
	}

	if part == 2 {
		time, err := strconv.Atoi(strings.ReplaceAll(strings.TrimPrefix(lines[0], "Time:"), " ", ""))
		if err != nil {
			panic(err)
		}
		distance, err := strconv.Atoi(strings.ReplaceAll(strings.TrimPrefix(lines[1], "Distance:"), " ", ""))
		if err != nil {
			panic(err)
		}
		fmt.Println(time, distance)

		return time - minWait(time, distance)*2 + 1
	}

	return 0
}
