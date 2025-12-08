package day08_test

import (
	"adventofcode2025/day08"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay(t *testing.T) {
	// assert.Equal(t, 40, day08.Solve1("testdata/input", 10))
	assert.Equal(t, 25272, day08.Solve2("testdata/input"))
}
