package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Mapping struct {
	DestinationStart int
	SourceStart      int
	Length           int
}

func (m *Mapping) String() string {
	return fmt.Sprintf("%v %v %v", m.DestinationStart, m.SourceStart, m.Length)
}

func (m *Mapping) Map(seed int) (int, bool) {
	if seed >= m.SourceStart && seed < m.SourceStart+m.Length {
		return m.DestinationStart + seed - m.SourceStart, true
	}
	return seed, false
}

func (m *Mapping) MapRange(src Range) ([]Range, bool) {
	// Overlap start
	os := int(math.Max(float64(src.Start), float64(m.SourceStart)))
	// Overlap end
	oe := int(math.Min(float64(src.End), float64(m.SourceStart+m.Length-1)))

	if os <= oe {
		ranges := []Range{{
			Start: m.DestinationStart + os - m.SourceStart,
			End:   m.DestinationStart + oe - m.SourceStart,
		}}
		if src.Start < os {
			ranges = append(ranges, Range{
				Start: src.Start,
				End:   os - 1,
			})
		}
		if src.End > oe {
			ranges = append(ranges, Range{
				Start: oe + 1,
				End:   src.End,
			})
		}
		return ranges, true
	}

	return []Range{src}, false
}

func MappingFromString(s string) (*Mapping, error) {
	parts := strings.Split(s, " ")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid map string: %v", s)
	}
	destStart, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}
	sourceStart, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}
	length, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, err
	}
	return &Mapping{
		DestinationStart: destStart,
		SourceStart:      sourceStart,
		Length:           length,
	}, nil
}

type Layer struct {
	Mappings []*Mapping
}

func (l *Layer) String() string {
	s := ""
	for _, m := range l.Mappings {
		s += fmt.Sprintf("%v\n", m.String())
	}
	return s
}

func (l *Layer) AddMap(line string) error {
	m, err := MappingFromString(line)
	if err != nil {
		return err
	}
	l.Mappings = append(l.Mappings, m)
	return nil
}

func (l *Layer) Map(seed int) int {
	for _, m := range l.Mappings {
		if seed, hit := m.Map(seed); hit {
			return seed
		}
	}
	return seed
}

func (l *Layer) MapRange(initial Range) []Range {
	for _, m := range l.Mappings {
		if ranges, hit := m.MapRange(initial); hit {
			newRanges := []Range{ranges[0]}
			for _, r := range ranges[1:] {
				newRanges = append(newRanges, l.MapRange(r)...)
			}
			return newRanges
		}
	}
	return []Range{initial}
}

type Range struct {
	Start int
	End   int
}

func Day(input string, part int) int {
	lines := strings.Split(input, "\n")

	seedsParts1 := strings.Split(lines[0], ":")
	seedsParts2 := strings.Split(strings.Trim(seedsParts1[1], " "), " ")
	seeds := []int{}
	for _, seedStr := range seedsParts2 {
		val, err := strconv.Atoi(seedStr)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, val)
	}

	allLayers := []Layer{}
	var currentLayer *Layer
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		if strings.HasSuffix(line, "map:") {
			if currentLayer != nil {
				allLayers = append(allLayers, *currentLayer)
			}
			currentLayer = &Layer{}
			continue
		}
		if err := currentLayer.AddMap(line); err != nil {
			panic(err)
		}
	}
	allLayers = append(allLayers, *currentLayer)

	if part == 1 {
		min := math.MaxInt
		for _, seed := range seeds {
			for _, layer := range allLayers {
				seed = layer.Map(seed)
			}
			if seed < min {
				min = seed
			}
		}
		return min
	}

	if part == 2 {
		seedRanges := map[int]int{}
		for i, val := range seeds {
			if i%2 == 0 {
				seedRanges[val] = seeds[i+1]
			}
		}

		min := math.MaxInt
		ranges := []Range{}
		for start, length := range seedRanges {
			ranges = []Range{{
				Start: start,
				End:   start + length,
			}}
			fmt.Println(ranges)
			for i, layer := range allLayers {
				layerRanges := ranges
				ranges = []Range{}
				for _, r := range layerRanges {
					ranges = append(ranges, layer.MapRange(r)...)
				}
				fmt.Printf("ranges after layer %d: %v\n", i, ranges)
			}
			sort.Slice(ranges, func(i, j int) bool {
				return ranges[i].Start < ranges[j].Start
			})
			fmt.Printf("ranges sorted: %v\n", ranges)
			if ranges[0].Start < min {
				min = ranges[0].Start
			}
		}

		return min
	}

	return 0
}
