package day03

import (
	"adventofcode2025/common"
)

func Solve(input string, count int) int {
	var sum int

	for line := range common.ReadInput(input).ReadLines {
		var current int
		idxMax := -1

		for back := count - 1; back >= 0; back-- {
			var max int
			var idxLocal int

			for idx := idxMax + 1; idx < len(line)-back; idx++ {
				num := int(line[idx] - '0')

				if num > max {
					max = num
					idxLocal = idx
				}
			}

			idxMax = idxLocal
			current = current*10 + max
		}

		sum += current
	}

	return sum
}
