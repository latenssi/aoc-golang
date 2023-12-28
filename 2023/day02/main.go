package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	Id      int
	Pickups []Pickup
}

func (g *Game) SetFromStr(s string) error {
	s = strings.Trim(s, " ")
	if s == "" {
		return nil
	}
	s = strings.Trim(s, "Game ")
	id, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	g.Id = id
	g.Pickups = []Pickup{}
	return nil
}

func (g *Game) AddPickups(s string) error {
	for _, round := range strings.Split(s, ";") {
		round = strings.Trim(round, " ")
		if round == "" {
			return nil
		}
		for _, p := range strings.Split(round, ",") {
			p = strings.Trim(p, " ")
			if p == "" {
				continue
			}
			parts := strings.Split(p, " ")
			if len(parts) != 2 {
				return fmt.Errorf("invalid pickup: %s", p)
			}
			count, err := strconv.Atoi(parts[0])
			if err != nil {
				return err
			}
			color := parts[1]
			g.Pickups = append(g.Pickups, Pickup{Color: color, Count: count})
		}
	}
	return nil
}

func (g *Game) MaxSeenCounts() (int, int, int) {
	red := 0
	green := 0
	blue := 0
	for _, p := range g.Pickups {
		if p.Color == "red" && p.Count > red {
			red = p.Count
		}
		if p.Color == "green" && p.Count > green {
			green = p.Count
		}
		if p.Color == "blue" && p.Count > blue {
			blue = p.Count
		}
	}
	return red, green, blue
}

type Pickup struct {
	Color string
	Count int
}

func Day(input string, part int) int {
	allGames := []Game{}
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}
		game := Game{}
		if err := game.SetFromStr(parts[0]); err != nil {
			panic(err)
		}
		if err := game.AddPickups(parts[1]); err != nil {
			panic(err)
		}
		allGames = append(allGames, game)
	}

	if part == 1 {
		sum := 0
		for _, game := range allGames {
			// 12 red cubes, 13 green cubes, and 14 blue cubes
			red, green, blue := game.MaxSeenCounts()
			if red <= 12 && green <= 13 && blue <= 14 {
				sum += game.Id
			}
		}
		return sum
	}

	if part == 2 {
		sum := 0
		for _, game := range allGames {
			red, green, blue := game.MaxSeenCounts()
			sum += red * green * blue
		}
		return sum
	}

	return 0
}
