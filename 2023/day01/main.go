package main

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func Day(input string, part int) int {
	if part == 2 {
		numbers := map[string]string{
			"one":   "1",
			"two":   "2",
			"three": "3",
			"four":  "4",
			"five":  "5",
			"six":   "6",
			"seven": "7",
			"eight": "8",
			"nine":  "9",
		}

		r := regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine`)
		input = r.ReplaceAllStringFunc(input, func(s string) string {
			if v, ok := numbers[s]; ok {
				return v + s[len(s)-1:]
			}
			return s
		})

		input = r.ReplaceAllStringFunc(input, func(s string) string {
			if v, ok := numbers[s]; ok {
				return v
			}
			return s
		})

	}

	var sum int64
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		var first, last string
		for _, r := range line {
			if unicode.IsDigit(r) {
				if first == "" {
					first = string(r)
				}
				last = string(r)
			}
		}
		v, err := strconv.ParseInt(first+last, 10, 64)
		if err != nil {
			panic(err)
		}
		sum += v
	}
	return int(sum)
}
