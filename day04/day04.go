package day04

import (
	"adventofcode2025/common"
)

type coords struct {
	row, col int
}

func Solve1(input string) int {
	var rolls []string

	for line := range common.ReadInput(input).ReadLines {
		rolls = append(rolls, line)
	}

	var accessibleRolls int

	for row, line := range rolls {
		for col, char := range line {
			if char == '@' {
				var countBorder int

				for _, y := range []int{-1, 0, 1} {
					for _, x := range []int{-1, 0, 1} {
						if (x == y && x == 0) ||
							row+y < 0 ||
							row+y == len(rolls) ||
							col+x < 0 ||
							col+x == len(rolls[row+y]) {
							continue
						}

						if rolls[row+y][col+x] == '@' {
							countBorder++
						}
					}
				}

				if countBorder < 4 {
					accessibleRolls++
				}
			}
		}
	}

	return accessibleRolls
}

func Solve2(input string) int {
	var rolls []string

	for line := range common.ReadInput(input).ReadLines {
		rolls = append(rolls, line)
	}

	var result int

	for {
		var accessibleRolls []coords

		for row, line := range rolls {
			for col, char := range line {
				if char == '@' {
					var countBorder int

					for _, y := range []int{-1, 0, 1} {
						for _, x := range []int{-1, 0, 1} {
							if (x == y && x == 0) ||
								row+y < 0 ||
								row+y == len(rolls) ||
								col+x < 0 ||
								col+x == len(rolls[row+y]) {
								continue
							}

							if rolls[row+y][col+x] == '@' {
								countBorder++
							}
						}
					}

					if countBorder < 4 {
						accessibleRolls = append(accessibleRolls, coords{row, col})
					}
				}
			}
		}

		if len(accessibleRolls) == 0 {
			break
		}

		result += len(accessibleRolls)

		for _, roll := range accessibleRolls {
			rolls[roll.row] = rolls[roll.row][:roll.col] + "." + rolls[roll.row][roll.col+1:]
		}
	}

	return result
}
