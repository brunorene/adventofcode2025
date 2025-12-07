package day06_test

import (
	"adventofcode2025/day06"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay(t *testing.T) {
	assert.Equal(t, int64(4277556), day06.Solve1("testdata/input"))
	assert.Equal(t, int64(3263827), day06.Solve2("testdata/input"))
}
