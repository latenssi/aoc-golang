package main

import (
	"strings"

	"github.com/latenssi/aoc-golang/lib/math"
	"github.com/latenssi/aoc-golang/util"
)

type Grid struct {
	Width  int
	Height int
	raw    []int
}

func NewGrid(w, h int) *Grid {
	return &Grid{
		Width:  w,
		Height: h,
		raw:    make([]int, 0, w*h),
	}
}

func (g *Grid) Add(i int) {
	g.raw = append(g.raw, i)
}

func (g *Grid) At(x, y int) int {
	return g.raw[y*g.Width+x]
}

type TreeGrid struct {
	*Grid
}

func NewTreeGrid(w, h int) *TreeGrid {
	return &TreeGrid{
		NewGrid(w, h),
	}
}

func (g *TreeGrid) TreeVisibleAt(x, y int) bool {
	if x == 0 || y == 0 {
		return true
	}

	if x == g.Width-1 || y == g.Height-1 {
		return true
	}

	tree := g.At(x, y)

	xLeftOffset := 0
	xRightOffset := 0
	yUpOffset := 0
	yDownOffset := 0

	iMax := math.Max(g.Width, g.Height)

	for i := 0; i < iMax; i++ {
		xLeft := math.Max(0, xLeftOffset+1)
		xRight := math.Max(0, xRightOffset+1)
		yUp := math.Max(0, yUpOffset+1)
		yDown := math.Max(0, yDownOffset+1)

		if g.At(xLeft, y) >= tree {
			xLeft--
		}

		if g.At(xRight, y) >= tree {
			xRight--
		}

		if g.At(x, yUp) >= tree {
			yUp--
		}

		if g.At(x, yDown) >= tree {
			yDown--
		}

	}

	// Check left
	for x2 := x - 1; x2 >= 0; x2-- {
		if g.At(x2, y) >= tree {
			break
		} else if x2 == 0 {
			return true
		}
	}

	// Check right
	for x2 := x + 1; x2 < g.Width; x2++ {
		if g.At(x2, y) >= tree {
			break
		} else if x2 == g.Width-1 {
			return true
		}
	}

	// Check up
	for y2 := y - 1; y2 >= 0; y2-- {
		if g.At(x, y2) >= tree {
			break
		} else if y2 == 0 {
			return true
		}
	}

	// Check down
	for y2 := y + 1; y2 < g.Height; y2++ {
		if g.At(x, y2) >= tree {
			break
		} else if y2 == g.Height-1 {
			return true
		}
	}

	return false
}

func Day(input string, part int) int {
	input, err := util.RemoveEmptyLines(input)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(input, "\n")

	g := NewTreeGrid(len(lines[0]), len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}

		for _, c := range line {
			g.Add(int(c - '0'))
		}
	}

	count := 0
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			// fmt.Print(g.At(x, y), g.TreeVisibleAt(x, y), " ")
			if g.TreeVisibleAt(x, y) {
				count++
			}
		}
		// fmt.Println()
	}

	return count
}
