package dominantarray_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/dominantarray"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input []int
	want  []int
}

func TestSolve(t *testing.T) {
	testCases := []testCase{
		{[]int{16, 17, 14, 3, 14, 5, 2}, []int{17, 14, 5, 2}},
		{[]int{92, 52, 93, 31, 89, 87, 77, 105}, []int{105}},
		{[]int{75, 47, 42, 56, 13, 55}, []int{75, 56, 55}},
		{[]int{67, 54, 27, 85, 66, 88, 31, 24, 49}, []int{88, 49}},
		{[]int{76, 17, 25, 36, 29}, []int{76, 36, 29}},
		{[]int{104, 18, 37, 9, 36, 47, 28}, []int{104, 47, 28}},
	}
	for _, tc := range testCases {
		got := dominantarray.Solve(tc.input)
		assert.Equal(t, tc.want, got)
	}
}
