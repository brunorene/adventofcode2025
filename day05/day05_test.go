package day05_test

import (
	"adventofcode2025/day05"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay(t *testing.T) {
	assert.Equal(t, 3, day05.Solve1("testdata/input"))
	assert.Equal(t, int64(14), day05.Solve2("testdata/input"))
}
