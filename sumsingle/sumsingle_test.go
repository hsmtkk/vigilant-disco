package sumsingle_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/sumsingle"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	numbers []int
	want    int
}

func TestSumSingle(t *testing.T) {
	testCases := []testCase{
		{[]int{4, 5, 7, 5, 4, 8}, 15},
		{[]int{9, 10, 19, 13, 19, 13}, 19},
		{[]int{16, 0, 11, 4, 8, 16, 0, 11}, 12},
		{[]int{5, 17, 18, 11, 13, 18, 11, 13}, 22},
		{[]int{5, 10, 19, 13, 10, 13}, 24},
	}
	for _, tc := range testCases {
		got := sumsingle.SumSingle(tc.numbers)
		assert.Equal(t, tc.want, got)
	}
}
