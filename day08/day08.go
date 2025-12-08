package day08

import (
	"adventofcode2025/common"
	"fmt"
	"maps"
	"math"
	"slices"
)

type Point struct {
	X, Y, Z int
}

func ShortestPair(points []Point, except map[[2]Point]struct{}) [2]Point {
	var pair [2]Point
	minDistance := math.MaxFloat64

	for _, p1 := range points {
		for _, p2 := range points {
			if p1 == p2 {
				continue
			}

			distance := math.Sqrt(math.Pow(float64(p1.X-p2.X), 2.0) + math.Pow(float64(p1.Y-p2.Y), 2.0) + math.Pow(float64(p1.Z-p2.Z), 2.0))
			_, exists := except[[2]Point{p1, p2}]
			_, existsRev := except[[2]Point{p2, p1}]

			if distance < minDistance && !exists && !existsRev {
				minDistance = distance
				pair = [2]Point{p1, p2}
			}
		}
	}

	return pair
}

func asSortedSlice(points map[Point]struct{}) (result []Point) {
	result = make([]Point, 0, len(points))

	for point := range points {
		result = append(result, point)
	}

	slices.SortFunc(result, SortSlice)

	return result
}

func SortSlice(p1, p2 Point) int {
	diffX := p1.X - p2.X
	if diffX != 0 {
		return diffX
	}

	diffY := p1.Y - p2.Y
	if diffY != 0 {
		return diffY
	}

	diffZ := p1.Z - p2.Z
	if diffZ != 0 {
		return diffZ
	}

	return 0
}

func Merge(pairs map[[2]Point]struct{}) map[Point]map[Point]struct{} {
	circuits := make(map[Point]map[Point]struct{})

	for pair := range pairs {
		circuit1, exists := circuits[pair[0]]
		if !exists {
			circuit1 = map[Point]struct{}{pair[0]: {}}
		}

		circuit2, exists := circuits[pair[1]]
		if !exists {
			circuit2 = map[Point]struct{}{pair[1]: {}}
		}

		if len(circuit1) > 0 && maps.Equal(circuit1, circuit2) {
			continue
		}

		for point := range circuit1 {
			circuit2[point] = struct{}{}
		}

		for box := range circuit2 {
			circuits[box] = circuit2
		}
	}

	return circuits
}

func Solve1(input string, size int) int {
	var boxes []Point

	fmt.Println("reading input")
	for line := range common.ReadInput(input).ReadLines {
		var x, y, z int
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		boxes = append(boxes, Point{X: x, Y: y, Z: z})
	}

	pairs := make(map[[2]Point]struct{})

	fmt.Println("shortest pairs")
	for range size {
		shortest := ShortestPair(boxes, pairs)
		pairs[shortest] = struct{}{}
	}

	fmt.Println("merging circuits")
	circuits := Merge(pairs)

	uniqueCircuits := make([][]Point, 0, len(circuits))

	fmt.Println("unique circuits")
	for _, current := range circuits {
		circuit := asSortedSlice(current)

		if slices.IndexFunc(uniqueCircuits, func(p []Point) bool {
			return slices.Equal(circuit, p)
		}) == -1 {
			uniqueCircuits = append(uniqueCircuits, circuit)
		}
	}

	fmt.Println("sorting circuits")
	slices.SortFunc(uniqueCircuits, func(p1, p2 []Point) int { return len(p2) - len(p1) })

	return len(uniqueCircuits[0]) * len(uniqueCircuits[1]) * len(uniqueCircuits[2])
}

func allConnected(pairs map[[2]Point]struct{}, boxCount int) bool {
	circuits := Merge(pairs)

	if len(circuits) < boxCount {
		return false
	}

	fmt.Println("all on circuit")

	for _, circuit := range circuits {
		if len(circuit) == boxCount {
			return true
		}
	}

	return false
}

func Solve2(input string) int {
	var boxes []Point

	fmt.Println("reading input")
	for line := range common.ReadInput(input).ReadLines {
		var x, y, z int
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		boxes = append(boxes, Point{X: x, Y: y, Z: z})
	}

	pairs := make(map[[2]Point]struct{})

	fmt.Println("shortest pairs")
	for {
		shortest := ShortestPair(boxes, pairs)
		pairs[shortest] = struct{}{}

		if allConnected(pairs, len(boxes)) {
			return shortest[0].X * shortest[1].X
		}
	}
}
