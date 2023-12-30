package main

import (
	"testing"

	"github.com/latenssi/aoc-golang/util"
)

func Test(t *testing.T) {
	testInput1, err := util.ReadFile("./test1.txt", nil)
	if err != nil {
		t.Fatal(err)
	}
	testInput2, err := util.ReadFile("./test2.txt", nil)
	if err != nil {
		t.Fatal(err)
	}
	testInput3, err := util.ReadFile("./test3.txt", nil)
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
		{"test1", testInput1, 1, 2},
		{"test2", testInput2, 1, 6},
		{"actual", input, 1, 20659},
		{"test1", testInput1, 2, 2},
		{"test2", testInput2, 2, 6},
		{"test3", testInput3, 2, 6},
		{"actual", input, 2, 15690466351717},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day(tt.input, tt.part); got != tt.want {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}
