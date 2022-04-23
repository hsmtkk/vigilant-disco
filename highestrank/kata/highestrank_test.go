package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/highestrank/kata"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input []int
	want  int
}

func TestXxx(t *testing.T) {
	testCases := []testCase{
		{[]int{12, 10, 8, 12, 7, 6, 4, 10, 12}, 12},
		{[]int{12, 10, 8, 12, 7, 6, 4, 10, 12, 10}, 12},
		{[]int{12, 10, 8, 8, 3, 3, 3, 3, 2, 4, 10, 12, 10}, 3},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.want, kata.HighestRank(tc.input))
	}
}
