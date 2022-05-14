package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/wordvalues/kata"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input []string
	want  []int
}

func TestWordValues(t *testing.T) {
	testCases := []testCase{
		{[]string{"abc", "abc", "abc", "abc"}, []int{6, 12, 18, 24}},
		{[]string{"codewars", "abc", "xyz"}, []int{88, 12, 225}},
	}
	for _, tc := range testCases {
		got := kata.NameValue(tc.input)
		assert.Equal(t, tc.want, got)
	}
}
