package kata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/hsmtkk/vigilant-disco/reducefraction/kata"
)

type testCase struct {
	input [2]int
	want  [2]int
}

func TestReduce(t *testing.T) {
	testCases := []testCase{
		{[2]int{60, 20}, [2]int{3, 1}},
		{[2]int{80, 120}, [2]int{2, 3}},
		{[2]int{4, 2}, [2]int{2, 1}},
		{[2]int{45, 120}, [2]int{3, 8}},
		{[2]int{1000, 1}, [2]int{1000, 1}},
		{[2]int{1, 1}, [2]int{1, 1}},
	}
	for _, tc := range testCases {
		got := kata.ReduceFraction(tc.input)
		assert.Equal(t, tc.want, got)
	}
}
