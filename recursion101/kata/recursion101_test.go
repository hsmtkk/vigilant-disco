package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/recursion101/kata"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	first  int
	second int
	want   []int
}

func TestSolve(t *testing.T) {
	testCases := []testCase{
		{6, 19, []int{6, 7}},
		{2, 1, []int{0, 1}},
		{22, 5, []int{0, 1}},
		{2, 10, []int{2, 2}},
		{8796203, 7556, []int{1019, 1442}},
		{19394, 19394, []int{19394, 19394}},
	}
	for _, tc := range testCases {
		got := kata.Solve(tc.first, tc.second)
		assert.Equal(t, tc.want, got)
	}
}
