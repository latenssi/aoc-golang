package main

import (
	"strconv"
	"strings"

	"github.com/latenssi/aoc-golang/lib/filesystem"
)

const (
	spaceRequiredForUpdate = 30000000
	fileSystemSize         = 70000000
)

func Day(input string, part int) int {
	fs := filesystem.NewFileSystem(fileSystemSize)
	shell := filesystem.NewShell(fs)

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		if line[0] == '$' {
			// Handle commands
			if parts[1] == "cd" {
				shell.ChangeDir(parts[2])
			}
		} else {
			// Handle ls output
			if parts[0] == "dir" {
				// Add directory
				shell.MakeDir(parts[1])
			} else {
				// Add file
				size, err := strconv.Atoi(parts[0])
				if err != nil {
					panic(err)
				}
				shell.AddFile(parts[1], size)
			}
		}
	}

	if part == 1 {
		sum := 0
		for _, dir := range fs.AllDirs() {
			size := dir.Size()
			if size <= 100000 {
				sum += size
			}
		}
		return sum
	}

	if part == 2 {
		spaceNeeded := spaceRequiredForUpdate - fs.FreeSpace()
		min := fs.TotalSize()
		for _, dir := range fs.AllDirs() {
			size := dir.Size()
			if size >= spaceNeeded && size < min {
				min = size
			}
		}
		return min
	}

	return 0
}
