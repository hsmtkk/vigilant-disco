package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/splitaddarray/kata"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input  []int
	repeat int
	want   []int
}

func TestSplitAdd(t *testing.T) {
	testCases := []testCase{
		{[]int{1, 2, 3, 4, 5}, 2, []int{5, 10}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{15}},
		{[]int{15}, 3, []int{15}},
		{[]int{1, 2, 3, 4}, 1, []int{4, 6}},
		{[]int{1, 2, 3, 4, 5, 6}, 20, []int{21}},
		{[]int{32, 45, 43, 23, 54, 23, 54, 34}, 2, []int{183, 125}},
		{[]int{32, 45, 43, 23, 54, 23, 54, 34}, 0, []int{32, 45, 43, 23, 54, 23, 54, 34}},
		{[]int{3, 234, 25, 345, 45, 34, 234, 235, 345}, 3, []int{305, 1195}},
		{[]int{3, 234, 25, 345, 45, 34, 234, 235, 345, 34, 534, 45, 645, 645, 645, 4656, 45, 3}, 4, []int{1040, 7712}},
		{[]int{23, 345, 345, 345, 34536, 567, 568, 6, 34536, 54, 7546, 456}, 20, []int{79327}},
	}
	for _, tc := range testCases {
		got := kata.SplitAndAdd(tc.input, tc.repeat)
		assert.Equal(t, tc.want, got)
	}
}
