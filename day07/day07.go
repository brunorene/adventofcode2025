package day07

import (
	"adventofcode2025/common"
)

type Point struct {
	X, Y int
}

func countSplits(manifold []string, splitters map[Point]struct{}, start Point) map[Point]struct{} {
	if start.Y+1 >= len(manifold) {
		return make(map[Point]struct{})
	}

	if _, exists := splitters[start]; exists {
		var split1, split2 map[Point]struct{}

		var splitted bool

		if (start.X-1 >= 0 && manifold[start.Y][start.X-1] != '^') && (start.X-2 < 0 || manifold[start.Y][start.X-2] != '^') {
			split1 = countSplits(manifold, splitters, Point{start.X - 1, start.Y})

			splitted = true
		}

		if start.X+1 < len(manifold[0]) && manifold[start.Y][start.X+1] != '^' {
			split2 = countSplits(manifold, splitters, Point{start.X + 1, start.Y})

			splitted = true
		}

		if splitted {
			for k := range split1 {
				split2[k] = struct{}{}
			}

			split2[start] = struct{}{}

			return split2
		}
	}

	return countSplits(manifold, splitters, Point{start.X, start.Y + 1})
}

func Solve1(input string) int {
	splitters := make(map[Point]struct{})
	var manifold []string
	var start Point
	var row int

	for line := range common.ReadInput(input).ReadLines {
		manifold = append(manifold, line)

		for col, char := range line {
			switch char {
			case '^':
				splitters[Point{col, row}] = struct{}{}
			case 'S':
				start = Point{col, row}
			}
		}

		row++
	}

	return len(countSplits(manifold, splitters, start))
}

func Solve2(input string) int {
	splitters := make(map[Point]struct{})
	var manifold []string
	var start Point
	var row int

	for line := range common.ReadInput(input).ReadLines {
		manifold = append(manifold, line)

		for col, char := range line {
			switch char {
			case '^':
				splitters[Point{col, row}] = struct{}{}
			case 'S':
				start = Point{col, row}
			}
		}

		row++
	}

	return countTimelines(manifold, splitters, start, make(map[Point]int))
}

func countTimelines(manifold []string, splitters map[Point]struct{}, start Point, cache map[Point]int) int {
	if start.Y+1 >= len(manifold) {
		return 1
	}

	if _, exists := cache[start]; exists {
		return cache[start]
	}

	if _, exists := splitters[start]; exists {
		var count int

		if start.X-1 >= 0 && manifold[start.Y][start.X-1] != '^' {
			count += countTimelines(manifold, splitters, Point{start.X - 1, start.Y}, cache)
		}

		if start.X+1 < len(manifold[0]) && manifold[start.Y][start.X+1] != '^' {
			count += countTimelines(manifold, splitters, Point{start.X + 1, start.Y}, cache)
		}

		cache[start] = count
		return count
	}

	return countTimelines(manifold, splitters, Point{start.X, start.Y + 1}, cache)
}
