package day12

import (
	"adventofcode2025/common"
	"fmt"
)

func countHashtags(s string) int {
	count := 0
	for _, r := range s {
		if r == '#' {
			count++
		}
	}

	return count
}

func presents() []string {
	return []string{
		`..#
.##
##.`,
		`###
#.#
#.#`,
		`###
.##
.##`,
		`#..
##.
###`,
		`#.#
###
##.`,
		`###
.#.
###`,
	}
}

func Solve1(input string) int {
	presents := presents()
	var width, height, count0, count1, count2, count3, count4, count5, total, valid int

	for line := range common.ReadInput(input).ReadLines {
		fmt.Sscanf(line, "%dx%d: %d %d %d %d %d %d", &width, &height, &count0, &count1, &count2, &count3, &count4, &count5)
		total++
		if width*height < count0*countHashtags(presents[0])+
			count1*countHashtags(presents[1])+
			count2*countHashtags(presents[2])+
			count3*countHashtags(presents[3])+
			count4*countHashtags(presents[4])+
			count5*countHashtags(presents[5]) {
			continue
		}

		valid++
	}

	return valid
}
