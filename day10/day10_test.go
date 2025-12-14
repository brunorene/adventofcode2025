package day10_test

import (
	"adventofcode2025/day10"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay(t *testing.T) {
	assert.Equal(t, 7, day10.Solve1("testdata/input"))
	assert.Equal(t, 33, day10.Solve2("testdata/input"))
}
