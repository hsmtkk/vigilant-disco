package findnthdigit_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/findnthdigit"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	number int
	nth    int
	want   int
}

func TestFindNthDigit(t *testing.T) {
	testCases := []testCase{
		{5673, 4, 5},
		{129, 2, 2},
		{-2825, 3, 8},
		{-456, 4, 0},
		{0, 20, 0},
		{65, 0, -1},
		{24, -8, -1},
	}
	for _, testCase := range testCases {
		got := findnthdigit.FindNthDigit(testCase.number, testCase.nth)
		assert.Equal(t, testCase.want, got)
	}
}
