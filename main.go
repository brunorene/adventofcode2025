package main

import (
	"adventofcode2025/day11"
	"fmt"
)

func main() {
	// fmt.Printf("day01 - part 1: %d\n", day01.Solve1("day01/input"))
	// fmt.Printf("day01 - part 2: %d\n", day01.Solve2("day01/input"))
	// fmt.Printf("day02 - part 1: %d\n", day02.Solve1("day02/input"))
	// fmt.Printf("day02 - part 1: %d\n", day02.Solve2("day02/input"))
	// fmt.Printf("day03 - part 1: %d\n", day03.Solve("day03/input", 2))
	// fmt.Printf("day03 - part 2: %d\n", day03.Solve("day03/input", 12))
	// fmt.Printf("day04 - part 1: %d\n", day04.Solve1("day04/input"))
	// fmt.Printf("day04 - part 2: %d\n", day04.Solve2("day04/input"))
	// fmt.Printf("day05 - part 1: %d\n", day05.Solve1("day05/input"))
	// fmt.Printf("day05 - part 2: %d\n", day05.Solve2("day05/input"))
	// fmt.Printf("day06 - part 1: %d\n", day06.Solve1("day06/input"))
	// fmt.Printf("day06 - part 2: %d\n", day06.Solve2("day06/input"))
	// fmt.Printf("day07 - part 1: %d\n", day07.Solve1("day07/input"))
	// fmt.Printf("day07 - part 2: %d\n", day07.Solve2("day07/input"))
	// fmt.Printf("day08 - part 1: %d\n", day08.Solve1("day08/input", 1000))
	// fmt.Printf("day08 - part 2: %d\n", day08.Solve2("day08/input"))
	// fmt.Printf("day09 - part 1: %d\n", day09.Solve1("day09/input"))
	// fmt.Printf("day09 - part 2: %d\n", day09.Solve2("day09/input"))
	// fmt.Printf("day10 - part 1: %d\n", day10.Solve1("day10/input"))
	// fmt.Printf("day10 - part 2: %d\n", day10.Solve2("day10/input"))
	fmt.Printf("day11 - part 1: %d\n", day11.Solve("day11/input", []string{}, "you"))
	fmt.Printf("day11 - part 1: %d\n", day11.Solve("day11/input", []string{"dac", "fft"}, "svr"))
}
