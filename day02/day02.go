package day02

import (
	"adventofcode2025/common"
	"fmt"
	"strconv"
	"strings"
)

func invalidID1(id int) bool {
	str := fmt.Sprintf("%d", id)

	if len(str)%2 == 1 {
		return false
	}

	return str[0:len(str)/2] == str[len(str)/2:]
}

func invalidID2(id int) bool {
	str := fmt.Sprintf("%d", id)

	for size := 1; size <= len(str)/2; size++ {
		if len(str)%size != 0 || len(str) == size {
			continue
		}

		if strings.Replace(str, str[0:size], "", len(str)/size) == "" {
			return true
		}
	}

	return false
}

func Solve(input string, invalidID func(int) bool) int {
	line := common.ReadInput(input).Read()

	ranges := strings.Split(line, ",")

	var sum int

	for _, rangeStr := range ranges {
		firstAndLast := strings.Split(rangeStr, "-")

		first, err := strconv.Atoi(firstAndLast[0])
		common.CheckError(err)

		last, err := strconv.Atoi(strings.ReplaceAll(firstAndLast[1], "\n", ""))
		common.CheckError(err)

		for i := first; i <= last; i++ {
			if invalidID(i) {
				sum += i
			}
		}
	}

	return sum
}

func Solve1(input string) int {
	return Solve(input, invalidID1)
}

func Solve2(input string) int {
	return Solve(input, invalidID2)
}
