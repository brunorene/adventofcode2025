package day08

import (
	"adventofcode2025/common"
	"fmt"
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

	circuits := make(map[Point]map[Point]struct{})

	fmt.Println("merging circuits")
	for pair := range pairs {
		circuit1, exists := circuits[pair[0]]
		if !exists {
			circuit1 = make(map[Point]struct{})
		}

		circuit2, exists := circuits[pair[1]]
		if !exists {
			circuit2 = make(map[Point]struct{})
		}

		for point := range circuit1 {
			circuit2[point] = struct{}{}
		}

		circuit2[pair[0]] = struct{}{}
		circuit2[pair[1]] = struct{}{}

		for box := range circuit2 {
			circuits[box] = circuit2
		}
	}

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
	circuits := make(map[Point]map[Point]struct{})

	for pair := range pairs {
		circuit1, exists := circuits[pair[0]]
		if !exists {
			circuit1 = make(map[Point]struct{})
		}

		circuit2, exists := circuits[pair[1]]
		if !exists {
			circuit2 = make(map[Point]struct{})
		}

		for point := range circuit1 {
			circuit2[point] = struct{}{}
		}

		circuit2[pair[0]] = struct{}{}
		circuit2[pair[1]] = struct{}{}

		for box := range circuit2 {
			circuits[box] = circuit2
		}
	}

	if len(circuits) != boxCount {
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
