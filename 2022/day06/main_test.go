package main

import (
	"testing"

	"github.com/latenssi/aoc-golang/util"
)

func Test(t *testing.T) {
	// testInput, err := util.ReadFile("./test.txt", nil)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	input, err := util.ReadFile("./input.txt", nil)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"test", "mjqjpqmgbljsphdztnvjfqwrcgsmlb", 1, 7},
		{"test", "bvwbjplbgvbhsrlpgdmjqwftvncz", 1, 5},
		{"test", "nppdvjthqldpwncqszvftbrmjlhg", 1, 6},
		{"test", "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 1, 10},
		{"test", "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 1, 11},
		{"actual", input, 1, 1356},
		{"test", "mjqjpqmgbljsphdztnvjfqwrcgsmlb", 2, 19},
		{"test", "bvwbjplbgvbhsrlpgdmjqwftvncz", 2, 23},
		{"test", "nppdvjthqldpwncqszvftbrmjlhg", 2, 23},
		{"test", "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 2, 29},
		{"test", "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 2, 26},
		{"actual", input, 2, 2564},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day(tt.input, tt.part); got != tt.want {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}
