package main

import (
	"testing"

	"github.com/latenssi/aoc-golang/util"
)

func Test(t *testing.T) {
	testInput, err := util.ReadFile("./test.txt", nil)
	if err != nil {
		t.Fatal(err)
	}
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
		{"test", testInput, 1, 114},
		{"actual", input, 1, 1868368343},
		// {"test", testInput, 2, 2},
		// {"actual", input, 2, 1022},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day(tt.input, tt.part); got != tt.want {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}
