package simpleconsecutivepairs_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/simpleconsecutivepairs"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input []int
	want  int
}

func TestCout(t *testing.T) {
	testCases := []testCase{
		{[]int{1, 2, 5, 8, -4, -3, 7, 6, 5}, 3},
		{[]int{21, 20, 22, 40, 39, -56, 30, -55, 95, 94}, 2},
		{[]int{81, 44, 80, 26, 12, 27, -34, 37, -35}, 0},
		{[]int{-55, -56, -7, -6, 56, 55, 63, 62}, 4},
		{[]int{73, 72, 8, 9, 73, 72}, 3},
	}
	for _, tc := range testCases {
		got := simpleconsecutivepairs.Count(tc.input)
		assert.Equal(t, tc.want, got)
	}
}
