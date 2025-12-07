package day02_test

import (
	"adventofcode2025/day02"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay(t *testing.T) {
	assert.Equal(t, 1227775554, day02.Solve1("testdata/input"))
	assert.Equal(t, 4174379265, day02.Solve2("testdata/input"))
}
