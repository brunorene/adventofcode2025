package day05

import (
	"adventofcode2025/common"
	"strconv"
	"strings"
)

type freshRange struct {
	start, end int
}

func Solve1(input string) int {
	var freshRanges []freshRange
	var freshCount int

	for line := range common.ReadInput(input).ReadLines {
		if !strings.Contains(line, "-") {
			id, _ := strconv.Atoi(line)

			for _, idRange := range freshRanges {
				if id >= idRange.start && id <= idRange.end {
					freshCount++

					break
				}
			}

			continue
		}

		ranges := strings.Split(line, "-")
		start, _ := strconv.Atoi(ranges[0])
		end, _ := strconv.Atoi(ranges[1])
		freshRanges = append(freshRanges, freshRange{start, end})
	}

	return freshCount
}

func Solve2(input string) int64 {
	var freshCount int64
	var freshRanges []freshRange

	for line := range common.ReadInput(input).ReadLines {
		if !strings.Contains(line, "-") {
			break
		}

		ranges := strings.Split(line, "-")
		start, _ := strconv.Atoi(ranges[0])
		end, _ := strconv.Atoi(ranges[1])
		freshRanges = append(freshRanges, freshRange{start, end})
	}

	for id1 := 0; id1 < len(freshRanges); id1++ {
		for id2 := 0; id2 < len(freshRanges); id2++ {
			if id1 == id2 {
				continue
			}

			range1 := freshRanges[id1]
			range2 := freshRanges[id2]

			candidate := freshRange{min(range1.start, range2.start), max(range1.end, range2.end)}

			size := int64(candidate.end - candidate.start + 1)
			size1 := int64(range1.end - range1.start + 1)
			size2 := int64(range2.end - range2.start + 1)

			if size > size1+size2 {
				continue
			}

			freshRanges = append(freshRanges[0:max(id1, id2)], freshRanges[max(id1, id2)+1:]...)
			freshRanges = append(freshRanges[0:min(id1, id2)], freshRanges[min(id1, id2)+1:]...)
			freshRanges = append(freshRanges, candidate)

			id1 = 0
			break
		}
	}

	for _, idRange := range freshRanges {
		freshCount += int64(idRange.end - idRange.start + 1)
	}

	return freshCount
}
