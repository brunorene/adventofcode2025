package day07_test

import (
	"adventofcode2025/day07"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay(t *testing.T) {
	assert.Equal(t, 21, day07.Solve1("testdata/input"))
	assert.Equal(t, 40, day07.Solve2("testdata/input"))
}
