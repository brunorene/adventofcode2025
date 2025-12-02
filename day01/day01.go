package day01

import (
	"adventofcode2025/common"
	"strconv"
)

func Solve1(input string) int {
	current := 50

	var zeroCount int

	for line := range common.ReadInput(input).ReadLines {
		turn, err := strconv.Atoi(line[1:])
		common.CheckError(err)

		if line[0] == 'R' {
			current = (current + turn) % 100
		} else {
			current = (current - turn) % 100
		}

		if current == 0 {
			zeroCount++
		}
	}

	return zeroCount
}

func Solve2(input string) int {
	current := 50

	var zeroCount int

	for line := range common.ReadInput(input).ReadLines {
		turn, err := strconv.Atoi(line[1:])
		common.CheckError(err)

		zeroCount += turn / 100
		turn %= 100

		if line[0] == 'R' {
			current += turn
		} else {
			current -= turn
		}

		if (current <= 0 || current >= 100) && (current != turn && current != -turn) {
			zeroCount++
		}

		if current < 0 {
			current = 100 + current%100

			continue
		}

		current %= 100
	}

	return zeroCount
}
