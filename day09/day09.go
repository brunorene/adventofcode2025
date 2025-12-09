package day09

import (
	"adventofcode2025/common"
	"fmt"
	"math"
	"os"
	"sort"
)

type color int

const (
	red color = iota
	green
)

type point struct {
	x, y int
}

func area(p1, p2 point) int {
	xDiff := math.Abs(float64(p1.x-p2.x)) + 1
	yDiff := math.Abs(float64(p1.y-p2.y)) + 1

	area := int(xDiff * yDiff)

	return area
}

func Solve1(input string) int {
	var points []point

	for line := range common.ReadInput(input).ReadLines {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		points = append(points, point{x: x, y: y})
	}

	var maxArea int

	for _, p1 := range points {
		for _, p2 := range points {
			area := area(p1, p2)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func connect(points []point) map[point]color {
	connected := make(map[point]color)
	for i := range points {
		connected[points[i]] = red
	}

	return connected
}

func borderGreen(points []point, connected map[point]color) {
	for idx, curr := range points {
		next := points[(idx+1)%len(points)]

		for y := min(curr.y, next.y); y <= max(curr.y, next.y); y++ {
			for x := min(curr.x, next.x); x <= max(curr.x, next.x); x++ {
				col, exists := connected[point{x, y}]
				if exists && col == red {
					continue
				}

				connected[point{x, y}] = green
			}
		}
	}
}

func maximums(connected map[point]color) (maxX, maxY int) {
	for p := range connected {
		if p.x > maxX {
			maxX = p.x
		}

		if p.y > maxY {
			maxY = p.y
		}
	}

	return maxX, maxY
}

func draw(connected map[point]color) {
	drawing, err := os.Create("drawing.txt")
	common.CheckError(err)
	defer drawing.Close()

	maxX, maxY := maximums(connected)

	for y := range maxY + 1 {
		for x := range maxX + 1 {
			col, exists := connected[point{x, y}]
			if !exists {
				drawing.WriteString(".")

				continue
			}

			switch col {
			case red:
				drawing.WriteString("#")
			case green:
				drawing.WriteString("X")
			}
		}

		drawing.WriteString("\n")
	}
}

func validArea(p1, p2 point, connected map[point]color) bool {
	for pt := range connected {
		if min(p1.x, p2.x) < pt.x && max(p1.x, p2.x) > pt.x && min(p1.y, p2.y) < pt.y && max(p1.y, p2.y) > pt.y {
			return false
		}
	}

	return true
}

func Solve2(input string) int {
	finish := common.Timer("day 09 part 2")
	defer finish()

	var points []point

	for line := range common.ReadInput(input).ReadLines {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		points = append(points, point{x: x, y: y})
	}

	connected := connect(points)

	borderGreen(points, connected)

	// draw(connected)

	// sort by area
	type rect struct {
		p1, p2 point
		area   int
	}

	areas := []rect{}
	visited := make(map[[2]point]bool)

	for _, p1 := range points {
		for _, p2 := range points {
			if p1 == p2 || visited[[2]point{p1, p2}] || visited[[2]point{p2, p1}] {
				continue
			}

			areas = append(areas, rect{p1: p1, p2: p2, area: area(p1, p2)})
			visited[[2]point{p1, p2}] = true
			visited[[2]point{p2, p1}] = true
		}
	}

	sort.Slice(areas, func(i, j int) bool {
		return areas[i].area > areas[j].area
	})

	for _, current := range areas {
		if validArea(current.p1, current.p2, connected) {
			return current.area
		}
	}

	return 0
}
