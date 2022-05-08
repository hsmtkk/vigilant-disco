package maxminarray_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/maxminarray"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input []int
	want  []int
}

func TestMaxMinArray(t *testing.T) {
	testCases := []testCase{
		{[]int{15, 11, 10, 7, 12}, []int{15, 7, 12, 10, 11}},
		{[]int{91, 75, 86, 14, 82}, []int{91, 14, 86, 75, 82}},
		{[]int{84, 79, 76, 61, 78}, []int{84, 61, 79, 76, 78}},
		{[]int{52, 77, 72, 44, 74, 76, 40}, []int{77, 40, 76, 44, 74, 52, 72}},
		{[]int{1, 6, 9, 4, 3, 7, 8, 2}, []int{9, 1, 8, 2, 7, 3, 6, 4}},
		{[]int{78, 79, 52, 87, 16, 74, 31, 63, 80}, []int{87, 16, 80, 31, 79, 52, 78, 63, 74}},
	}
	for _, tc := range testCases {
		got := maxminarray.MaxMinArray(tc.input)
		assert.Equal(t, tc.want, got)
	}
}
