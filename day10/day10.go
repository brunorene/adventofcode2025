package day10

import (
	"adventofcode2025/common"
	"fmt"
	"math"
	"os"
	"os/exec"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type machine struct {
	lights  int
	buttons []int
}

func (m machine) String() string {
	return fmt.Sprintf("lights: %d, buttons: %v", m.lights, m.buttons)
}

var lightRegex = regexp.MustCompile(`\[(.*?)\]`)
var buttonRegex = regexp.MustCompile(`\((.*?)\)`)

func parse1(line string) machine {
	lightMatch := lightRegex.FindStringSubmatch(line)
	buttonMatch := buttonRegex.FindAllStringSubmatch(line, -1)

	var lights int
	for idx, light := range lightMatch[1] {
		if light == '#' {
			lights += int(math.Pow(2, float64(idx)))
		}
	}

	buttons := make([]int, len(buttonMatch))
	for idx := range buttonMatch {
		for num := range strings.SplitSeq(buttonMatch[idx][1], ",") {
			buttons[idx] += int(math.Pow(2, float64(common.AsInt(num))))
		}
	}

	return machine{lights, buttons}
}

func (m *machine) solve() int {
	n := len(m.buttons)
	minPresses := math.MaxInt

	for mask := 1; mask < (1 << n); mask++ {
		result := 0
		presses := 0
		for i := range n {
			if mask&(1<<i) != 0 {
				result ^= m.buttons[i]
				presses++
			}
		}

		if result == m.lights && presses < minPresses {
			minPresses = presses
		}
	}

	return minPresses
}

func Solve1(input string) int {
	var sum int

	for line := range common.ReadInput(input).ReadLines {
		machine := parse1(line)

		fmt.Println(machine)

		minSteps := machine.solve()

		fmt.Println("min", minSteps)

		sum += minSteps
	}

	return sum
}

// Part 2 types and functions

type machine2 struct {
	buttons [][]int // each button is a list of counter indices it affects
	targets []int   // joltage requirements
}

var joltageRegex = regexp.MustCompile(`\{(.*?)\}`)

func parse2(line string) machine2 {
	buttonMatch := buttonRegex.FindAllStringSubmatch(line, -1)
	joltageMatch := joltageRegex.FindStringSubmatch(line)

	buttons := make([][]int, len(buttonMatch))
	for idx := range buttonMatch {
		for num := range strings.SplitSeq(buttonMatch[idx][1], ",") {
			buttons[idx] = append(buttons[idx], common.AsInt(num))
		}
	}

	var targets []int
	for num := range strings.SplitSeq(joltageMatch[1], ",") {
		targets = append(targets, common.AsInt(num))
	}

	return machine2{buttons, targets}
}

func (m *machine2) solveGLPK() int {
	var lp strings.Builder
	nButtons := len(m.buttons)

	// Objective: minimize sum of all button presses
	lp.WriteString("Minimize\n obj: ")
	for i := range nButtons {
		if i > 0 {
			lp.WriteString(" + ")
		}
		lp.WriteString(fmt.Sprintf("x%d", i))
	}
	lp.WriteString("\n")

	// Constraints: each counter must reach its target
	lp.WriteString("Subject To\n")
	for j, target := range m.targets {
		lp.WriteString(fmt.Sprintf(" c%d: ", j))
		first := true
		for i, button := range m.buttons {
			if slices.Contains(button, j) {
				if !first {
					lp.WriteString(" + ")
				}
				lp.WriteString(fmt.Sprintf("x%d", i))
				first = false
			}
		}
		lp.WriteString(fmt.Sprintf(" = %d\n", target))
	}

	// Bounds: all variables >= 0
	lp.WriteString("Bounds\n")
	for i := range nButtons {
		lp.WriteString(fmt.Sprintf(" x%d > 0\n", i))
	}

	// Integer variables
	lp.WriteString("General\n")
	for i := range nButtons {
		lp.WriteString(fmt.Sprintf(" x%d\n", i))
	}

	lp.WriteString("End\n")

	// Write to temp file
	tmpFile, err := os.CreateTemp("", "problem*.lp")
	common.CheckError(err)
	tmpFile.WriteString(lp.String())
	tmpFile.Close()

	outFile, err := os.CreateTemp("", "solution*.txt")
	common.CheckError(err)
	outFile.Close()

	// Run GLPK solver
	cmd := exec.Command("glpsol", "--lp", tmpFile.Name(), "-o", outFile.Name())
	err = cmd.Run()
	common.CheckError(err)

	// Parse solution
	solution, err := os.ReadFile(outFile.Name())
	common.CheckError(err)

	// Clean up
	os.Remove(tmpFile.Name())
	os.Remove(outFile.Name())

	// Extract objective value
	re := regexp.MustCompile(`Objective:\s+obj\s+=\s+(\d+)`)
	match := re.FindStringSubmatch(string(solution))
	if match != nil {
		result, _ := strconv.Atoi(match[1])
		return result
	}

	return -1
}

func Solve2(input string) int {
	var sum int

	for line := range common.ReadInput(input).ReadLines {
		machine := parse2(line)

		minPresses := machine.solveGLPK()

		fmt.Println("min", minPresses)

		sum += minPresses
	}

	return sum
}
