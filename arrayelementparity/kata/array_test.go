package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/arrayelementparity/kata"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input []int
	want  int
}

func TestSolve(t *testing.T) {
	testCases := []testCase{
		{[]int{1, -1, 2, -2, 3}, 3},
		{[]int{-3, 1, 2, 3, -1, -4, -2}, -4},
		{[]int{1, -1, 2, -2, 3, 3}, 3},
		{[]int{-110, 110, -38, -38, -62, 62, -38, -38, -38}, -38},
		{[]int{-9, -105, -9, -9, -9, -9, 105}, -9},
	}
	for _, tc := range testCases {
		got := kata.Solve(tc.input)
		assert.Equal(t, tc.want, got)
	}
}
