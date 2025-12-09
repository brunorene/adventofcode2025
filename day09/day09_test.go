package day09_test

import (
	"adventofcode2025/day09"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay(t *testing.T) {
	assert.Equal(t, 50, day09.Solve1("testdata/input"))
	assert.Equal(t, 24, day09.Solve2("testdata/input"))
}
