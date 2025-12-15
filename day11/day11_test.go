package day11_test

import (
	"adventofcode2025/day11"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay(t *testing.T) {
	assert.Equal(t, 5, day11.Solve("testdata/input1", []string{}, "you"))
	assert.Equal(t, 2, day11.Solve("testdata/input2", []string{"dac", "fft"}, "svr"))
}
