package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input []string
	want  []int
}

func Test0(t *testing.T) {
	testCases := []testCase{
		{[]string{"abode", "ABc", "xyzD"}, []int{4, 3, 1}},
		{[]string{"abide", "ABc", "xyz"}, []int{4, 3, 0}},
		{[]string{"IAMDEFANDJKL", "thedefgh", "xyzDEFghijabc"}, []int{6, 5, 7}},
		{[]string{"encode", "abc", "xyzD", "ABmD"}, []int{1, 3, 1, 3}},
	}
	for _, testCase := range testCases {
		got := solve(testCase.input)
		assert.Equal(t, testCase.want, got)
	}
}
