package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/casinochip/kata"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input []int
	want  int
}

func Test0(t *testing.T) {
	cases := []testCase{
		{[]int{1, 1, 1}, 1},
		{[]int{1, 2, 1}, 2},
		{[]int{4, 1, 1}, 2},
		{[]int{8, 2, 8}, 9},
		{[]int{8, 1, 4}, 5},
		{[]int{7, 4, 10}, 10},
		{[]int{12, 12, 12}, 18},
		{[]int{1, 23, 2}, 3},
	}
	for _, test := range cases {
		got := kata.Solve(test.input)
		assert.Equal(t, test.want, got)
	}
}
