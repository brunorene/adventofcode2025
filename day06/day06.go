package day06

import (
	"adventofcode2025/common"
	"slices"
	"strings"
)

func Solve1(input string) int64 {
	var size int

	for line := range common.ReadInput(input).ReadLines {
		parts := strings.Fields(line)
		size = len(parts)
		break
	}

	numbers := make([][]int, size)

	var result int64

	for line := range common.ReadInput(input).ReadLines {
		parts := strings.Fields(line)

		if slices.Contains([]string{"*", "+"}, parts[0]) {

			for idx, operator := range parts {
				column := int64(-1)

				switch operator {
				case "*":
					for _, num := range numbers[idx] {
						if column == -1 {
							column = int64(num)

							continue
						}

						column *= int64(num)

					}
				case "+":
					for _, num := range numbers[idx] {
						if column == -1 {
							column = int64(num)

							continue
						}

						column += int64(num)
					}
				}

				result += column
			}

			continue
		}

		for idx, part := range parts {
			numbers[idx] = append(numbers[idx], common.AsInt(part))
		}
	}

	return result

}

func isSeparator(homework []string, idx int) bool {
	for _, row := range homework {
		if row[idx] != ' ' {
			return false
		}
	}

	return true
}

func components(homework []string, idx int) (number int64, operator byte) {
	var result int64

	for idxRow, row := range homework {
		if idxRow == len(homework)-1 {
			if row[idx] != ' ' {
				return result, row[idx]
			}

			break
		}

		if row[idx] == ' ' {
			continue
		}

		result *= 10
		result += int64(row[idx] - '0')
	}

	return result, ' '
}

func Solve2(input string) int64 {
	var homework []string
	for line := range common.ReadInput(input).ReadLines {
		homework = append(homework, line)
	}

	var result int64
	var operator byte
	column := int64(-1)

	for idxCol := range homework[0] {
		if isSeparator(homework, idxCol) {
			result += column
			column = int64(-1)

			continue
		}

		num, current := components(homework, idxCol)

		if current != ' ' {
			operator = current
		}

		switch operator {
		case '*':
			if column == -1 {
				column = int64(num)

				continue
			}

			column *= int64(num)
		case '+':
			if column == -1 {
				column = int64(num)

				continue
			}

			column += int64(num)
		}
	}
	
	result += column

	return result
}
