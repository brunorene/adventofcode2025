package day04_test

import (
	"adventofcode2025/day04"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay01(t *testing.T) {
	assert.Equal(t, 13, day04.Solve1("testdata/input"))
	assert.Equal(t, 43, day04.Solve2("testdata/input"))
}
