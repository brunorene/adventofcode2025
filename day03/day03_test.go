package day03_test

import (
	"adventofcode2025/day03"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay01(t *testing.T) {
	assert.Equal(t, 357, day03.Solve("testdata/input", 2))
	assert.Equal(t, 3121910778619, day03.Solve("testdata/input", 12))
}
