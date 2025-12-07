package day01_test

import (
	"adventofcode2025/day01"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay(t *testing.T) {
	assert.Equal(t, 3, day01.Solve1("testdata/input"))
	assert.Equal(t, 6, day01.Solve2("testdata/input"))
}
